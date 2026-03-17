package controllers

import (
	"errors"
	"goshare/models"
	"net/http"
)

func InsertUser(formValue map[string][]string) (models.User, error) {
	user := models.User{
		Name: formValue["name"][0],
	}
	res := DB.Create(&user)
	return user, res.Error
}

func InsertFile(r *http.Request) error {
	if DB.First(&models.File{}, "id = ?", r.FormValue("url")).RowsAffected > 0 {
		return errors.New("URL sudah digunakan, silakan pilih URL lain")
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return err
	}

	filename, filePath, err := SaveFile(file, fileHeader)
	if err != nil {
		return err
	}

	fileRecord := models.File{
		ID:       r.FormValue("url"),
		FileName: filename,
		FilePath: filePath,
	}
	if err := DB.Create(&fileRecord).Error; err != nil {
		return err
	}

	return nil
}

func GetFileByID(id string) (models.File, error) {
	var fileRecord models.File
	if err := DB.First(&fileRecord, "id = ?", id).Error; err != nil {
		return models.File{}, errors.New("File tidak ditemukan")
	}
	return fileRecord, nil
}
