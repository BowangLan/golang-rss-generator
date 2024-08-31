package main

import (
  "fmt"
  "os"
  "log"

  "github.com/joho/godotenv"

  "net/http"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
  "github.com/go-chi/cors"
)

func main() {
	fmt.Println("Hello, World!")

  godotenv.Load(".env")
  
  portString := os.Getenv("PORT");
  if portString == "" {
    log.Fatal("PORT is not found in environmental variables");
  }

  router := chi.NewRouter();

  router.Use(middleware.RequestID)
  router.Use(middleware.RealIP)
  router.Use(middleware.Logger)
  router.Use(middleware.Recoverer) 

  router.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

  v1Router := chi.NewRouter()
  v1Router.Get("/ready", handlerReadiness)
  v1Router.Get("/error", handlerError)

  router.Mount("/v1", v1Router)

  srv := &http.Server{
    Handler: router,
    Addr: ":" + portString,
  }

  fmt.Printf("Server starting at port %v", portString)

  err := srv.ListenAndServe()
  if err != nil {
    log.Fatal(err)
  }
}
