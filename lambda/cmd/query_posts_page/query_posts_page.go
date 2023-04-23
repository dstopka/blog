package main

import (
	"fmt"
	"log"
	"net/http"

	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/dstopka/blog/lambda/internal/env"
	"github.com/dstopka/blog/lambda/internal/infra"
	"github.com/dstopka/blog/lambda/internal/service"
)

func main() {
	cfg, err := env.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	handler := service.NewQueryPostsPageHandler(cfg)
	lambda := infra.NewQueryPostsPageLambda(handler)

	if cfg.IsDev {
		addr := fmt.Sprintf(":%d", cfg.Port)
		http.ListenAndServe(addr, lambda)
	} else {
		awslambda.Start(lambda.ServeLambda)
	}
}
