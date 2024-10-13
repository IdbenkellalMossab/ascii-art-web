package main

import (
	"fmt"
	"html/template"
	"net/http"

	function "function/Functions"
)

type result struct {
	Res  string
	Res1 string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))
var res result

// Page not found
func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Page not found 404")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/" {
			// Display the homepage
			renderTemplate(w, "Home Page", &res)
			res = result{
				Res:  "", // Clear previous values
				Res1: "", // Clear previous values
			}
		} else {
			// Return an error if the path is incorrect
			errorHandler(w, http.StatusNotFound)
		}
	case http.MethodPost:
		if r.URL.Path == "/ascii-art" {
			// Handle POST requests for /ascii-art
			r.ParseForm()
			res = result{
				Res:  r.FormValue("banner"),
				Res1: "\n" + artHandler(r.FormValue("text"), r.FormValue("banner")),
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			// Return an error if the path is incorrect
			errorHandler(w, http.StatusNotFound)
		}
	default:
		// Return an error if the request method is not supported
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func renderTemplate(w http.ResponseWriter, title string, result *result) {
	err := templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Title":  title,
		"Result": result, // Pass the result to the template
	})
	// Server error
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
	}
}

func artHandler(sentence string, banner string) string {
	if len(sentence) == 0 {
		return ""
	}

	symboles, err := function.ReadSymbols(banner)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return function.PrintWords(function.Split(sentence), symboles)
}
