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

type PostController struct {
	controller *Controller
}

func (postController *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	post := &models.Post{}
	if err := json.Unmarshal(requestBody, post); err != nil {
		log.Fatal(err)
	}

	postController.controller.DB.GormDB.Create(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (postController *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	post := &models.Post{}
	postController.controller.DB.GormDB.First(post, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (postController *PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	requestBody, _ := ioutil.ReadAll(r.Body)
	post := &models.Post{}
	json.Unmarshal(requestBody, post)
	post.ID, _ = strconv.Atoi(key)
	postController.controller.DB.GormDB.Save(post)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (postController *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	post := &models.Post{}
	id, _ := strconv.ParseInt(key, 10, 64)
	postController.controller.DB.GormDB.Where("id = ?", id).Delete(post)
	w.WriteHeader(http.StatusNoContent)
}
