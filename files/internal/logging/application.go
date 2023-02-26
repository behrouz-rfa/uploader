package logging

import (
	"context"
	"github.com/rs/zerolog"
	"uploader/files/internal/application"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

func NewApplication(app application.App, logger zerolog.Logger) Application {
	return Application{App: app, logger: logger}
}

var _ application.App = (*Application)(nil)

func (a Application) Upload(ctx context.Context, body application.FileBody) (err error) {
	a.logger.Info().Msgf("-->FileInfo.Upload")
	defer func() { a.logger.Info().Err(err).Msg("<-- FileInfo.Upload") }()
	return a.App.Upload(ctx, body)
}
