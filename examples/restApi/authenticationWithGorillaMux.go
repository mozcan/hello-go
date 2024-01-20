package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "secret_token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	securedRouter := router.PathPrefix("/secured").Subrouter()
	securedRouter.Use(AuthenticationMiddleware)

	securedRouter.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Secured Endpoint"))
	}).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
