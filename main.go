package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type RouteResponse struct {
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
}

func main() {
	log.Println("Starting server...")

	router := mux.NewRouter()

	log.Println("Setting up routes...")

	router.Handle("/register", alice.New(loggingMiddleware).ThenFunc(register)).Methods("POST")
	router.Handle("/login", alice.New(loggingMiddleware).ThenFunc(login)).Methods("POST")
	router.Handle("/projects", alice.New(loggingMiddleware).ThenFunc(createProject)).Methods("POST")
	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(updateProject)).Methods("PUT")
	router.Handle("/projects", alice.New(loggingMiddleware).ThenFunc(getProjects)).Methods("GET")
	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(getProject)).Methods("GET")
	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(deleteProject)).Methods("DELETE")

	log.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
			next.ServeHTTP(w, r)
		})
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from register"})
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from login"})
}

func createProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from createProject"})
}

func updateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from updateProject", ID: id})
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from getProjects"})
}

func getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from getProject", ID: id})
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from deleteProject", ID: id})
}
