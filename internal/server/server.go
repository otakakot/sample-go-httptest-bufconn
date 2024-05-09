package server

import (
	"context"
	"log/slog"

	"github.com/otakakot/sample-go-httptest-bufconn/pkg/proto/proto"
)

var _ proto.HealthServer = (*Server)(nil)

type Server struct{}

// Get implements proto.HealthServer.
func (s *Server) Get(ctx context.Context, _ *proto.GetHealthRequest) (*proto.GetHealthResponse, error) {
	slog.InfoContext(ctx, "[gPRC] begin GetHealth")
	defer slog.InfoContext(ctx, "[gRPC] end GetHealth")

	return &proto.GetHealthResponse{}, nil
}
