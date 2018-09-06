package clix_test

import (
	"errors"
	"flag"
	"testing"
	"time"

	"github.com/msales/pkg/clix"
	"github.com/msales/pkg/log"
	"github.com/msales/pkg/stats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
)

func TestLogger(t *testing.T) {
	ctx := &clix.Context{}

	fn := clix.Logger(log.Null)

	assert.IsType(t, clix.CtxOptionFunc(nil), fn)

	fn(ctx)

	l, ok := log.FromContext(ctx)

	assert.True(t, ok)
	assert.Equal(t, l, log.Null)
}

func TestStats(t *testing.T) {
	ctx := &clix.Context{}

	fn := clix.Stats(stats.Null)

	assert.IsType(t, clix.CtxOptionFunc(nil), fn)

	fn(ctx)

	s, ok := stats.FromContext(ctx)

	assert.True(t, ok)
	assert.Equal(t, s, stats.Null)
}

func TestNewContext(t *testing.T) {
	c := cli.NewContext(nil, flag.NewFlagSet("", flag.ContinueOnError), nil)

	ctx, err := clix.NewContext(c, clix.Logger(log.Null), clix.Stats(stats.Null))

	assert.IsType(t, &clix.Context{}, ctx)
	assert.NoError(t, err)
}

func TestContext_Close(t *testing.T) {
	tests := []struct {
		err error
	}{
		{nil},
		{errors.New("")},
	}

	for _, tt := range tests {
		s := new(MockStats)
		s.On("Close").Return(tt.err)

		c := cli.NewContext(nil, flag.NewFlagSet("", flag.ContinueOnError), nil)
		ctx, err := clix.NewContext(c, clix.Logger(log.Null), clix.Stats(s))
		assert.NoError(t, err)

		err = ctx.Close()

		assert.Equal(t, err, tt.err)
	}
}

type MockStats struct {
	mock.Mock
}

func (m *MockStats) Inc(name string, value int64, rate float32, tags ...interface{}) error {
	return nil
}

func (m *MockStats) Dec(name string, value int64, rate float32, tags ...interface{}) error {
	return nil
}

func (m *MockStats) Gauge(name string, value float64, rate float32, tags ...interface{}) error {
	return nil
}

func (m *MockStats) Timing(name string, value time.Duration, rate float32, tags ...interface{}) error {
	return nil
}

func (m *MockStats) Close() error {
	args := m.Called()
	return args.Error(0)
}