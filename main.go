package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Time struct {
	CurrentTime string `json:"current_time"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Helper Boooooiiiiiii!"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := []Time{
		{CurrentTime: http.TimeFormat},
	}
	json.NewEncoder(w).Encode(currentTime)
}

func main() {

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/time", timeHandler)

	logger.Println("Server is starting...")
	http.ListenAndServe("localhost:8080", mux)
}
