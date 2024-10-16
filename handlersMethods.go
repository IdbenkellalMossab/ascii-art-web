package main

import (
	"net/http"
	"strings"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
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

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, http.StatusMethodNotAllowed)
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
			//setError(w, r, "Text or Banner cannot be empty")
			errorHandler(w, http.StatusBadRequest)
			return
		}
		if len(text) > 700 { // Check if text exceeds 700 characters
			//setError(w, r, "Please enter less than 700 characters.")
			errorHandler(w, http.StatusBadRequest)
			return
		}
		// Generate ascii-art
		artResult := artHandler(text, banner)

		// Check if special characters are present
		if artResult == "Special charactere is not allowed." {
			// Set error message for non-printable characters
			//setError(w, r, "Please enter printable ASCII characters only.")
			errorHandler(w, http.StatusBadRequest)
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

// Function to handle requests to the /Js/ path
func jsHandler(w http.ResponseWriter, r *http.Request) {
	// Check the path
	if strings.HasPrefix(r.URL.Path, "/Js/") {
		// If the request is directly to /Js/, return a forbidden error
		if r.URL.Path == "/Js/" {
			errorHandler(w, http.StatusForbidden)
			return
		}
		// If the request is for a specific file, pass the request to http.FileServer
		http.StripPrefix("/Js/", http.FileServer(http.Dir("Js"))).ServeHTTP(w, r)
	} else {
		// If the path is incorrect, return a 404 error
		errorHandler(w, http.StatusNotFound)
	}
}