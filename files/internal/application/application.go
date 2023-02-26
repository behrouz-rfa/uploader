package application

import (
	"context"
	"uploader/files/internal/domain"
)

type (
	FileBody struct {
		ID       string
		Name     string
		Size     int64
		Location string
	}
	App interface {
		Upload(ctx context.Context, body FileBody) error
	}
	Application struct {
		files domain.FileInfoRepository
	}
)

func NewApplication(files domain.FileInfoRepository) Application {
	return Application{files: files}
}

var _ App = (*Application)(nil)

func (a Application) Upload(ctx context.Context, body FileBody) error {
	//TODO implement me
	var fileInfo = domain.NewFileInfo(body.ID, body.Name, body.Size, body.Location)
	err := a.files.Upload(fileInfo)
	if err != nil {
		return err
	}
	return nil
}
