package controllers

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func SaveFile(file multipart.File, fileHeader *multipart.FileHeader) (string, string, error) {
	defer file.Close()
	filename := filepath.Base(fileHeader.Filename)
	jktLocation, _ := time.LoadLocation("Asia/Jakarta")
	filePath := "uploads/" + time.Now().In(jktLocation).Format("20060102150405") + "_" + filename
	dst, err := os.Create(filePath)
	if err != nil {
		return "", "", err
	}

	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return "", "", err
	}

	return filename, filePath, nil
}

func DownloadFile(w http.ResponseWriter, filePath string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("Gagal membuka file / File tidak ditemukan")
	}
	defer file.Close()
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", http.DetectContentType([]byte(fileName)))
	_, err = io.Copy(w, file)
	if err != nil {
		return errors.New("Gagal mengirim file")
	}
	return nil
}
