package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

var (
	client          *mongo.Client
	usersCollection *mongo.Collection
)

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	// create testing database and users collection
	usersCollection = client.Database("testing").Collection("users")
}

func insert() {
	//  insert one user
	one, err := usersCollection.InsertOne(context.TODO(), bson.M{"name": "John", "age": 25})
	if err != nil {
		panic(err)
	}
	log.Println("insert one id: ", one.InsertedID)

	two, err := usersCollection.InsertOne(context.TODO(), User{Name: "ray", Age: 40})
	if err != nil {
		panic(err)
	}
	log.Println("insert two id: ", two.InsertedID)
}

func find() {
	cur, err := usersCollection.Find(context.TODO(), bson.M{"name": "John"})
	if err != nil {
		panic(err)
	}

	var results []User
	err = cur.All(context.TODO(), &results)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		data, _ := json.Marshal(result)
		log.Println(string(data))
	}
}

func main() {
	find()
}
