package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zachlatta/angelhack/database"
	"github.com/zachlatta/angelhack/hander"
)

func httpLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := database.Init("postgres",
		os.ExpandEnv("postgres://docker:docker@$DB_1_PORT_5432_TCP_ADDR/docker"))
	if err != nil {
		panic(err)
	}
	defer database.Close()

	r := mux.NewRouter()

	r.Handle("/users",
		handler.AppHandler(handler.CreateUser)).Methods("POST")
	r.Handle("/users/authenticate",
		handler.AppHandler(handler.Authenticate)).Methods("POST")
	r.Handle("/users/me", handler.AppHandler(handler.CurrentUser)).Methods("GET")
	r.Handle("/users/{id}", handler.AppHandler(handler.User)).Methods("GET")

	r.Handle("/entries", handler.AppHandler(handler.CreateEntry)).Methods("POST")
	r.Handle("/entries", handler.AppHandler(handler.Entries)).Methods("GET")
	r.Handle("/entries/{id}", handler.AppHandler(handler.Entity)).Methods("GET")
	r.Handle("/entries/{id}",
		handler.AppHandler(handler.DeleteEntry)).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":"+port, httpLog(http.DefaultServeMux))
}
