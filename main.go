package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type User struct {
	fullName string
	userName string
	email    string
}

type Post struct {
	title  string
	body   string
	author User
}

var data []Post

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/post", getPost).Methods("GET")
	router.HandleFunc("/post", addPost).Methods("POST")
	router.HandleFunc("/post/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")
	http.ListenAndServe(":8080", router) // Runs the Application at Port 8080
}

// Method for getPost Request
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(data) // Sends Data in JSON Format
}

// Method for addPost Request
func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost) // Decoding new Post from Http Request
	data = append(data, newPost)
	json.NewEncoder(w).Encode(newPost) // Sends Data in JSON Format
}

// Method for updatePost Request
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam) // Converting String into Integer
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID failed to convert into Integer"))
		return
	}
	if id >= len(data) {
		w.WriteHeader(400)
		w.Write([]byte("No such data found"))
		return
	}
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost) // Decoding new Post from Http Request
	data[id] = newPost
	json.NewEncoder(w).Encode(newPost) // Sends Data in JSON Format
}

// Method fot deletePost Request
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/josn")
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID failed to convert into Integer"))
		return
	}
	if id >= len(data) {
		w.WriteHeader(400)
		w.Write([]byte("No such data found"))
		return
	}
	data = append(data[:id], data[id+1:]...)
	w.WriteHeader(200)
}
