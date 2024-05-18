package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Stupnikjs/pharmago/pkg/repo"
	"github.com/joho/godotenv"
)

type application struct {
	config map[string]string
	URI_DB string
	DB     repo.DatabaseRepo
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := make(map[string]string)
	config["SECRET_TOKEN"] = os.Getenv("SECRET_TOKEN")

	app := application{}

	app.config = config

	/* conn, err := app.connectToDB()

	if err != nil {
		fmt.Println("error connection to db")
	}
	// app.DB = &db.PostgresDBRepo{DB: conn}

	// app.InitTables()

	*/

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, app.routes())
}
