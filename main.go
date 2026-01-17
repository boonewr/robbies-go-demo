package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func main() {
	// Serve React static files from the "dist" directory
	fs := http.FileServer(http.Dir("./dist"))

	// API Handler
	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		current := count
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Let's Go",
			"count":   current,
		})
	})

	// serve index.html 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("dist", r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join("dist", "index.html"))
			return
		}
		fs.ServeHTTP(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}