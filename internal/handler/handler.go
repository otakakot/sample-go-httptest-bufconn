package handler

import (
	"context"
	"log/slog"

	"github.com/otakakot/sample-go-httptest-bufconn/pkg/api"
	"github.com/otakakot/sample-go-httptest-bufconn/pkg/proto/proto"
)

var _ api.Handler = (*Handler)(nil)

type Handler struct {
	cli proto.HealthClient
}

func New(
	cli proto.HealthClient,
) *Handler {
	return &Handler{
		cli: cli,
	}
}

// GetHealth implements api.Handler.
func (h *Handler) GetHealth(ctx context.Context, params api.GetHealthParams) (api.GetHealthRes, error) {
	slog.InfoContext(ctx, "[REST] begin GetHealth")
	defer slog.InfoContext(ctx, "[REST] end GetHealth")

	if _, err := h.cli.Get(ctx, &proto.GetHealthRequest{}); err != nil {
		return nil, err
	}

	return &api.GetHealthOK{}, nil
}
