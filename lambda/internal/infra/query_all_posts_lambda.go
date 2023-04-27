package infra

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dstopka/blog/lambda/internal/app"
)

var _ Lambda = (*QueryAllPostsLambda)(nil)
var _ http.Handler = (*QueryAllPostsLambda)(nil)

// QueryAllPostsLambda allows serving an endpoint for querying all blog posts
// as a lambda or as an http endpoint.
type QueryAllPostsLambda struct {
	handler app.QueryAllPostsHandler
}

// NewQueryAllPostsLambda creates a new QueryAllPostsLambda using the given handler.
func NewQueryAllPostsLambda(handler app.QueryAllPostsHandler) QueryAllPostsLambda {
	return QueryAllPostsLambda{handler: handler}
}

// ServeLambda is an entrypoint for serving an aws lambda.
func (l QueryAllPostsLambda) ServeLambda(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	postsModel, err := l.handler.Handle(ctx)
	if err != nil {
		log.Println(err)
		return lambdaErrorResponse(InternalServerError()), nil
	}

	posts := postsModelToResponse(postsModel)
	return lambdaResponse(posts, http.StatusOK), nil
}

// ServeHTTP is an entrypoint for serving an http endpoint.
func (l QueryAllPostsLambda) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	postsModel, err := l.handler.Handle(r.Context())
	// log.Printf("%#+v\n", postsModel)
	if err != nil {
		httpErrorResponse(w, r, InternalServerError())
		return
	}

	posts := postsModelToResponse(postsModel)
	// log.Printf("%#+v\n", posts)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	httpResponse(w, r, posts, http.StatusOK)
}
