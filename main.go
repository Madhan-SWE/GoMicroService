package main

import (
  "fmt"
  "net/http"
  "time"
)

func appHome(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world! \n")
  currentTime := time.Now()
  fmt.Fprintf(w, currentTime.Format(time.UnixDate))
}

func setupRoutes() {
  http.HandleFunc("/", appHome)
}

func main() {
  fmt.Println("Running on port: 3000")
  setupRoutes()
  http.ListenAndServe(":3000", nil)
}

