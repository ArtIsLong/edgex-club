package main

import (
	controller "edgex-club/controller"
	"net/http"

	mux "github.com/gorilla/mux"
)

func initRouter() http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/v1").Subrouter()

	//+++++++++++++++++认证API++++++++++++++++++++++++++++++++++++++++++++++++

	s.HandleFunc("/login", controller.Login).Methods("POST", "GET")
	// s.HandleFunc("/register",controller.Register).Methods("POST")
	s.HandleFunc("/loginbygithub", controller.LoginByGitHub).Methods("POST", "GET")
	s.HandleFunc("/isvalid/{token}", controller.ValidToken).Methods("GET")

	//======================简易消息API========================================
	//更新消息状态，已读、未读
	s.HandleFunc("/auth/message/{id}", controller.UpdateMessage).Methods("PUT")
	//find 用户所有消息
	s.HandleFunc("/auth/message/{userName}", controller.FindAllMessage).Methods("GET")
	//未读消息总数
	s.HandleFunc("/auth/message/{userName}/count", controller.FindAllMessageCount).Methods("GET")

	//======================回复API============================================
	//find一个文章的所有评论
	s.HandleFunc("/comments/{articleId}", controller.FindAllCommentByArticleId).Methods("GET")
	//find一个评论下的所有回复
	s.HandleFunc("/replys/{commentId}", controller.FindAllReplyByCommentId).Methods("GET")

	//给一个文章发表评论
	s.HandleFunc("/auth/comment/{articleId}", controller.PostComment).Methods("POST")
	//回复某个评论
	s.HandleFunc("/auth/reply/{commentId}/{toUserName}", controller.PostReply).Methods("POST")

	//======================文章API============================================
	//更新文章
	s.HandleFunc("/auth/article/{articleId}", controller.UpdateArticle).Methods("PUT")
	//发表文章
	s.HandleFunc("/auth/article/{userId}", controller.SaveNewArticle).Methods("POST")
	//首页加载最新发表的文章
	s.HandleFunc("/article/findNewArticles", controller.FindNewArticles).Methods("GET")

	//加载用户所有文章，包括已审核和未审核
	s.HandleFunc("/auth/article/{userId}/all", controller.FindAllArticlesByUser).Methods("GET")
	s.HandleFunc("/article/{userId}/public", controller.FindAllArticlesByUser).Methods("GET")

	//===================其他API=================================================

	s.HandleFunc("/hotAuthor", controller.HotAuthor).Methods("GET")
	s.HandleFunc("/hotArticle", controller.HotArticle).Methods("GET")

	s.HandleFunc("/ping", ping).Methods("GET")

	//====================用户主页===============================================
	s1 := r.PathPrefix("").Subrouter()
	//用户主页模板
	s1.HandleFunc("/user/{userName}", controller.UserHome).Methods("GET")
	//阅读某个用户的文章，加载公共文章模板
	s1.HandleFunc("/user/{userName}/article/{articleId}", controller.FindOneArticle).Methods("GET")
	//加载编辑、发帖模板
	s1.HandleFunc("/article/edit/{userName}/{articleId}", controller.LoadEditArticleTemplate).Methods("GET")
	return r
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte("pong"))
}
