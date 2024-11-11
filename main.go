package main

import (
	"golang-beginner-chap24/routers"
	"log"
	"net/http"
)

func main() {
	r := routers.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	fu := http.FileServer(http.Dir("./uploads"))
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", fu))

	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}

	http.ListenAndServe(":8080", r)
}
