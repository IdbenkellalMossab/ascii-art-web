package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type result struct {
	Res  string
	Res1 string
	Err  string
}

var (
	templates = template.Must(template.ParseGlob("Templates/*.html"))
	res       result
)

func main() {
	http.Handle("/Js/", http.StripPrefix("/Js/", http.FileServer(http.Dir("Js"))))
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// handleGet processes GET requests and renders the home page.
		handleGet(w, r)
	case http.MethodPost:
		// handlePost processes POST requests for the ASCII Art generation.
		handlePost(w, r)
	default:
		// 405 Error for unsupported methods
		errorHandler(w, http.StatusMethodNotAllowed)
	}
}

func renderTemplate(w http.ResponseWriter, title string, result *result) {
	err := templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Title":  title,
		"Result": result, // Pass the result to the template
	})
	// Server error
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
}
