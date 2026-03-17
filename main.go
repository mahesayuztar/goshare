package main

import (
	"fmt"
	"goshare/controllers"
	"html/template"
	"net/http"
)

func main() {
	controllers.DB = controllers.Connect()
	fmt.Println("Database Goshare connected")

	controllers.TmpPtr, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/submit", controllers.SubmitHandler)
	http.HandleFunc("/download/", controllers.DownloadHandler)
	// http.HandleFunc("/file/", handlers.DownloadHandler(fileService))
	http.ListenAndServe(":8080", nil)
}
