package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/emilyyf/friendly-backend/internal/handlers"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MODGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
		// log.Fatal("Set your 'MONDOGB_URI' environment variable. ")
	}

	db, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	mux := http.NewServeMux()

	// mux.HandleFunc("/login", handlers.NewPostLoginHandler())
}
