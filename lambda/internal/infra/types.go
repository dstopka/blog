package infra

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

const (
	QueryParamSlug     = "slug"
)

type Lambda interface {
	ServeLambda(context.Context, events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse
}

type Post struct {
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Slug          string     `json:"slug"`
	Content       string     `json:"content,omitempty"`
	CoverImageURL string     `json:"coverImageURL"`
	Tags          []string   `json:"tags,omitempty"`
	PublishedTime time.Time  `json:"publishedTime"`
	UpdatedTime   *time.Time `json:"updatedTime,omitempty"`
}

type Posts struct {
	Posts      []Post `json:"posts"`
}

type ErrorResponse struct {
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
	httpStatus  int
}

func BadRequest(desc string) ErrorResponse {
	return ErrorResponse{
		Slug:        "bad-request",
		Description: desc,
		httpStatus:  http.StatusBadRequest,
	}
}

func InternalServerError() ErrorResponse {
	return ErrorResponse{
		Slug:       "internal-server-error",
		httpStatus: http.StatusInternalServerError,
	}
}

func NotFound() ErrorResponse {
	return ErrorResponse{
		Slug:       "not-found",
		httpStatus: http.StatusNotFound,
	}
}
