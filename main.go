package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(handleHome))
	mux.Handle("GET /menu", http.HandlerFunc(handleMenu))
	mux.Handle("GET /reviews", http.HandlerFunc(handleReviews))
	mux.Handle("GET /review", http.HandlerFunc(handleReviewForm))
	mux.Handle("POST /submitReview", http.HandlerFunc(handleReviewSubmission))

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting the server at http://localhost:4001")

	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error occurred while starting the server:", err)
	}
}
