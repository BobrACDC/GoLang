package main

import (
	"log"
	"net/http"
)

type Handler struct {}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("cat /etc/passwd"))
	if err != nil {
		log.Fatalln()
	}
}

func main() {
	h := Handler{}
	server := http.Server{
		Addr:    ":9999",
		Handler: h,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}


