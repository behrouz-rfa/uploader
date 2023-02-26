package application

import "context"

type (
	FileBody struct {
		FileBody interface{}
	}
	App interface {
		Upload(ctx context.Context, body FileBody) error
	}
	Application struct {
	}
)

func NewApplication() Application {
	return Application{}
}

var _ App = (*Application)(nil)

func (a Application) Upload(ctx context.Context, body FileBody) error {
	//TODO implement me
	panic("implement me")
}
