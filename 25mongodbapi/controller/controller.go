package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// used in main file only cuz just to learn
const connectionString = "mongodb+srv://madhavv8528:KOPtQpI2361QNcO7@cluster0.knn998u.mongodb.net/"
const db_name = "cinema"
const col_name = "movies"

var collection *mongo.Collection

func init() {

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected!")

	collection = client.Database(db_name).Collection(col_name)
	fmt.Println("Collection is ready.")
}
