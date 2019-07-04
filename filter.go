package main

import (
	"edgex-club/authorization"
	"encoding/json"
	"log"
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

		//检测认证API是否携带有效jwt token
		if strings.HasPrefix(path, "/api/v1/auth") {
			token := r.Header.Get("edgex-club-token")
			if token == "" {
				log.Println("token 为空！")
				tokenS, _ := r.Cookie("edgex-club-token")
				token = tokenS.Value
				// w.WriteHeader(http.StatusUnauthorized)
				// return
			}
			claims := &authorization.Claims{}
			ok := false
			if ok, claims = authorization.CheckToken(token); !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			credsByte, err := json.Marshal(claims.Credentials)

			if err != nil {
				log.Println("转换creds失败！")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			//重写请求头，用于下游服务中使用，比如一些controller中需要用到用户信息
			r.Header.Set("inner-user", string(credsByte))
		}

		h.ServeHTTP(w, r)
	})
}
