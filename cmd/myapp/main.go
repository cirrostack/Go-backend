package main

import (
    "log"
    "net/http"
    "github.com/cirrostack/Go-backend/internal/handler"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/hello", handler.HelloHandler)

    log.Println("Serverul rulează pe :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}