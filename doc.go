package design

import (
	"io"
	"os"
	"reflect"
	"text/template"
)

type Document struct {
	Parts Stringers
	Style Stringer
}

func NewDocument() *Document {
	return &Document{}
}

func (doc *Document) Editor() Editor {
	return doc.edit
}

type Editor func(...interface{})

func (doc *Document) edit(arguments ...interface{}) {
	for _, arg := range arguments {
		doc.appendByType(arg)
	}
}

func (doc *Document) appendByType(arg interface{}) {
	var valid Stringer
	switch a := arg.(type) {
	case string:
		valid = Plain(a)
	case *Graph:
		doc.Parts = append(doc.Parts, Plain("\n"))
		valid = a
	case Stringer:
		valid = a
	default:
		valid = Plain(reflect.TypeOf(arg).Name())
	}
	doc.Parts = append(doc.Parts, valid)
}

func (doc *Document) SaveAs(filename string) error {
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	_, err = doc.WriteTo(fh)
	return err
}

func (doc *Document) WriteTo(w io.Writer) (int, error) {
	r, writer := io.Pipe()
	go func() {
		tpl := template.Must(template.New("html").Parse(htmlSource))
		tpl.Execute(writer, doc)
		writer.Close()
	}()
	n, err := io.Copy(w, r)
	// Safe for most cases, we're dealing with small documents
	return int(n), err
}

const htmlSource = `<!DOCTYPE html>

<html>
  <head>
    <style>
      .component, .smallbox {
        fill:#ffffcc;
        stroke:black;
        stroke-width:1;
      }
      line {
        stroke:black;
        stroke-width:1;
      }
      {{.Style}}
    </style>
  </head>
<body>
{{range .Parts}}{{.}}{{end}}
</body>
</html>`

type Stringers []Stringer

type Stringer interface {
	String() string
}

type StringerFunc func() string

func (fn StringerFunc) String() string {
	return fn()
}

type Plain string

func (p Plain) String() string {
	return string(p)
}
