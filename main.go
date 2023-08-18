package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type FlashCard struct {
	ID    string `json:"id"`
	Side1 string `json:"side1"`
	Side2 string `json:"side2"`
}

// Global variable for storing flashcards
var flashcards []FlashCard
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handleMainRoute)
	http.HandleFunc("/add", handleAddFlashCard)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	// Return the current list of flashcards
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(flashcards); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func handleAddFlashCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

	// Parse the flashcard from the request body
	var card FlashCard
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Add to our in-memory list
	mu.Lock()
	flashcards = append(flashcards, card)
	mu.Unlock()

	// Send a 201 Created response
	w.WriteHeader(http.StatusCreated)
}

func init() {
	// Initialize with some default flashcards
	flashcards = []FlashCard{
		{ID: "1", Side1: "Tierra", Side2: "Land"},
		{ID: "2", Side1: "Agua", Side2: "Water"},
	}
}
