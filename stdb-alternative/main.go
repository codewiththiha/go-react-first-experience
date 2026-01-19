package stdbalternative

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TODO struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error Loading .env File")
		}
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		log.Fatal("Error connecting to mongo DB")
	}

	defer client.Disconnect(context.Background())
	collection = client.Database("golang_db").Collection("todos")
	fmt.Println("Connected to DB")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/todos", getTodos)
	mux.HandleFunc("POST /api/todos", createTodo)
	mux.HandleFunc("PATCH /api/todos/{id}", updateTodo)
	mux.HandleFunc("DELETE /api/todos/{id}", deleteTodo)

	log.Fatal(http.ListenAndServe(":4000", mux))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	var todos []TODO
	cursor, err := collection.Find(r.Context(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// it's decode and put the results into the todos "&" is like a key
	if err := cursor.All(r.Context(), &todos); err != nil {
		http.Error(w, err.Error(), 500)
		return

	}
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo TODO
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid input", 400)
		return
	}

	// we are using go mongo driver which only accept go's struct not jsons so we need to decode back into go struct from json
	insertBody, err := collection.InsertOne(r.Context(), newTodo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id, ok := insertBody.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Fatal("Type Mismatched")
	}

	newTodo.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": true}}
	collection.UpdateOne(context.Background(), filter, update)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	filter := bson.M{"_id": objectId}
	collection.DeleteOne(r.Context(), filter)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
