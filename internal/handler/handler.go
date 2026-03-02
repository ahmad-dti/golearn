package handler

import (
    "fmt"
    "net/http"
)

// HelloHandler is a simple HTTP handler used by the server.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go in Docker!")
}
