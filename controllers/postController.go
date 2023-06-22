package controllers

import (
	"forum/models"
	"log"
	"net/http"
	"strconv"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	err = Tpl.ExecuteTemplate(w, "create_post.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Likes(w http.ResponseWriter, r *http.Request) {
	pID := r.URL.Query().Get("postID")
	postID, err := strconv.Atoi(pID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	likes, err := models.GetLikes(postID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	likes++
}