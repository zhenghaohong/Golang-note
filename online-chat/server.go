package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 4个文件一起启动
func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8181", router); err != nil {
		fmt.Println("err:", err)
	}
}
