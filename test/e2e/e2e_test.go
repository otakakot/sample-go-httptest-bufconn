package e2e_test

import (
	"context"
	"net"
	"net/http/httptest"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/otakakot/sample-go-httptest-bufconn/internal/handler"
	"github.com/otakakot/sample-go-httptest-bufconn/internal/server"
	"github.com/otakakot/sample-go-httptest-bufconn/pkg/api"
	"github.com/otakakot/sample-go-httptest-bufconn/pkg/proto/proto"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	lis := bufconn.Listen(1024 * 1024)

	srv := grpc.NewServer()

	proto.RegisterHealthServer(srv, &server.Server{})

	go func() {
		if err := srv.Serve(lis); err != nil {
			t.Errorf("failed to serve grpc: %v", err)
		}
	}()

	dialer := func(string, time.Duration) (net.Conn, error) {
		return lis.Dial()
	}

	defer srv.GracefulStop()

	cc, err := grpc.Dial(lis.Addr().String(), grpc.WithDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial grpc: %v", err)
	}

	defer cc.Close()

	hdl, err := api.NewServer(handler.New(proto.NewHealthClient(cc)))
	if err != nil {
		t.Fatalf("failed to create server: %v", err)
	}

	rest := httptest.NewServer(hdl)

	defer rest.Close()

	cli, err := api.NewClient(rest.URL)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	res, err := cli.GetHealth(context.Background(), api.GetHealthParams{})
	if err != nil {
		t.Errorf("failed to get health: %v", err)
	}

	switch res.(type) {
	case *api.GetHealthOK:
	default:
		t.Errorf("unexpected response: %v", res)
	}
}
