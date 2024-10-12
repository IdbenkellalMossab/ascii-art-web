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

// Page not found
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
	}
	a := result{Res: r.FormValue("banner"), Res1: "\n" + artHandler(r.FormValue("text"), r.FormValue("banner"))}
	renderTemplate(w, "index.html", "Home Page", &a)

}

func renderTemplate(w http.ResponseWriter, tmpl string, title string, result *result) {
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
