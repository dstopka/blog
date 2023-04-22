package notion2md

import (
	"fmt"

	"github.com/jomei/notionapi"
)

type Block interface {
	Type() BlockType
}

type BlockType string

type BulletedListBlock struct {
	Text     string
	Children []Block
}

func (blb BulletedListBlock) Type() BlockType {
	return "bulletedList"
}

func asBulletedListBlock(b *notionapi.BulletedListItemBlock) *BulletedListBlock {
	return &BulletedListBlock{
		Text: formatRichText(b.BulletedListItem.RichText),
	}
}

type CodeBlock struct {
	Language string
	Code     string
	Caption  string
}

func (cb CodeBlock) Type() BlockType {
	return "code"
}

func asCodeBlock(b *notionapi.CodeBlock) *CodeBlock {
	return &CodeBlock{
		Language: b.Code.Language,
		Caption:  formatRichText(b.Code.Caption),
		Code:     formatRichText(b.Code.RichText),
	}
}

type DividerBlock struct{}

func (db DividerBlock) Type() BlockType {
	return "divider"
}

func asDividerBlock(_ *notionapi.DividerBlock) *DividerBlock {
	return &DividerBlock{}
}

type EmptyBlock struct{}

func (eb EmptyBlock) Type() BlockType {
	return "empty"
}

type Heading interface {
	*notionapi.Heading1Block | *notionapi.Heading2Block | *notionapi.Heading3Block
}

type HeadingBlock struct {
	Text  string
	Level int
}

func (hb HeadingBlock) Type() BlockType {
	return "heading"
}

func asHeadingBlock[H Heading](b H) *HeadingBlock {
	var hb *HeadingBlock
	
	switch v := any(b).(type) {
	case *notionapi.Heading1Block:
		hb = &HeadingBlock{
			Text:  formatRichText(v.Heading1.RichText),
			Level: 1,
		}
	case *notionapi.Heading2Block:
		hb = &HeadingBlock{
			Text:  formatRichText(v.Heading2.RichText),
			Level: 2,
		}
	case *notionapi.Heading3Block:
		hb = &HeadingBlock{
			Text:  formatRichText(v.Heading3.RichText),
			Level: 3,
		}
	default:
		panic("unsupported type")
	}

	return hb
}

type ImageBlock struct {
	URL     string
	Caption string
}

func (ib ImageBlock) Type() BlockType {
	return "image"
}

func asImageBlock(b *notionapi.ImageBlock) *ImageBlock {
	return &ImageBlock{
		URL:     b.Image.GetURL(),
		Caption: formatRichText(b.Image.Caption),
	}
}

type NumberedListBlock struct {
	Text     string
	Children []Block
}

func (nlb NumberedListBlock) Type() BlockType {
	return "numberedList"
}

func asNumberedListBlock(b *notionapi.NumberedListItemBlock) *NumberedListBlock {
	return &NumberedListBlock{
		Text: formatRichText(b.NumberedListItem.RichText),
	}
}

type ParagraphBlock struct {
	Text string
}

func (pb ParagraphBlock) Type() BlockType {
	return "paragraph"
}

func asParagraphBlock(b *notionapi.ParagraphBlock) *ParagraphBlock {
	return &ParagraphBlock{
		Text: formatRichText(b.Paragraph.RichText),
	}
}

type QuoteBlock struct {
	Text string
}

func (qb QuoteBlock) Type() BlockType {
	return "quote"
}

func asQuoteBlock(b *notionapi.QuoteBlock) *QuoteBlock {
	return &QuoteBlock{
		Text: formatRichText(b.Quote.RichText),
	}
}

type TodoBlock struct {
	Checked  bool
	Text     string
	Children []Block
}

func (tb TodoBlock) Type() BlockType {
	return "todo"
}

func asTodoBlock(b *notionapi.ToDoBlock) *TodoBlock {
	return &TodoBlock{
		Checked: b.ToDo.Checked,
		Text:    formatRichText(b.ToDo.RichText),
	}
}

type ToggleBlock struct {
	Text     string
	Children []Block
}

func (tb ToggleBlock) Type() BlockType {
	return "toggle"
}

func asToggleBlock(b *notionapi.ToggleBlock) *ToggleBlock {
	return &ToggleBlock{
		Text: formatRichText(b.Toggle.RichText),
	}
}

func formatRichText(rts []notionapi.RichText) string {
	var text string

	for _, rt := range rts {
		textPart := rt.PlainText
		if rt.Annotations != nil {
			if rt.Annotations.Code {
				textPart = "`" + textPart + "`"
			}
			if rt.Annotations.Bold {
				textPart = "**" + textPart + "**"
			}
			if rt.Annotations.Italic {
				textPart = "*" + textPart + "*"
			}
			if rt.Annotations.Strikethrough {
				textPart = "~~" + textPart + "~~"
			}
			if rt.Annotations.Underline {
				textPart = "<u>" + textPart + "</u>"
			}
		}

		if rt.Href != "" {
			textPart = fmt.Sprintf("[%s](%s)", textPart, rt.Href)
		}

		text += textPart
	}

	return text
}
