package main

import (
	_ "edgex-club/cache"
	"net/http"
	"strings"
)

func GeneralFilter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		//log.Println("GeneralFilter===path==" + path)
		if path == "/" {
			http.FileServer(http.Dir("static")).ServeHTTP(w, r)
			return
		}
		if strings.HasSuffix(path, ".html") ||
			strings.HasSuffix(path, ".css") ||
			strings.HasSuffix(path, ".js") ||
			strings.HasSuffix(path, ".json") ||
			strings.HasSuffix(path, ".png") ||
			strings.HasSuffix(path, ".md") ||
			strings.HasSuffix(path, ".svg") ||
			strings.HasPrefix(path, "/vendors") {
			http.FileServer(http.Dir("static")).ServeHTTP(w, r)
			return
		}

		// if strings.HasPrefix(path,"/api/v1/") &&
		//    path != "/api/v1/loginbygithub" &&
		// 	 path != "/api/v1/login" && {
		// 	token := r.Header.Get("edgex-club-token")
		//
		// 	//判断token是否在缓存中，且有效
		// 	_,ok := cache.TokenCache[token]
		//
		// 	if !ok {
		// 		log.Println("非法用户尝试访问认证API")
		// 		// http.Error(w, "非法用户", http.StatusBadRequest)
		// 		http.Redirect(w, r, "/login.html", 302)
		// 		return
		// 	}
		// }

		h.ServeHTTP(w, r)
	})
}
