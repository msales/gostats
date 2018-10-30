package stats_test

import (
	"context"
	"testing"
	"time"

	"github.com/msales/pkg/v3/stats"
	"github.com/stretchr/testify/mock"
)

func TestRuntime(t *testing.T) {
	m := new(MockStats)
	m.On("Gauge", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	m.On("Timing", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	stats.DefaultRuntimeInterval = time.Microsecond

	go stats.Runtime(m)

	time.Sleep(100 * time.Millisecond)

	m.AssertCalled(t, "Gauge", "runtime.cpu.goroutines", mock.Anything, mock.Anything, mock.Anything)
}

func TestRuntimeFromContext(t *testing.T) {
	m := new(MockStats)
	m.On("Gauge", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	m.On("Timing", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ctx := stats.WithStats(context.Background(), m)

	go stats.RuntimeFromContext(ctx, time.Millisecond)

	time.Sleep(100 * time.Millisecond)

	m.AssertCalled(t, "Gauge", "runtime.cpu.goroutines", mock.Anything, mock.Anything, mock.Anything)
}
