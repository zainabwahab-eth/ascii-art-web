package main

import (
	"ascii-art-web/handler"
	"log"
	"net/http"
)

func main() {
	//Serve static file eg css and js
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))


	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/ascii-art", handler.AsciiArtHandler)
	http.HandleFunc("/download", handler.DownloadHandler)


	//Create server
	log.Println("Serving on :8080...")
	http.ListenAndServe(":8080", nil)
}
