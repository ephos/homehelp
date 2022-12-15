package main

import (
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Tempy!</h1>"))
}

func main() {

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	logger.Println("Server is starting...")
	http.ListenAndServe(":"+port, mux)
}
