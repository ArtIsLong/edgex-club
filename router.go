package main

import (
	controller "edgex-club/controller"
	"net/http"

	mux "github.com/gorilla/mux"
)

func initRouter() http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/ping", ping).Methods("GET")
	s.HandleFunc("/login", controller.Login).Methods("POST", "GET")
	// s.HandleFunc("/register",controller.Register).Methods("POST")
	s.HandleFunc("/loginbygithub", controller.LoginByGitHub).Methods("POST", "GET")
	s.HandleFunc("/isvalid/{token}", controller.ValidToken).Methods("GET")

	//======================简易消息API========================================
	//跟新消息状态，已读、未读
	s.HandleFunc("/message/{id}", controller.UpdateMessage).Methods("PUT")
	//find 所有消息
	s.HandleFunc("/message", controller.FindAllMessage).Methods("GET")
	//未读消息总数
	s.HandleFunc("/message/count", controller.FindAllMessageCount).Methods("GET")

	//======================回复API============================================
	//find一个文章的所有评论
	s.HandleFunc("/comments/{articleId}", controller.FindAllCommentByArticleId).Methods("GET")
	//find一个评论下的所有回复
	s.HandleFunc("/replys/{commentId}", controller.FindAllReplyByCommentId).Methods("GET")
	//回复某个评论
	s.HandleFunc("/reply/{commentId}/{toUserName}", controller.PostReply).Methods("POST")
	//给一个文章发表评论
	s.HandleFunc("/comment/{articleId}", controller.PostComment).Methods("POST")

	//======================文章API============================================
	//更新文章
	s.HandleFunc("/updateArticle/{articleId}", controller.UpdateArticle).Methods("POST")
	//发表文章
	s.HandleFunc("/{userId}/saveNewArticle", controller.SaveNewArticle).Methods("POST")
	//首页加载最新最热门的文章
	s.HandleFunc("/findNewArticles", controller.FindNewArticles).Methods("GET")

	s.HandleFunc("/hotAuthor", controller.HotAuthor).Methods("GET")
	s.HandleFunc("/hotArticle", controller.HotArticle).Methods("GET")

	s1 := r.PathPrefix("").Subrouter()
	//阅读某个用户的文章
	s1.HandleFunc("/articles/users/{userName}/{articleId}", controller.FindOneArticle).Methods("GET")
	s1.HandleFunc("/articles/users/{userId}", controller.FindAllArticlesByUser).Methods("GET")
	//用户主页
	s1.HandleFunc("/users/{userName}", controller.UserHome).Methods("GET")
	//加载编辑、发帖模板【这个地方需要用户登录认证后才能操作，但是在URL中不能显示/api/v1的路径，因此在此处单独处理】
	s1.HandleFunc("/article/edit/{userName}/{articleId}", controller.LoadEditArticleTemplate).Methods("GET")

	return r
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte("pong"))
}
