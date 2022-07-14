package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error
var collection *mongo.Collection

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Panic(err)
	}

	collection = client.Database("baz").Collection("qux")

	fmt.Println("Connected Successfully")

	r := httprouter.New()

	r.GET("/users", handleGet)
	r.GET("/users/:id", handleGetById)
	r.POST("/users/create", handleSave)
	r.POST("/users/delete/:id", handleDelete)

	http.ListenAndServe(":8080", r)
}
