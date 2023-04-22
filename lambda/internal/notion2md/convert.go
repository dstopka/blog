package notion2md

import (
	"context"
	_ "embed"
	"errors"
	"html/template"
	"io"
)

//go:embed templates/main.tmpl
var markdownTmpl string

// Converter fetches notion page data and converts it to markdown using go templates.
type Converter struct {
	client NotionChildrenGetter
	tmpl   *template.Template
}

// NewConverter returns a new Converter using client.
func NewConverter(client NotionChildrenGetter) (*Converter, error) {
	if client == nil {
		return nil, errors.New("client cannot be nil")
	}

	tmpl := template.Must(template.New("notion2markdown").Funcs(tmplFuncs()).Parse(markdownTmpl))
	return &Converter{client: client, tmpl: tmpl}, nil
}

// ConvertContext converts notion page under the given pageID to markdown and writes the result to wr.
func (c Converter) ConvertContext(ctx context.Context, pageID string, wr io.Writer) error {
	data, err := c.buildTmplData(ctx, pageID)
	if err != nil {
		return err
	}

	return c.execute(wr, data)
}

func (c Converter) execute(wr io.Writer, data *TemplateData) error {
	return c.tmpl.Execute(wr, data)
}

func tmplFuncs() template.FuncMap {
	return template.FuncMap{
		"loop": func(n int) []int {
			return make([]int, n)
		},
		"noescape": func(s string) template.HTML {
			return template.HTML(s)
		},
	}
}
