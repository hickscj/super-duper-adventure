package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createDatabase() {
	dbref := 
	return dbref;
}

func main() {
	uri := "mongodb://localhost:27017"
	fmt.Println("HEY MONGO SCRIPTS")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("dex-development").Collection("users")

	filter := bson.D{{Key: "email", Value: "ak-cms-admin@semanticbits.com"}}
	projection := bson.D{}
	opts := options.Find().SetProjection(projection)
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
