package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pedrolfr404/encurtador-url-go/handlers"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	godotenv.Load()
	mongoURI := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Erro ao conectar no MongoDB: %v", err)
	}

	collection := client.Database("encurtador").Collection("encurtador")

	indexModel := mongo.IndexModel{
		Keys: bson.M{"original_url": 1},
	}
	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Printf("Aviso: Não foi possível criar o índice: %v", err)
	}

	urlHandler := &handlers.URLHandler{
		Collection: collection,
		BaseURL:    "http://localhost:8080/",
	}

	r := gin.Default()

	r.POST("/shorten", urlHandler.ShortenURL)
	r.GET("/:id", urlHandler.Redirect)

	log.Println("Servidor rodando em http://localhost:8080")
	r.Run(":8080")
}
