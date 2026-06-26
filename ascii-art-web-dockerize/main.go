package main

import (
	controller "ascii-art-web/handler"
	"log"
	"net/http"
)

func main() {
	//Serve static file eg css and js
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))


	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/ascii-art", controller.AsciiArtHandler)

	//Create server
	log.Println("Serving on :8080...")
	http.ListenAndServe(":8080", nil)
}
