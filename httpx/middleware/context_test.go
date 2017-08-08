package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/msales/pkg/httpx/middleware"
	"github.com/stretchr/testify/assert"
)

func TestWithContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "test", "test")

	h := middleware.WithContext(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, ctx, r.Context())
	}), ctx)

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	h.ServeHTTP(resp, req)
}