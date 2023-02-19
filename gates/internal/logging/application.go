package logging

import (
	"github.com/rs/zerolog"
	"uploader/gates/internal/application"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)
