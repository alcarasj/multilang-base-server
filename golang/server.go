package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

const PORT = 8080

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome! This is a base server written in Golang.")
}

func main() {
    http.HandleFunc("/", index)
    fmt.Println("Server listening on port " + strconv.Itoa(PORT) + ".")
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(PORT), nil))
}
