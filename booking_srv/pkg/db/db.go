package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var SessionCollection *mongo.Collection

type Handler struct {
	SessionCollection *mongo.Collection
}

func Init(url string) Handler {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	db, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		log.Fatalln(err)
	}
	SessionCollection = db.Database("booking").Collection("session")

	return Handler{SessionCollection: SessionCollection}
}
