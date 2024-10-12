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

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		// Just return without processing or redirecting
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	res = result{Res: r.FormValue("banner"), Res1: "\n" + artHandler(r.FormValue("text"), r.FormValue("banner"))}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}
	renderTemplate(w, "Home Page", &res)
	res = result{
		Res:  "", // Set to empty string
		Res1: "", // Set to empty string
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
