package templates

import (
	"html/template"
	"os"
	"strings"
)

const sampleTemplate = `
This template demonstrates printing a {{ .Variable |
	printf "%#v" }}.
	{{if .Condition}}
	If condition is set, we'll print this
	{{else}}
	Otherwise, we'll print this instead
	{{end}}
	Next we'll iterate over an array of strings:
	{{range $index, $item := .Items}}
	{{$index}}: {{$item}}
	{{end}}
	We can also easily import other functions like
	strings.Split
	then immediately used the array created as a result:
	{{ range $index, $item := split .Words ","}}
	{{$index}}: {{$item}}
	{{end}}
	Blocks are a way to embed templates into one another
	{{ block "block_example" .}}
	No Block defined!
	{{end}}
	{{/*
	This is a way
	to insert a multi-line comment
*/}}
`

const secondTemplate = `
 {{ define "block_example" }}
 {{.OtherVariable}}
 {{end}}
`

// RunTemplate initializes a template and demonstrates a variety of template helper functions
func RunTemplate() error {
	data := struct {
		Condition     bool
		Variable      string
		Items         []string
		Words         string
		OtherVariable string
	}{
		Condition:     true,
		Variable:      "variable",
		Items:         []string{"one", "two", "three"},
		Words:         "another one, another two, another three",
		OtherVariable: "I'm defined in a second template",
	}
	funcamp := template.FuncMap{
		"split": strings.Split,
	}
	// these can also be chained together
	t := template.New("example").Funcs(funcamp)
	// We could use MUst instead to panic on error
	// template.Must(t.Parse(sampleTemplate))
	t, err := t.Parse(sampleTemplate)
	if err != nil {
		return err
	}
	// to demonstrate block we'll create another template
	// by cloning the first tmeplate, then parsing a second
	t2, err := t.Clone()
	if err != nil {
		return err
	}
	t2, err = t2.Parse(secondTemplate)
	if err != nil {
		return err
	}
	// write the template to stdout and populate it
	// with data
	err = t2.Execute(os.Stdout, &data)
	if err != nil {
		return err
	}
	return nil
}
