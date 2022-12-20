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

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := handlers.YamlHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", yamlHandler)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from Tanzania"))
}
