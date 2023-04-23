package notion2md

import (
	"testing"

	"github.com/jomei/notionapi"
	"github.com/stretchr/testify/assert"
)

func TestFormatRichText(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    []notionapi.RichText
		expected string
	}{
		"code annotation": {
			input: []notionapi.RichText{
				{
					Annotations: &notionapi.Annotations{
						Code: true,
					},
					PlainText: "code",
				},
			},
			expected: "`code`",
		},
		"bold annotation": {
			input: []notionapi.RichText{
				{
					Annotations: &notionapi.Annotations{
						Bold: true,
					},
					PlainText: "bold",
				},
			},
			expected: "**bold**",
		},
		"italic annotation": {
			input: []notionapi.RichText{
				{
					Annotations: &notionapi.Annotations{
						Italic: true,
					},
					PlainText: "italic",
				},
			},
			expected: "*italic*",
		},
		"strikethrough annotation": {
			input: []notionapi.RichText{
				{
					Annotations: &notionapi.Annotations{
						Strikethrough: true,
					},
					PlainText: "strikethrough",
				},
			},
			expected: "~~strikethrough~~",
		},
		"underline annotation": {
			input: []notionapi.RichText{
				{
					Annotations: &notionapi.Annotations{
						Underline: true,
					},
					PlainText: "underline",
				},
			},
			expected: "<u>underline</u>",
		},
		"with link": {
			input: []notionapi.RichText{
				{
					Href:      "github.com/dstopka",
					PlainText: "github",
				},
			},
			expected: "[github](github.com/dstopka)",
		},
		"multiple annotations at once": {
			input: []notionapi.RichText{
				{
					Annotations: &notionapi.Annotations{
						Bold:          true,
						Italic:        true,
						Strikethrough: true,
					},
					PlainText: "multiple",
				},
			},
			expected: "~~***multiple***~~",
		},
		"annotation with link": {
			input: []notionapi.RichText{
				{
					Annotations: &notionapi.Annotations{
						Italic: true,
					},
					Href:      "example.com",
					PlainText: "link",
				},
			},
			expected: "[*link*](example.com)",
		},
		"multiple words": {
			input: []notionapi.RichText{
				{
					Annotations: &notionapi.Annotations{
						Bold: true,
					},
					PlainText: "Bold beginning",
				},
				{
					PlainText: ", then plain, followed by ",
				},
				{
					Annotations: &notionapi.Annotations{
						Italic: true,
					},
					PlainText: "italic",
				},
				{
					PlainText: " and a ",
				},
				{
					PlainText: "link",
					Href:      "example.com",
				},
			},
			expected: "**Bold beginning**, then plain, followed by *italic* and a [link](example.com)",
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := formatRichText(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}
