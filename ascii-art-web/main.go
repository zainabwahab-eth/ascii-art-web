package main

import (
	"ascii-art-web/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/ascii-art", controller.AsciiArtHandler)
	http.ListenAndServe(":8080", nil)
}
