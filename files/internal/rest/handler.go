package rest

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"uploader/files/internal/application"
	"uploader/internal/storage"
)

type Handler interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

type FileHandler struct {
	app         application.App
	fileStorage storage.FileStorage
}

var _ Handler = (*FileHandler)(nil)

func New(app application.App, fileStorage storage.FileStorage) FileHandler {
	return FileHandler{app: app, fileStorage: fileStorage}
}

func (h FileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = os.MkdirAll("./tmpfile", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileName := uuid.New().String() + filepath.Ext(fileHader.Filename)
	dst, err := os.Create(fmt.Sprintf("./tmpfile/%s", fileName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.fileStorage.UploadData("website", fileName, file, fileHader.Size, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	url, _ := h.fileStorage.SignUrl("original", fileName)

	h.app.Upload(context.Background(), application.FileBody{
		ID:       uuid.New().String(),
		Name:     fileName,
		Size:     fileHader.Size,
		Location: url,
	})
	fmt.Fprintf(w, url)
}
