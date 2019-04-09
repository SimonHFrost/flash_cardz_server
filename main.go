package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Translation struct {
	Side1 string `json:"side1"`
	Side2 string `json:"side2"`
}

func main() {
	http.HandleFunc("/", handleMainRoute)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	translations := []Translation{
		{
			Side1: "Tierra",
			Side2: "Land",
		},
		{
			Side1: "Agua",
			Side2: "Water",
		},
	}

	// Set Content-Type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	if err := json.NewEncoder(w).Encode(translations); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
