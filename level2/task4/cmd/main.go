package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"task4/controller"
)

func main() {
	control:= controller.New()

	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/create/post", control.PostController().CreatePost).Methods("POST")
	router.HandleFunc("/get/post/{id}", control.PostController().GetPost).Methods("GET")
	router.HandleFunc("/update/post/{id}", control.PostController().UpdatePost).Methods("PUT")
	router.HandleFunc("/delete/post/{id}", control.PostController().DeletePost).Methods("DELETE")

	router.HandleFunc("/create/comment", control.CommentController().CreateComment).Methods("POST")
	router.HandleFunc("/get/comment/{id}", control.CommentController().GetComment).Methods("GET")
	router.HandleFunc("/update/comment/{id}", control.CommentController().UpdateComment).Methods("PUT")
	router.HandleFunc("/delete/comment/{id}", control.CommentController().DeleteComment).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", router))
}
