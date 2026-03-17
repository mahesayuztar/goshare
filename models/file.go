package models

import (
	"time"
)

type File struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	FileName  string // Nama file yang diupload
	FilePath  string // Path penyimpanan file di server
}
