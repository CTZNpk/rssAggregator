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

func respondWithEror(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}

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
		Addr:    ":8000",
	}

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerRediness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT KEY NOT FOUND")
	}

	fmt.Println(portString)
	fmt.Println("The Server is listening")
	log.Fatal(srv.ListenAndServe())

}
