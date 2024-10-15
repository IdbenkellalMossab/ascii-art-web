package main

import (
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusBadRequest)
		return
	}
	if r.URL.Path == "/" {
		// Display the homepage
		renderTemplate(w, "Home Page", &res)
		res = result{
			Res:  "", // Clear previous values
			Res1: "",
			Err:  "",
		}
	} else {
		errorHandler(w, http.StatusNotFound)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	if r.URL.Path == "/ascii-art" {
		// Parse the form data
		if err := r.ParseForm(); err != nil {
			setError(w, r, "Failed to parse form data.")
			return
		}
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		if text == "" || banner == "" { // Set error message if any field is empty
			setError(w, r, "Text or Banner cannot be empty")
			return
		}
		if len(text) > 700 { // Check if text exceeds 700 characters
			setError(w, r, "Please enter less than 700 characters.")
			return
		}
		// Generate ascii-art
		artResult := artHandler(text, banner)

		// Check if special characters are present
		if artResult == "Special charactere is not allowed." {
			// Set error message for non-printable characters
			setError(w, r, "Please enter printable ASCII characters only.")
			return
		} else {
			// Process form values and generate ASCII art if valid
			res = result{
				Res:  banner,
				Res1: "\n" + artResult,
				Err:  "", // Clear error
			}
		}

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// 404 Error for incorrect path
		errorHandler(w, http.StatusNotFound)
	}
}
