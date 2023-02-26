package postgres

import (
	"gorm.io/gorm"
	"uploader/files/internal/domain"
)

type FileInfoRepository struct {
	db *gorm.DB
}

func NewFileInfoRepository(db *gorm.DB) FileInfoRepository {
	return FileInfoRepository{db: db}
}

var _ domain.FileInfoRepository = (*FileInfoRepository)(nil)

func (f FileInfoRepository) Upload(info domain.FileInfo) error {
	//TODO implement me
	f.db.Create(info)
	return nil
}
