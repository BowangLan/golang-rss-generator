package main

import (
  "encoding/json"
  "net/http"
  "log"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
  if code > 499 {
    log.Println("Responding with 5XX level error")
  }

  type errResponse struct {
    Error string `json:"error"`
  }

  respondWithJson(w, code, errResponse{
    Error: msg,
  })
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
  dat, err := json.Marshal(payload);

  if err != nil {
    w.WriteHeader(500)
    log.Fatal("Failed to marshal JSON response: %v", payload)
    return
  }

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(dat)
}
