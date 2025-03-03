package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type templateEnum struct {
	Package string
	Type    string
	Entries []templateEnumEntry
}

type templateEnumEntry struct {
	Name  string
	Value string
	Alert string
}

const templateStr = `
// Code generated by go generate; DO NOT EDIT.
// This file was generated by genenum.go

package {{.Package}}

import "fmt"

var names{{.Type}} = map[{{.Type}}]string{
{{range .Entries}}{{.Value}}: "{{.Name}}",{{if .Alert}} // {{.Alert}}{{end}}
{{end}}
}

func (v {{.Type}}) String() string {
    return names{{.Type}}[v]
}

{{range .Entries}}
type {{.Name}}Error struct {
    Description string
}

func (v {{.Name}}Error) Error() string {
    return fmt.Sprintf("HTTP {{.Value}} %s", v.Description) 
}

{{if .Alert}}
func (v {{.Name}}Error) Alert() string {
    return "{{.Alert}}"
}
{{end}}
{{end}}
`

var tmpl = template.Must(template.New("enum").Parse(templateStr))

func main() {
	fname := os.Getenv("GOFILE") // имя файла, из которого вызван go:generate,
	// установлено в переменной окружения $GOFILE
	typ := os.Args[1]

	// получаем AST
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var entries []templateEnumEntry
	// итерируемся по всем определениям в корне файла
	for _, d := range f.Decls {
		// ищем блок с константами
		var gd *ast.GenDecl
		var ok bool
		if gd, ok = d.(*ast.GenDecl); !ok {
			continue
		}
		if gd.Tok != token.CONST {
			continue
		}

		// итерируемся по всем конcтантам в блоке
		for _, s := range gd.Specs {
			vs := s.(*ast.ValueSpec)

			// получаем тип константы
			var it *ast.Ident
			if it, ok = vs.Type.(*ast.Ident); !ok {
				continue
			}
			if it.Name != typ {
				continue
			}

			// получаем имя константы
			var name string
			if len(vs.Names) < 1 {
				continue
			}
			name = vs.Names[0].Name

			// получаем значение константы
			var value string
			var bl *ast.BasicLit
			if len(vs.Values) < 1 {
				continue
			}
			if bl, ok = vs.Values[0].(*ast.BasicLit); !ok {
				continue
			}
			value = bl.Value

			// получаем значение комментария
			var alert string
			if vs.Comment != nil {
				for _, c := range vs.Comment.List {
					alertIdx := strings.Index(c.Text, "alert: ")
					if alertIdx == -1 {
						continue
					}
					alert = c.Text[alertIdx+len("alert: "):]
				}
			}

			// сохраняем собранные данные для последующей передачи в шаблон
			entries = append(entries,
				templateEnumEntry{Name: name, Value: value, Alert: alert})
		}
	}

	// генерируем код по шаблону
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, templateEnum{Package: f.Name.Name, Type: typ, Entries: entries})
	if err != nil {
		panic(err)
	}

	// форматируем код
	bufFmt, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	// записываем сгенерированный код в файл
	basename := strings.TrimSuffix(fname, filepath.Ext(fname))
	err = os.WriteFile(fmt.Sprintf("%s_methods.go", basename), bufFmt, 0644)
	if err != nil {
		panic(err)
	}
}
