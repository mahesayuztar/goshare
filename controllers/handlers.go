package controllers

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
	"strings"
)

var TmpPtr *template.Template

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	TmpPtr.ExecuteTemplate(w, "index.html", nil)
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	url := r.FormValue("url")
	fmt.Println("URL:", url)
	_, fileHeader, _ := r.FormFile("file")
	if len(url) > 192 {
		TmpPtr.ExecuteTemplate(w, "index.html", map[string]string{
			"Error": "URL terlalu panjang, maksimal 192 karakter",
		})
		return
	}
	if fileHeader.Size > 10*1024*1024 {
		TmpPtr.ExecuteTemplate(w, "index.html", map[string]string{
			"Error": "Ukuran file terlalu besar, maksimal 10MB",
		})
		return
	}

	insertErr := InsertFile(r)
	if insertErr != nil {
		TmpPtr.ExecuteTemplate(w, "index.html", map[string]string{
			"Error": "Gagal Upload File: " + html.UnescapeString(insertErr.Error()),
		})
		return
	}

	TmpPtr.ExecuteTemplate(w, "success.html", map[string]string{
		"FileID": url,
	})
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/download/")
	fmt.Println("Download request for ID:", id)
	row, err := GetFileByID(id)
	if err != nil {
		TmpPtr.ExecuteTemplate(w, "index.html", err.Error())
		return
	}
	err = DownloadFile(w, row.FilePath, row.FileName)
	if err != nil {
		TmpPtr.ExecuteTemplate(w, "index.html", err.Error())
		return
	}
	TmpPtr.ExecuteTemplate(w, "success.html", map[string]string{
		"Download": id,
	})
}
