package main

import (
    "net/http"
    "fmt"
)


func main() {
    http.HandleFunc("/", indexHandler)
    fmt.Println("Starting server in 8080 ...")
    http.ListenAndServe(":8080", nil)
}
