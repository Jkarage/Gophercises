package handlers

import "net/http"

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortener := r.URL.Path
		if original, ok := pathToUrls[shortener]; ok {
			http.Redirect(w, r, original, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
