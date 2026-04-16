package main

import (
	"ascii-art-web/operations"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type PageData struct {
	Input  string
	Banner string
	Result string
	Error  string
}

// var tmpl = template.Must(template.ParseFiles("templates/index.html"))
var tmpl = template.Must(template.ParseFiles("templates/index.html", "templates/notFound.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//Make sure path is "/". Check if path is not "/"
	if r.URL.Path != "/" {
		//Write status 404
		w.WriteHeader(http.StatusNotFound)

		//Load error page template and return Error
		err := tmpl.ExecuteTemplate(w, "notFound.html", PageData{Error: "Page Not Found"})

		//If there is template loading error
		if err != nil {
			log.Println("template error:", err)
		}

		return
	}

	//If the correct route "/" is entered print the index page
	if err := tmpl.ExecuteTemplate(w, "index.html", PageData{}); err != nil {
    log.Println("template error:", err)
    return
}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// 400 — bad input from the user
	if text == "" || banner == "" {
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, PageData{Error: "text and banner are required"})
		return
	}

	bannerLoc := banner + ".txt"

	// read banner
	data, err := operations.ReadTextFile(bannerLoc)

	if err != nil {
		// 404 — banner file wasn't found
		w.WriteHeader(http.StatusNotFound)
		tmpl.Execute(w, PageData{Error: "banner file not found"})
		return
	}

	inputSlice := strings.Split(text, "\\n")

	//Call ascii art logic
	result, err := operations.AsciiArt(inputSlice, data)
	if err != nil {
		// 500 — template itself is broken/missing
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, PageData{Error: "Something went wrong"})
		return
	}

	tmpl.Execute(w, PageData{Input: text, Banner: banner, Result: result})
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.ListenAndServe(":8080", nil)
}
