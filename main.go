package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("./dist"))

	// API Handler
	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "let's go",
			"time":    "Server active",
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