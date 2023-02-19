package gates

import (
	"context"
	"uploader/internal/monolith"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	return nil
}
