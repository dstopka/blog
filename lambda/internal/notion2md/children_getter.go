package notion2md

import (
	"context"

	"github.com/jomei/notionapi"
)

// NotionChildrenGetter defines functions needed by Converter to fetch the data.
//
//go:generate go run github.com/vektra/mockery/v2 --name NotionChildrenGetter
type NotionChildrenGetter interface {
	GetChildren(context.Context, notionapi.BlockID, *notionapi.Pagination) (*notionapi.GetChildrenResponse, error)
}