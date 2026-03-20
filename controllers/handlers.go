package controllers

import (
	"encoding/json"
	"fmt"
	"goshare/models"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var TmpPtr *template.Template

type APIResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	TmpPtr.ExecuteTemplate(w, "index.html", nil)
}

func CreateFileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseMultipartForm(5 << 20)

	fileName, filePath, err := SaveFile(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Status: "fail",
			Error:  "Gagal menyimpan file: " + err.Error(),
		})
		return
	}

	file := &models.File{
		ID:       r.FormValue("url"),
		FileName: fileName,
		FilePath: filePath,
	}

	if err := CreateFile(file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Status: "fail",
			Error:  "Gagal menyimpan file: " + err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(APIResponse{
		Status: "success",
	})
}

func GetFileHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	id = strings.TrimSpace(id)
	fmt.Println("Download request for ID:", id)

	row, err := GetFileByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(APIResponse{
			Status: "fail",
			Error:  "File tidak ditemukan",
		})
		return
	}

	err = DownloadFile(w, row.FilePath, row.FileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Status: "fail",
			Error:  "Gagal mengunduh file: " + err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(APIResponse{
		Status: "success",
	})
}
