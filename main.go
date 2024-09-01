package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/BowangLan/golang-rss-generator/internal/database"
	"github.com/joho/godotenv"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in environmental variables")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is not found in environmental variables")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Failed to establish a database connection")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	fmt.Println("Database connection established")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
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

	v1Router.Route("/users", func(r chi.Router) {
		r.Post("/", apiCfg.handleCreateUser)
		r.Get("/", apiCfg.handleListUsers)
		r.Get("/{id}", apiCfg.handleGetUser)
		r.Get("/me", apiCfg.handleGetUserWithAuth)
	})

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Printf("Server starting at port %v\n", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
