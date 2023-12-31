package controllers

import (
	"encoding/json"
	"fmt"
	m "forum/models"
	"net/http"
	"strconv"
)

type SessionStatusResponse struct {
	LoggedIn bool `json:"loggedIn"`
}

func CheckSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		response := SessionStatusResponse{
			LoggedIn: false,
		}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println("cant marshal the response into json")

		}
		return
	}

	sessionId := cookie.Value

	_, loggedIn, err := m.SessionIsActive(sessionId)

	if err != nil {
		http.Error(w, "unathorized: invalid sesh id", http.StatusUnauthorized)
		fmt.Println(err, "invalid sesh id")
		return
	}

	response := SessionStatusResponse{
		LoggedIn: loggedIn,
	}
	fmt.Println(response, "this is the response u nkeoeoffffffffffffffffffffffeoeo")

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("cant marshal the response into json")

	}

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad req", http.StatusBadRequest)
		return
	}
	var err error
	var postId int
	fmt.Println("making post began")
	user, err := m.GetUserByCookie(r)
	if err != nil {
		http.Error(w, "user has no cookie", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	file, fileHeader, err := r.FormFile("image")

	if file == nil {
		err = nil
	}
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	if fileHeader != nil {
		if fileHeader.Size > 20<<20 { // 20MB
			http.Error(w, "400 Bad Request, File size greater than 20MB", http.StatusBadRequest)
		}
	}
	categories := r.Form["category"]
	if content == "" || title == "" {
		http.Error(w, "can't create an empty post", http.StatusBadRequest)
	}
	fmt.Println(categories)
	ids := m.GetCategoriesID(categories)
	if file == nil {
		postId, err = m.SavePost(title, content, nil, user.ID)
	} else {
		postId, err = m.SavePost(title, content, file, user.ID)
	}
	if err != nil {
		fmt.Println("couldnt save post", err)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	m.LinkPostCategories(postId, ids)
}

type LikeReq struct {
	PostId string `json:"postId"`
}

func checkCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil

}

func LikePost(w http.ResponseWriter, r *http.Request) {
	// check if client has a cookie
	sessionId, _ := checkCookie(w, r)
	userId, session, err1 := m.SessionIsActive(sessionId)
	if err1 != nil {
		fmt.Println("post.controller.go Error func LikePost: ", err1)
		return
	}

	if session {
		LikeReqData := new(LikeReq)
		json.NewDecoder(r.Body).Decode(&LikeReqData)
		postIdStr := LikeReqData.PostId
		fmt.Println("Liked post id is: ", postIdStr)
		postId, _ := strconv.Atoi(postIdStr)

		m.SaveLike(postId, userId)

	}

	// path := path.Base(r.URL.Path)

}
