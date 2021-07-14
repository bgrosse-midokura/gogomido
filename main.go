package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "gogomido/endpoints"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Location", "/health")
    w.WriteHeader(http.StatusMovedPermanently)
    w.Write([]byte("GoGo Mido!\n"))
}


func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", RootHandler)

    r.HandleFunc("/health", endpoints.HealthCheckHandler)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
