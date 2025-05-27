package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodbapi/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//helper function

func addAMovie(movie model.Cinema) {

	insertedMovie, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie inserted with uniqueId", insertedMovie.InsertedID)
}

func updateAMovie(id string) {

	ObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": ObjectId}
	update := bson.M{"$set": bson.M{"watched": true}}

	updatedDoc, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Doc modified count", updatedDoc.ModifiedCount)
}

func deleteOneMovie(id string) {

	ObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	options := bson.M{"_id": ObjectId}

	deletedDoc, err := collection.DeleteOne(context.Background(), options)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted document", deletedDoc)
}

func deleteAll() int64 {
	//bson.M does'nt support empty map
	option := bson.D{{}}
	deletedCol, err := collection.DeleteMany(context.Background(), option, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("No. of deleted movie", deletedCol.DeletedCount)
	return deletedCol.DeletedCount
}

func getAllMovies() []primitive.M {

	filter := bson.D{{}}

	cursor, err := collection.Find(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	fmt.Println("List of all movies is: ")
	return movies
}

//controllers

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode("<h1>Welcome to Cinema</h1>")
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	movies := getAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func CreateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	//just to try new header
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Cinema
	json.NewDecoder(r.Body).Decode(&movie)

	addAMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func UpdateAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateAMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAll()

	json.NewEncoder(w).Encode(count)
}
