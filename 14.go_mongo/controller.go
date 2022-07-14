package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var user User
var users []User

func handleGetById(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Panic(err)
	}

	filter := bson.D{{Key: "_id", Value: oid}}
	opts := options.FindOne().SetSort(bson.D{{Key: "age", Value: 1}})
	err = collection.FindOne(context.Background(), filter, opts).Decode(&user)

	if err != nil {
		fmt.Fprintln(w, "No documents in result")
		return
	}

	fmt.Fprintln(w, user)
}

func handleGet(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	if err = cur.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, users)

}

func handleSave(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		log.Panic(err)
	}

	data, err := bson.Marshal(user)

	if err != nil {
		log.Panic(err)
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Panic(err)
	}
	id := res.InsertedID
	fmt.Println("Inserted id ", id)
}

func handleDelete(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Panic(err)
	}

	filter := bson.D{{Key: "_id", Value: oid}}
	opts := options.FindOneAndDelete().
		SetProjection(bson.D{{Key: "firstname", Value: 1}, {Key: "lastname", Value: 1}, {Key: "age", Value: 1}})

	err = collection.FindOneAndDelete(context.Background(), filter, opts).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Println(err)
	}

	fmt.Fprintln(w, "User deleted ", user)
}
