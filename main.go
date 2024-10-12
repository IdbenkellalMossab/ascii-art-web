package main

import (
    "net/http"
    "fmt"
)


func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/ascii-art", asciiArtHandler)

    fmt.Println("Starting server in 8080 ...")
    http.ListenAndServe(":8080", nil)
}
