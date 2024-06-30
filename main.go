package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteResponse struct {
	Message string `json:"message"`
}

func main() {
	log.Println("Starting server...")

	router := mux.NewRouter()

	log.Println("Setting up routes...")

	router.HandleFunc("/").Methods("GET")

	log.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
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
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from updateProject"})
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from getProjects"})
}

func getProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from getProject"})
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Hello from deleteProject"})
}
