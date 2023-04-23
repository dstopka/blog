package notion2md

import (
	"context"

	"github.com/jomei/notionapi"
)

// TemplateData represents data used by go template to build markdown.
type TemplateData struct {
	Blocks []Block
}

func (c Converter) buildTmplData(ctx context.Context, pageID string) (*TemplateData, error) {
	blocks, err := c.buildChildrenBlocks(ctx, notionapi.BlockID(pageID))
	if err != nil {
		return nil, err
	}

	return &TemplateData{
		Blocks: blocks,
	}, nil
}

func (c Converter) getAllChildren(ctx context.Context, blockID notionapi.BlockID) (notionapi.Blocks, error) {
	resp, err := c.client.GetChildren(ctx, blockID, nil)
	if err != nil {
		return nil, err
	}

	blocks := resp.Results
	for resp.HasMore {
		resp, err = c.client.GetChildren(ctx, blockID, &notionapi.Pagination{
			StartCursor: notionapi.Cursor(resp.NextCursor),
		})
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, resp.Results...)
	}

	return blocks, nil
}

func (c Converter) buildChildrenBlocks(ctx context.Context, blockID notionapi.BlockID) ([]Block, error) {
	blocks, err := c.getAllChildren(ctx, blockID)
	if err != nil {
		return nil, err
	}

	tmplBlocks := make([]Block, 0, len(blocks))
	for _, block := range blocks {
		tmplBlock, err := c.buildBlock(ctx, block)
		if err != nil {
			return nil, err
		}
		tmplBlocks = append(tmplBlocks, tmplBlock)
	}

	return tmplBlocks, nil
}

func (c Converter) buildBlock(ctx context.Context, b notionapi.Block) (Block, error) {
	var block Block

	switch notionBlock := b.(type) {
	case *notionapi.BulletedListItemBlock:
		tmplBlock := asBulletedListBlock(notionBlock)
		if b.GetHasChildren() {
			children, err := c.buildChildrenBlocks(ctx, b.GetID())
			if err != nil {
				return nil, err
			}
			tmplBlock.Children = children
		}
		block = tmplBlock
	case *notionapi.CodeBlock:
		block = asCodeBlock(notionBlock)
	case *notionapi.DividerBlock:
		block = asDividerBlock(notionBlock)
	case *notionapi.Heading1Block:
		block = asHeadingBlock(notionBlock)
	case *notionapi.Heading2Block:
		block = asHeadingBlock(notionBlock)
	case *notionapi.Heading3Block:
		block = asHeadingBlock(notionBlock)
	case *notionapi.ImageBlock:
		block = asImageBlock(notionBlock)
	case *notionapi.NumberedListItemBlock:
		tmplBlock := asNumberedListBlock(notionBlock)
		if b.GetHasChildren() {
			children, err := c.buildChildrenBlocks(ctx, b.GetID())
			if err != nil {
				return nil, err
			}
			tmplBlock.Children = children
		}
		block = tmplBlock
	case *notionapi.ToDoBlock:
		tmplBlock := asTodoBlock(notionBlock)
		if b.GetHasChildren() {
			children, err := c.buildChildrenBlocks(ctx, b.GetID())
			if err != nil {
				return nil, err
			}
			tmplBlock.Children = children
		}
		block = tmplBlock
	case *notionapi.ToggleBlock:
		tmplBlock := asToggleBlock(notionBlock)
		if b.GetHasChildren() {
			children, err := c.buildChildrenBlocks(ctx, b.GetID())
			if err != nil {
				return nil, err
			}
			tmplBlock.Children = children
		}
		block = tmplBlock
	case *notionapi.ParagraphBlock:
		block = asParagraphBlock(notionBlock)
	case *notionapi.QuoteBlock:
		block = asQuoteBlock(notionBlock)
	default:
		block = &EmptyBlock{}
	}

	return block, nil
}
