package main

import (
	"net/http"

	"github.com/jkarage/gophercise/quiz/urlshortener/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	pathsToUrl := map[string]string{
		"/jokes": "https://google.com/",
		"/serve": "https://alinker.tk",
	}
	mapHandler := handlers.MapHandler(pathsToUrl, mux)
	http.ListenAndServe(":8080", mapHandler)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from Tanzania"))
}
