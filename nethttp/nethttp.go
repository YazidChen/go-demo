package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// [1:]表示从根"/"后第一个字符开始到末尾
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func main() {
	// 处理所有跟路径的路由走handler方法
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
