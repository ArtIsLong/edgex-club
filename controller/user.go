package controller

import (
	model "edgex-club/model"
	respos "edgex-club/repository"
	"encoding/json"
	mux "github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type TodoPageUserData struct {
	UserId       string
	UserName     string
	AvatarUrl    string
	ArticleCount int
}

//Register methond
func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	userId, err := respos.UserRepos.Insert(u)
	log.Println("userId : " + userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte(userId))
}

func UserHome(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	userName := vars["userName"]
	u := respos.UserRepos.FindOneByName(userName)
	log.Println("userHome===")
	log.Println(u)
	articleCount := respos.ArticleRepos.UserArticleCount(userName)

	t, _ := template.ParseFiles("static/users/home.html")
	data := TodoPageUserData{
		UserId:       u.Id.Hex(),
		UserName:     userName,
		AvatarUrl:    u.AvatarUrl,
		ArticleCount: articleCount,
	}
	t.Execute(w, data)
}
