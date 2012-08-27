package main

import (
  "fmt"
  "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(w, "Hello World!")
  fmt.Println("just hello'd", req)
}

func main() {
  http.HandleFunc("/", hello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
      panic(err)
  }
}
