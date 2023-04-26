package service

import (
	"github.com/dstopka/blog/lambda/internal/adapters"
	"github.com/dstopka/blog/lambda/internal/app"
	"github.com/jomei/notionapi"
)

// PostsHandlersConfig defines functions that need to be provided
// by the config used to provision Post related handlers.
type PostsHandlersConfig interface {
	GetNotionAuthToken() string
	GetPostsDatabaseID() string
}

// NewQueryPostHandler creates a new QueryPostHandler using the provided config.
func NewQueryPostHandler(cfg PostsHandlersConfig) app.QueryPostHandler {
	client := notionapi.NewClient(notionapi.Token(cfg.GetNotionAuthToken()))
	notion := adapters.NewNotion(client, cfg.GetPostsDatabaseID())

	return app.NewQueryPostHandler(notion)
}

// NewQueryAllPostsHandler creates a new QueryAllPostsHandler using the provided config.
func NewQueryAllPostsHandler(cfg PostsHandlersConfig) app.QueryAllPostsHandler {
	client := notionapi.NewClient(notionapi.Token(cfg.GetNotionAuthToken()))
	notion := adapters.NewNotion(client, cfg.GetPostsDatabaseID())

	return app.NewQueryAllPostsHandler(notion)
}
