package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World hahah ")

}

func main() {

	godotenv.Load()

	http.HandleFunc("/", handler)

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT KEY NOT FOUND")
	}

	fmt.Println(portString)
  fmt.Println("The Server is listening")
  log.Fatal(http.ListenAndServe(":8080", nil))

}
