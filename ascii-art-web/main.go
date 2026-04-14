package main

import (
	"ascii-art-web/operations"
	"net/http"
	"strings"
	"text/template"
)

type PageData struct {
	Input  string
	Banner string
	Result string
	Error  string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		tmpl.Execute(w, PageData{Error: "page not found"})
		return
	}

	// 500 — template itself is broken/missing
	if err := tmpl.Execute(w, PageData{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
