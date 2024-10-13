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
	Err  string
}

var (
	templates = template.Must(template.ParseGlob("templates/*.html"))
	res       result
)

// Error handler
func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		// Handle 404 Not Found
		fmt.Fprint(w, "Page not found (404)")
	case http.StatusMethodNotAllowed:
		// Handle 405 Method Not Allowed
		fmt.Fprint(w, "Method not allowed (405)")
	case http.StatusInternalServerError:
		// Handle 500 Internal Server Error
		fmt.Fprint(w, "Internal server error (500)")
	default:
		// Handle other statuses generically
		fmt.Fprintf(w, "Error %d", status)
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
				Res1: "",
				Err:  "",
			}
		} else {
			// Return an error if the path is incorrect
			errorHandler(w, http.StatusNotFound)
		}
	case http.MethodPost:
		if r.URL.Path == "/ascii-art" {
			// Handle POST requests for /ascii-art
			r.ParseForm()
			text := r.FormValue("text")
			banner := r.FormValue("banner")
			if text == "" || banner == "" { // Set error message if any field is empty
				res = result{
					Err: "Text or Banner cannot be empty",
				}
			} else if len(text) > 700 { // Check if text exceeds 700 characters
				res = result{
					Err: "Please enter less than 700 characters.",
				}
			} else {
				// Generate ascii-art
				artResult := artHandler(text, banner)

				// Check if special characters are present
				if artResult == "Special charactere is not allowed." {
					// Set error message for non-printable characters
					res = result{
						Err: "Please enter printable ASCII characters only.",
					}
				} else {
					// Process form values and generate ASCII art if valid
					res = result{
						Res:  banner,
						Res1: "\n" + artResult,
						Err:  "", // Clear error
					}
				}
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			// 404 Error for incorrect path
			errorHandler(w, http.StatusNotFound)
		}
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

// func Generate ascii-art
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
