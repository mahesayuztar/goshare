package controllers

import (
	"errors"
	"goshare/models"
)

func GetFileByID(id string) (*models.File, error) {
	var fileRecord models.File
	if err := DB.First(&fileRecord, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &fileRecord, nil
}

func GetAllFiles() ([]models.File, error) {
	var files []models.File
	if err := DB.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func CreateFile(f *models.File) error {
	if DB.First(&models.File{}, "id = ?", f.ID).RowsAffected > 0 {
		return errors.New("URL sudah digunakan, silakan pilih URL lain")
	}
	return DB.Create(f).Error
}

func DeleteFile(f *models.File) error {
	return DB.Delete(f).Error
}
