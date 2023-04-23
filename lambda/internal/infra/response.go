package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dstopka/blog/lambda/internal/app"
)

func httpErrorResponse(w http.ResponseWriter, r *http.Request, err ErrorResponse) {
	httpResponse(w, r, err, err.httpStatus)
}

func httpResponse(w http.ResponseWriter, _ *http.Request, v any, statusCode int) {
	body := &bytes.Buffer{}
	enc := json.NewEncoder(body)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applicatiob/json")
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, body)
}

func lambdaErrorResponse(err ErrorResponse) *events.APIGatewayProxyResponse {
	return lambdaResponse(err, err.httpStatus)
}

func lambdaResponse(v any, statusCode int) *events.APIGatewayProxyResponse {
	body := &bytes.Buffer{}
	enc := json.NewEncoder(body)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}

	return &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: statusCode,
		Body:       body.String(),
	}
}

func postModelToResponse(postModel *app.Post) Post {
	return Post{
		Title:         postModel.Title,
		Description:   postModel.Description,
		Slug:          postModel.Slug,
		Content:       postModel.Content,
		CoverImageURL: postModel.CoverImageURL,
		Tags:          postModel.Tags,
		PublishedTime: postModel.PublishedTime,
		UpdatedTime:   postModel.UpdatedTime,
	}
}

func postsPageModelToResponse(pageModel *app.PostsPage) PostsPage {
	posts := make([]Post, 0, len(pageModel.Posts))
	for i := range pageModel.Posts {
		posts = append(posts, postModelToResponse(&pageModel.Posts[i]))
	}

	return PostsPage{
		Posts:      posts,
		NextCursor: pageModel.Next,
	}
}
