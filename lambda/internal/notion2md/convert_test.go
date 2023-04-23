package notion2md

import (
	"bytes"
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConverter_execute(t *testing.T) {
	t.Parallel()

	testdata := &TemplateData{
		Blocks: []Block{
			&HeadingBlock{
				Text:  "Heading1",
				Level: 1,
			},
			&TodoBlock{
				Text:    "Checked",
				Checked: true,
				Children: []Block{
					&TodoBlock{Text: "Nested todo1"},
					&TodoBlock{Text: "Nested todo2"},
				},
			},
			&TodoBlock{
				Text: "Unchecked",
			},
			&ParagraphBlock{},
			&NumberedListBlock{
				Text: "List1",
				Children: []Block{
					&NumberedListBlock{Text: "List1.1"},
					&NumberedListBlock{Text: "List1.2"},
				},
			},
			&NumberedListBlock{Text: "List2"},
			&ParagraphBlock{},
			&ToggleBlock{Text: "Empty"},
			&ToggleBlock{
				Text: "With content",
				Children: []Block{
					&BulletedListBlock{
						Text: "List1",
						Children: []Block{
							&BulletedListBlock{Text: "List1.1"},
							&BulletedListBlock{Text: "List1.2"},
						},
					},
					&BulletedListBlock{Text: "List2"},
				},
			},
			&ParagraphBlock{},
			&CodeBlock{
				Language: "go",
				Code:     "func main() {\n    fmt.Println(\"Hello World\")\n}",
				Caption:  "",
			},
			&ParagraphBlock{},
			&ImageBlock{
				URL:     "https://s3.us-west-2.amazonaws.com/image",
				Caption: "Img Caption",
			},
			&ParagraphBlock{},
			&QuoteBlock{
				Text: "this is\na multiline quote",
			},
			&DividerBlock{},
			&ParagraphBlock{
				Text: "Text can be **bold**, *italic* or ***both***.",
			},
			&ParagraphBlock{
				Text: "A paragraph that has multiple lines on the mobile screen.",
			},
		},
	}

	conv := Converter{tmpl: template.Must(template.New("notion2markdown").Funcs(tmplFuncs()).Parse(markdownTmpl))}
	md := &bytes.Buffer{}

	err := conv.execute(md, testdata)
	require.NoError(t, err)

	assertEqualToFile(t, "./testdata/expected.md", md.String())
}

func assertEqualToFile(t *testing.T, filePath string, actual string) {
	fileContent, err := os.ReadFile(filePath)
	require.NoError(t, err)

	assert.Equal(t, string(fileContent), actual)
}
