package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func origin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request in server-2 at: %s\n", time.Now())
	responseString := "Hello, from server 2!"

	// Convert the string to a JSON object
	jsonResponse := map[string]string{"message": responseString}

	//Encoding of json data using marshal
	jsonData, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response body
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {

	//creating a router
	r := mux.NewRouter()
	r.HandleFunc("/", origin).Methods("GET")
	fmt.Println("Server-2 running on port:8002")
	log.Fatal(http.ListenAndServe(":8002", r))
}
