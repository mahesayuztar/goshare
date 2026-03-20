package main

import (
	"fmt"
	"goshare/controllers"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	controllers.DB = controllers.Connect()
	fmt.Println("Database Goshare connected")

	controllers.TmpPtr, _ = template.ParseGlob("templates/*.html")

	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HomeHandler)
	router.HandleFunc("/files", controllers.CreateFileHandler).Methods("POST")
	router.HandleFunc("/files/{id}", controllers.GetFileHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, router)
}
