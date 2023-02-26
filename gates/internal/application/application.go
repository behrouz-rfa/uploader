package application

import "context"

type (
	TranscodeBody struct {
		// ID
		// Required: true
		ID string `json:"id"`
		// File URL
		// Required: true
		URL string `json:"url"`
		// Callback URL
		// Required: false
		CallbackURL string `json:"callback_url,omitempty"`
	}
	App interface {
		Transcode(ctx context.Context, body TranscodeBody) error
	}
	Application struct {
	}
)

func (a Application) Transcode(ctx context.Context, body TranscodeBody) error {
	//TODO implement me
	panic("implement me")
}

var _ App = (*Application)(nil)
