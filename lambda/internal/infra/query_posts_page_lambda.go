package infra

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dstopka/blog/lambda/internal/app"
)

var _ Lambda = (*QueryPostsPageLambda)(nil)
var _ http.Handler = (*QueryPostsPageLambda)(nil)

// QueryPostsPageLambda allows serving an endpoint for querying a blog posts page
// as a lambda or as an http endpoint.
type QueryPostsPageLambda struct {
	handler app.QueryPostsPageHandler
}

// NewQueryPostsPageLambda creates a new QueryPostsPageLambda using the given handler.
func NewQueryPostsPageLambda(handler app.QueryPostsPageHandler) QueryPostsPageLambda {
	return QueryPostsPageLambda{handler: handler}
}

// ServeLambda is an entrypoint for serving an aws lambda.
func (l QueryPostsPageLambda) ServeLambda(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) *events.APIGatewayProxyResponse {
	query := app.QueryPostsPage{
		Cursor: req.QueryStringParameters[QueryParamCursor],
	}

	pageSizeStr := req.QueryStringParameters[QueryParamPageSize]
	if pageSizeStr != "" {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			return lambdaErrorResponse(BadRequest(fmt.Sprintf("Invalid format for parameter 'pageSize': %s", pageSizeStr)))
		}
		query.PageSize = pageSize
	}

	pageModel, err := l.handler.Handle(ctx, query)
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			return lambdaErrorResponse(NotFound())
		}
		log.Println(err)
		return lambdaErrorResponse(InternalServerError())
	}

	page := postsPageModelToResponse(pageModel)
	return lambdaResponse(page, http.StatusOK)
}

// ServeHTTP is an entrypoint for serving an http endpoint.
func (l QueryPostsPageLambda) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := app.QueryPostsPage{
		Cursor: r.URL.Query().Get(QueryParamCursor),
	}

	if r.URL.Query().Has(QueryParamPageSize) {
		pageSizeStr := r.URL.Query().Get(QueryParamPageSize)
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			httpErrorResponse(w, r, BadRequest(fmt.Sprintf("Invalid format for parameter 'pageSize': %s", pageSizeStr)))
			return
		}
		query.PageSize = pageSize
	}

	pageModel, err := l.handler.Handle(r.Context(), query)
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			httpErrorResponse(w, r, NotFound())
			return
		}
		httpErrorResponse(w, r, InternalServerError())
		return
	}

	page := postsPageModelToResponse(pageModel)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	httpResponse(w, r, page, http.StatusOK)
}
