package infra

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dstopka/blog/lambda/internal/app"
)

var _ Lambda = (*QueryPostLambda)(nil)
var _ http.Handler = (*QueryPostLambda)(nil)

// QueryPostLambda allows serving an endpoint for querying a blog post
// as a lambda or as an http endpoint.
type QueryPostLambda struct {
	handler app.QueryPostHandler
}

// NewQueryPostLambda creates a new QueryPostLambda using the given handler.
func NewQueryPostLambda(handler app.QueryPostHandler) QueryPostLambda {
	return QueryPostLambda{handler: handler}
}

// ServeLambda is an entrypoint for serving an aws lambda.
func (l QueryPostLambda) ServeLambda(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	slug, ok := req.QueryStringParameters[QueryParamSlug]
	if !ok {
		return lambdaErrorResponse(BadRequest("Query argument 'slug' is required, but not found")), nil
	}

	postModel, err := l.handler.Handle(ctx, app.QueryPost{Slug: slug})
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			return lambdaErrorResponse(NotFound()), nil
		}
		log.Println(err)
		return lambdaErrorResponse(InternalServerError()), nil
	}

	post := postModelToResponse(postModel)

	resp := lambdaResponse(post, http.StatusOK)
	resp.Headers["Cache-Control"] = "public, max-age=3600"

	return resp, nil
}

// ServeHTTP is an entrypoint for serving an http endpoint.
func (l QueryPostLambda) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get(QueryParamSlug)
	if slug == "" {
		httpErrorResponse(w, r, BadRequest("Query argument 'slug' is required, but not found"))
		return
	}

	postModel, err := l.handler.Handle(r.Context(), app.QueryPost{Slug: slug})
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			httpErrorResponse(w, r, NotFound())
			return
		}
		log.Println(err)
		httpErrorResponse(w, r, InternalServerError())
		return
	}

	post := postModelToResponse(postModel)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("cache-control", "public, max-age=3600")

	httpResponse(w, r, post, http.StatusOK)
}
