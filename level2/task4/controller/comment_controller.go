package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"task4/models"
)

type CommentController struct {
	controller *Controller
}

func (commentController *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	comment := &models.Comment{}
	if err := json.Unmarshal(requestBody, comment); err != nil {
		log.Fatal(err)
	}

	commentController.controller.DB.GormDB.Create(comment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func (commentController *CommentController) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	comment := &models.Comment{}
	commentController.controller.DB.GormDB.First(comment, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

func (commentController *CommentController) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)
	comment := &models.Comment{}
	json.Unmarshal(requestBody, comment)
	comment.ID, _ = strconv.Atoi(key)
	commentController.controller.DB.GormDB.Save(comment)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

func (commentController *CommentController) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	comment := &models.Comment{}
	id, _ := strconv.ParseInt(key, 10, 64)
	commentController.controller.DB.GormDB.Where("id = ?", id).Delete(comment)
	w.WriteHeader(http.StatusNoContent)
}
