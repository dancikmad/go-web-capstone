package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Error loading home page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleMenu(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:4002/data")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading menu items", http.StatusInternalServerError)
	}
	defer r.Body.Close()
	fmt.Println(string(body))

	var menuItems []MenuItem
	err = json.Unmarshal(body, &menuItems)
	if err != nil {
		http.Error(w, "Error decoding menu items", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/menu.html")
	if err != nil {
		http.Error(w, "Error loading menu page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, menuItems)
}

func handleReviewForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/review_form.html")
	if err != nil {
		http.Error(w, "Error loading review form", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleReviewSubmission(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:4002/addReview"
	contentType := "application/json"
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	dish := r.FormValue("dish")
	comments := r.FormValue("comments")
	rating := stringToInt(r.FormValue("rating"))

	review := Review{Name: name, Dish: dish, Comments: comments, Rating: rating}

	reviewData, err := json.Marshal(review)
	if err != nil {
		http.Error(w, "Error encoding review data", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(reviewData))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	http.Redirect(w, r, "/review", http.StatusSeeOther)
}

func handleReviews(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:4002/reviews")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println(string(body))

	if err != nil {
		http.Error(w, "Error reading menu items", http.StatusInternalServerError)
		return
	}

	var reviews []Review
	err = json.Unmarshal(body, &reviews)
	if err != nil {
		http.Error(w, "Error decoding reviews", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/reviews.html")
	if err != nil {
		http.Error(w, "Error loading reviews page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, reviews)
}

func stringToInt(s string) int {
	var i int
	fmt.Scan(s, "%d", &i)
	return i
}
