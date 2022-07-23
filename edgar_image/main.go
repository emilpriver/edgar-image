package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleImageHttp)

	log.Println("listening on", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
