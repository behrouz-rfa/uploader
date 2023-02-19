package application

type (
	App interface {
	}
	Application struct {
	}
)

var _ App = (*Application)(nil)
