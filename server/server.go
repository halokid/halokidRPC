package main

import (
  "fmt"
  "golang.org/x/net/http2"
  "log"
  "net/http"
)

var svc map[string]interface{}

func main() {
  // register service
  svc = make(map[string]interface{})

  // make server
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "HoloKidRPC")
  })

  s := &http.Server{
    Addr:       ":9527",
    Handler:    mux,
  }
  http2.ConfigureServer(s, &http2.Server{})
  log.Fatal(s.ListenAndServe())
}

