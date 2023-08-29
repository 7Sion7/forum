package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"net/http"
	"strconv"
)

func CheckSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "unathorized to post", http.StatusUnauthorized)
		return
	}
	sessionId := cookie.Value
	loggedIn, err := m.SessionIsActive(sessionId)
	if err != nil {
		http.Error(w, "unathorized: invalid sesh id", http.StatusUnauthorized)
		return
	}
	if loggedIn {

		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func CreatePost(r *http.Request) error {
	user, err := m.GetUserByCookie(r)
	if err != nil {
		return err
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["category"]
	ids := m.GetCategoriesID(categories)
	postId, err := m.SavePost(title, content, user.ID)
	if err != nil {
		return err
	}
	m.LinkPostCategories(postId, ids)
	return nil
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var post m.Post

		var message string

		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		}

		switch post.Username {
		case "edit":
			message, err = post.EditPost()
			if err != nil {
				http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			}
	
		case "delete":
			message, err = post.DeletePost()
			if err != nil {
				http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	} else {
		http.Error(w, "404 Page Not Found", 404)
	}
}



func CommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var Comment m.Comment

		err := json.NewDecoder(r.Body).Decode(&Comment)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		}

		user, err := m.GetUserByCookie(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Comment = m.Comment{UserID: user.ID, PostID: Comment.PostID, Comment: Comment.Comment}

		err = Comment.SaveComment()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Comment received"))
	} else {
		http.Error(w, "404 Page Not Found", 404)
	}
}

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := m.GetUserByCookie(r)
	if err != nil {
		if user != nil {
			return
		}
		http.Error(w, "No active Cookie", 404)
	}
	if user != nil {
		var postActionReqData m.PostActionReq
		err = json.NewDecoder(r.Body).Decode(&postActionReqData)
		if err != nil {
			http.Error(w, "500", 500)
			return
		}
		postId, _ := strconv.Atoi(postActionReqData.PostId)
		commentId, _ := strconv.Atoi(postActionReqData.CommentId)
		action := postActionReqData.Action
		switch action {
		case "like":
			m.SaveLike(commentId, postId, user.ID)
		case "unlike":
			m.RemoveLike(commentId, postId, user.ID)
		case "dislike":
			m.SaveDislike(commentId, postId, user.ID)
		case "removeDislike":
			m.RemoveDislike(commentId, postId, user.ID)
		}
	}
}

func GetPostLikes(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			var postID = r.URL.Query().Get("postID")
			ID, err := strconv.Atoi(postID)
			if err != nil {
				return
			}
			likesData, err := m.GetLikedPosts(ID)
			if err != nil {
				return
			}

			JSsender(w, likesData)
			return
		} else {
			http.Error(w, err.Error(), 500)
		}
	}
	user, err := m.GetUserByCookie(r)
	if err != nil {
		return
	}
	likesData, err := m.GetUserLikedPosts(user.ID)
	if err != nil {
		return
	}

	JSsender(w, likesData)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		var postID = r.URL.Query().Get("postID")

		ID, err := strconv.Atoi(postID)
		if err != nil {
			fmt.Println(err)
			return
		}

		comments, err := m.GetCommentsForPost(ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		JSsender(w, comments)
	}
}
