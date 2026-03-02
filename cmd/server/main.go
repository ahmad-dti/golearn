package main

import (
    "fmt"
    "net/http"
    "goleam/internal/handler"
)

func main() {
    http.HandleFunc("/", handler.HelloHandler)

    fmt.Println("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}
