package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ledongthuc/pdf"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

var Dirname = "fiches"

func GetLlm() {
	arg := os.Args
	curr, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if arg[1] == "play" {

		entries, err := os.ReadDir(Dirname)

		if err != nil {
			log.Fatal(err)
		}
		for _, e := range entries {
			fmt.Println(e.Name())

			ctx := context.Background()
			llm, err := googleai.New(ctx, googleai.WithAPIKey("AIzaSyBoaJBQnI-Y6Af8E6e1Dqa-dSAHFoZaLC0"))
			if err != nil {
				log.Fatal(err)
			}

			pdf_file, err := os.Open(path.Join(curr, Dirname, e.Name()))
			if err != nil {
				log.Fatal(err)
			}
			pdf_stat, err := os.Stat((path.Join(curr, Dirname, e.Name())))
			if err != nil {
				log.Fatal(err)
			}
			defer pdf_file.Close()

			reader, err := pdf.NewReader(pdf_file, pdf_stat.Size())

			if err != nil {
				log.Fatal(err)
			}
			page := reader.Page(1)
			text, err := page.GetPlainText(nil)

			if err != nil {
				log.Fatal(err)
			}
			prompt := fmt.Sprintf("peux tu extraire les informations et les presenter en format json et en francais avec les clé 'dci' 'effets' 'indications' 'mode_de_prise' %s", text)
			answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(answer)

		}

	}
}
