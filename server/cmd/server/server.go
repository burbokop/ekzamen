package main

import (
	"log"
	"net/http"
)

func StartServer() error {
	log.Print("Server started on\n127.0.0.1:")
	http.HandleFunc("/", RootHandleFunc)
	return http.ListenAndServe(":8080", nil)
}

func RootHandleFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("gogadoda"))
}
