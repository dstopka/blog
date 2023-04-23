package notion2md

import (
	"context"
	"testing"

	"github.com/dstopka/blog/lambda/internal/notion2md/mocks"
	"github.com/google/uuid"
	"github.com/jomei/notionapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestConverter_getAllChildren_success(t *testing.T) {
	t.Parallel()

	expectedChildren := make(notionapi.Blocks, 0)
	blockID := notionapi.BlockID(uuid.NewString())

	mockClient := mocks.NewNotionChildrenGetter(t)
	mockClient.On("GetChildren", mock.Anything, blockID, (*notionapi.Pagination)(nil)).Return(
		func(_ context.Context, _ notionapi.BlockID, _ *notionapi.Pagination) (*notionapi.GetChildrenResponse, error) {
			childrenBlocks := notionapi.Blocks{
				&notionapi.Heading1Block{
					Heading1: notionapi.Heading{
						RichText: []notionapi.RichText{{PlainText: "Heading1"}},
					},
				},
				&notionapi.ParagraphBlock{
					Paragraph: notionapi.Paragraph{
						RichText: []notionapi.RichText{{PlainText: "First paragraph"}},
					},
				},
			}

			expectedChildren = append(expectedChildren, childrenBlocks...)

			return &notionapi.GetChildrenResponse{Results: childrenBlocks, HasMore: true, NextCursor: "next-cursor-1"}, nil
		},
	).Once()

	mockClient.On("GetChildren", mock.Anything, blockID, &notionapi.Pagination{
		StartCursor: notionapi.Cursor("next-cursor-1"),
	}).Return(func(
		_ context.Context,
		_ notionapi.BlockID,
		_ *notionapi.Pagination,
	) (*notionapi.GetChildrenResponse, error) {
		childrenBlocks := notionapi.Blocks{
			&notionapi.ParagraphBlock{
				Paragraph: notionapi.Paragraph{
					RichText: []notionapi.RichText{{PlainText: "Second paragraph"}},
				},
			},
			&notionapi.ImageBlock{
				Image: notionapi.Image{
					File: &notionapi.FileObject{URL: "https://s3.us-west-2.amazonaws.com/image"},
				},
			},
		}

		expectedChildren = append(expectedChildren, childrenBlocks...)

		return &notionapi.GetChildrenResponse{Results: childrenBlocks, HasMore: true, NextCursor: "next-cursor-2"}, nil
	}).Once()

	mockClient.On("GetChildren", mock.Anything, blockID, &notionapi.Pagination{
		StartCursor: notionapi.Cursor("next-cursor-2"),
	}).Return(func(
		_ context.Context,
		_ notionapi.BlockID,
		_ *notionapi.Pagination,
	) (*notionapi.GetChildrenResponse, error) {
		childrenBlocks := notionapi.Blocks{
			&notionapi.ParagraphBlock{
				Paragraph: notionapi.Paragraph{
					RichText: []notionapi.RichText{{PlainText: "Third paragraph"}},
				},
			},
		}

		expectedChildren = append(expectedChildren, childrenBlocks...)

		return &notionapi.GetChildrenResponse{Results: childrenBlocks, HasMore: false}, nil
	}).Once()

	c := Converter{client: mockClient}
	res, err := c.getAllChildren(context.Background(), blockID)

	require.NoError(t, err)
	assert.Equal(t, expectedChildren, res)
}
