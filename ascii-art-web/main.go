package main

import (
	"ascii-art-web/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/ascii-art", handler.AsciiArtHandler)
	http.ListenAndServe(":8080", nil)
}
