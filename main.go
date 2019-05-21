package main

import (
	"edgex-club/config"
	"edgex-club/repository"
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {

	confpath := flag.String("confpath", "env/dev.yaml", "配置文件路径")
	flag.Parse()

	err := config.InitConfig(*confpath)
	if err != nil {
		log.Println("Cann't parse config file, exit!")
		return
	}

	r := initRouter()

	b := repository.DBConnect()
	if b {
		log.Println("connect to db success!")
	} else {
		log.Println("failed connect to db!")
	}

	//用户访问限制功能，定时清除3分钟内已经被锁定的用户，
	//防止map缓存越过内存边界
  go cleanupVisitors()

	server := &http.Server{
		Handler:      GeneralFilter(limit(r)),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server Listen On Port: 8080")
	log.Fatal(server.ListenAndServe())
}
