package main

import (
	"net/http"
	"time"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/routers"
)

func main() {
	r := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1048576 Byte = 1024 KB = 1 M
	}
	_ = s.ListenAndServe()
}
