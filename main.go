package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {

	godotenv.Load()

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://"},
		AllowedMethods:   []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT KEY NOT FOUND")
	}

	fmt.Println(portString)
	fmt.Println("The Server is listening")
	log.Fatal(srv.ListenAndServe())

}
