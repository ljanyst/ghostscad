// Copyright 2021 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var codeTemplate = `
// Do not edit. Automatically generated on {{ .Timestamp }}.

package {{ .PackageName }}

import (
{{ range $path, $qualifier := .Qualifiers }}
  {{ $qualifier }} "{{ $path }}"
{{ end }}
)

{{ range $type, $optionals := .Optionals }}
{{ range $i, $opt := $optionals }}
func (o *{{ $type }}) Set{{ $opt.FieldName }}(val {{ $opt.Type }}) *{{ $type }} {
  o.{{ $opt.FieldName }} = val
  return o
}
{{ end }}
{{ end }}

{{ range $type, $forwards := .Forwards }}
{{ range $i, $forward := $forwards }}
func (o *{{ $type }}) {{ $forward.MethodName }}{{ $forward.Signature }} *{{ $type }} {
  o.{{ $forward.FieldName }}.{{ $forward.MethodName }}({{ $forward.VariableList }})
  return o
}
{{ end }}
{{ end }}
`

type Optional struct {
	FieldName string
	Type      string
}

type Forward struct {
	FieldName    string
	MethodName   string
	Signature    string
	VariableList string
}

type TemplateParams struct {
	PackageName string
	Optionals   map[string][]Optional
	Forwards    map[string][]Forward
	Timestamp   string
	Qualifiers  map[string]string
}

type Generator struct {
	ParsedFiles map[string]*ast.File
	Types       map[*ast.Ident]types.Object
	Params      TemplateParams
}

const qualifierChars = "abcdefghijklmnopqrstuvwxyz"

// Generate a random qualifier of given length
func RandomQualifier(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = qualifierChars[rand.Intn(len(qualifierChars))]
	}
	return string(b)
}

// Parse the go files in current package and type check the identifiers
func (g *Generator) ParseCurrentPackage() error {
	pkg, err := build.Default.ImportDir(".", 0)
	if err != nil {
		return err
	}

	g.Params.PackageName = pkg.Name
	g.Params.Timestamp = time.Now().Format(time.UnixDate)
	g.Params.Optionals = make(map[string][]Optional)
	g.Params.Forwards = make(map[string][]Forward)

	// Initialize the map of qualifiers for parameter types, we generate them at random
	// for each package path to avoid conflicts
	g.Params.Qualifiers = make(map[string]string)

	g.ParsedFiles = make(map[string]*ast.File)
	g.Types = make(map[*ast.Ident]types.Object)
	parsedFiles := []*ast.File{}

	// Parse all the files in the current package
	fileSet := token.NewFileSet()
	for _, name := range pkg.GoFiles {
		parsedFile, err := parser.ParseFile(fileSet, name, nil, parser.AllErrors)
		if err != nil {
			return err
		}
		log.Debugf("Successfully parsed %s", name)
		g.ParsedFiles[name] = parsedFile
		parsedFiles = append(parsedFiles, parsedFile)
	}

	// Type check all the identifiers
	config := types.Config{
		IgnoreFuncBodies: true,
		Importer:         importer.For("source", nil),
		FakeImportC:      true,
	}

	info := types.Info{
		Defs: g.Types,
	}

	_, err = config.Check(".", fileSet, parsedFiles, &info)
	if err != nil {
		return err
	}

	return nil
}

// Get the qualifier for current type and register the appropriate import, if the
// package the type is defined in has not yet been encountered
func (g *Generator) getQualifier(p *types.Package) string {
	// No qualifiers for the current package
	if p.Path() == "." {
		return ""
	}

	if qualifier, ok := g.Params.Qualifiers[p.Path()]; ok {
		return qualifier
	}

	// Since we generate short random identifiers, we need to be sure we don't have
	// collisions
QualifierGen:
	for i := 0; i < 1000; i++ {
		qualifier := RandomQualifier(3)
		for k, _ := range g.Params.Qualifiers {
			if k == qualifier {
				continue QualifierGen
			}
		}
		g.Params.Qualifiers[p.Path()] = qualifier
		return qualifier
	}

	// Tough luck, run out of combinations
	log.Fatalf("Cannot generate a qualifier for package %s", p.Path())
	return ""
}

func (g *Generator) processOptional(typeName string, field *ast.Field) error {
	optional := Optional{
		FieldName: field.Names[0].Name,
	}

	typ, ok := g.Types[field.Names[0]]
	if !ok {
		return fmt.Errorf("Cannot resolve type of field %s of %s.%s",
			optional.FieldName, g.Params.PackageName, typeName)
	}

	optional.Type = types.TypeString(typ.Type(), g.getQualifier)
	log.Infof("Found optional field %s of type %s", optional.FieldName, typ.Type())
	g.Params.Optionals[typeName] = append(g.Params.Optionals[typeName], optional)
	return nil
}

func (g *Generator) processForward(typeName string, field *ast.Field, methods []string) error {
	// Resolve the type of the field to figure out its methods
	typ, ok := g.Types[field.Names[0]]
	if !ok {
		return fmt.Errorf("Cannot resolve type of field %s of %s.%s",
			field.Names[0].Name, g.Params.PackageName, typeName)
	}

	methodSet := types.NewMethodSet(typ.Type())
	for _, m := range methods {
		forward := Forward{
			FieldName:  field.Names[0].Name,
			MethodName: m,
		}

		// Check if the method we're looking for is what it's supposed to be
		sel := methodSet.Lookup(typ.Pkg(), m)
		if sel == nil || sel.Kind() != types.MethodVal {
			return fmt.Errorf("Cannot find method %s in %s", m, typ.Type())
		}

		funcObj, ok := sel.Obj().(*types.Func)
		if !ok {
			return fmt.Errorf("Selection %s is not a function", m)
		}

		// Write down the signature of the method including type qualification
		sign, ok := funcObj.Type().(*types.Signature)
		if !ok {
			return fmt.Errorf("Cannot get function signature for %s", m)
		}

		var buffer bytes.Buffer
		newSign := types.NewSignature(nil, sign.Params(), nil, sign.Variadic())
		types.WriteSignature(&buffer, newSign, g.getQualifier)
		forward.Signature = buffer.String()

		// Create the variable list to call the original method if there are any
		// If the function is variadic, then the last parameter needs to be expanded
		params := newSign.Params()
		length := params.Len()
		if length != 0 {
			buffer.Reset()
			length--

			for i := 0; i < length; i++ {
				buffer.WriteString(params.At(i).Name())
				buffer.WriteString(", ")
			}

			buffer.WriteString(params.At(length).Name())
			if sign.Variadic() {
				buffer.WriteString("...")
			}
			forward.VariableList = buffer.String()
		}

		log.Infof("Found forward method %s%s in field %s",
			forward.MethodName, forward.Signature, forward.FieldName)
		g.Params.Forwards[typeName] = append(g.Params.Forwards[typeName], forward)
	}
	return nil
}

func (g *Generator) DiscoverTypes() error {
	log.Infof("Discovering types")

	// Search through the files in the package for type that we can generate methods for
	var fields []*ast.Field
	for fileName, parsedFile := range g.ParsedFiles {
		log.Infof("Looking at %s", fileName)
		for _, decl := range parsedFile.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE || len(genDecl.Specs) != 1 {
				continue
			}

			typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
			if !ok {
				continue
			}
			typeName := typeSpec.Name.Name
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			fields = structType.Fields.List
			if fields == nil {
				continue
			}

			log.Debugf("Processing %s", typeName)
			for _, field := range fields {
				if field.Tag == nil || field.Tag.Kind != token.STRING {
					continue
				}

				tag := field.Tag.Value[1 : len(field.Tag.Value)-1]
				if tag == "optional" {
					if err := g.processOptional(typeName, field); err != nil {
						return err
					}
				}

				if strings.HasPrefix(tag, "forward") {
					if err := g.processForward(typeName, field, strings.Split(tag[8:], ",")); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (g *Generator) GenerateCode(fileName string) error {
	t, err := template.New("code").Parse(codeTemplate)
	if err != nil {
		panic(err)
	}

	log.Infof("Writing code to %s", fileName)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = t.Execute(file, g.Params)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "fmt", fileName)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	logLevel := flag.String("log-level", "Info", "verbosity of the diagnostic information")
	flag.Parse()

	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})

	level := log.InfoLevel
	if *logLevel != "" {
		l, err := log.ParseLevel(*logLevel)
		if err != nil {
			log.Fatal(err)
		}
		level = l
	}
	log.SetLevel(level)

	g := Generator{}

	// We want repeatable identifiers
	rand.Seed(1354)

	if err := g.ParseCurrentPackage(); err != nil {
		log.Fatalf("Cannot parse current package: %s", err)
	}

	if err := g.DiscoverTypes(); err != nil {
		log.Fatalf("Cannot discover types: %s", err)
	}

	if err := g.GenerateCode("tagged_methods.go"); err != nil {
		log.Fatalf("Cannot generate code: %s", err)
	}
}
