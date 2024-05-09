// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// GetHealth implements getHealth operation.
//
// Health.
//
// GET /health
func (UnimplementedHandler) GetHealth(ctx context.Context, params GetHealthParams) (r GetHealthRes, _ error) {
	return r, ht.ErrNotImplemented
}