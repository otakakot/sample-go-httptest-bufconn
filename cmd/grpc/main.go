package main

import (
	"cmp"
	"context"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/otakakot/sample-go-httptest-bufconn/internal/server"
	"github.com/otakakot/sample-go-httptest-bufconn/pkg/proto/proto"
	"google.golang.org/grpc"
)

func main() {
	port := cmp.Or(os.Getenv("PORT"), "8888")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterHealthServer(srv, &server.Server{})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		slog.Info("start grpc server. listen")

		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()

	slog.Info("start grpc server shutdown")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	srv.GracefulStop()

	slog.Info("done grpc server shutdown")
}
