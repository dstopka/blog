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

	handler := service.NewQueryPostHandler(cfg)
	lambda := infra.NewQueryPostLambda(handler)

	if cfg.IsDev {
		mux := http.NewServeMux()
		mux.Handle("/post", lambda)
		addr := fmt.Sprintf(":%d", cfg.Port)
		http.ListenAndServe(addr, mux)
	} else {
		awslambda.Start(lambda.ServeLambda)
	}
}
