package router

import (
	"encoding/json"
	"kanolibrary/crud"
	"kanolibrary/models"
	"kanolibrary/util"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := crud.FindAll("books")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		requestId := r.URL.Query().Get("id")

		objID, err := primitive.ObjectIDFromHex(requestId)
		result, err := crud.FindOne("books", bson.M{"_id": objID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func InsertBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := crud.Insert("books", book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PATCH" {
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		util.ErrorChecker(err)
		result, err := crud.UpdateOne("books", bson.M{"_id": objId}, bson.M{"$set": bson.M{"title": book.Title, "author": book.Author, "synopsis": book.Synopsis, "UpdateAt": primitive.NewDateTimeFromTime(time.Now())}})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		result, err := crud.DeleteOne("books", bson.M{"_id": objId})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
