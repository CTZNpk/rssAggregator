package main

import (
	"net/http"
)

func handlerRediness(w http.ResponseWriter, _ *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
