package main

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms/googleai"
)

var Dirname = "fiches"

func GetLlm(apikey string) *googleai.GoogleAI {

	ctx := context.Background()

	llm, err := googleai.New(ctx, googleai.WithAPIKey(apikey))

	if err != nil {
		log.Fatal(err)
	}

	return llm

}
