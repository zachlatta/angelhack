package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zachlatta/angelhack/database"
	"github.com/zachlatta/angelhack/hander"
	"github.com/zachlatta/cors"
)

const (
	ApplicationEnvironment = "ANGELHACK_ENV"
	Production             = "PRODUCTION"
)

var corsHandler http.Handler

func httpLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		corsHandler.ServeHTTP(w, r)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("Application started.")

	production := os.Getenv(ApplicationEnvironment) == Production

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if production {
		log.Println("Production detected. Looking for $DATABASE_URL.")

		databaseURL := os.Getenv("DATABASE_URL")

		if databaseURL == "" {
			log.Fatal("DATABASE_URL is empty")
		}

		err := database.Init("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			panic(err)
		}
	} else {
		err := database.Init("postgres",
			os.ExpandEnv("postgres://docker:docker@$DB_1_PORT_5432_TCP_ADDR/docker"))
		if err != nil {
			panic(err)
		}
	}
	defer database.Close()

	corsHandler = cors.Allow(&cors.Options{
		AllowAllOrigins: true,
	})

	r := mux.NewRouter()

	r.Handle("/users",
		handler.AppHandler(handler.CreateUser)).Methods("POST")
	r.Handle("/users/authenticate",
		handler.AppHandler(handler.Authenticate)).Methods("POST")
	r.Handle("/users/{id}", handler.AppHandler(handler.User)).Methods("GET")

	r.Handle("/journals",
		handler.AppHandler(handler.CreateJournal)).Methods("POST")
	r.Handle("/journals",
		handler.AppHandler(handler.Journals)).Methods("GET")
	r.Handle("/journals/{id}",
		handler.AppHandler(handler.Journal)).Methods("GET")

	r.Handle("/journals/{journalID}/entries",
		handler.AppHandler(handler.CreateJournalEntry)).Methods("POST")
	r.Handle("/journals/{journalID}/entries",
		handler.AppHandler(handler.JournalEntries)).Methods("GET")

	r.Handle("/entries/{id}",
		handler.AppHandler(handler.Entry)).Methods("GET")

	// OPTIONS catchall for CORS.
	r.HandleFunc("/{**}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	http.Handle("/", r)
	http.ListenAndServe(":"+port, httpLog(http.DefaultServeMux))
}
