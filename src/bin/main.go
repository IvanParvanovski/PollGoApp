package main

import (
	"fmt"
	// "mainapp/pkg/models"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}

func main() {
	router := http.NewServeMux()
	
	router.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Success ", router)
}
