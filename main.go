package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "courses")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/courses", coursesHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	// Serve static files (CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := 90
	fmt.Printf("Server running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

