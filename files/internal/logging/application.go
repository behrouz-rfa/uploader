package logging

import (
	"github.com/rs/zerolog"
	"uploader/files/internal/application"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)
