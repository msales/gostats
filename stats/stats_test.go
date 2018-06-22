package stats_test

import (
	"testing"
	"time"

	"context"

	"github.com/msales/pkg/stats"
	"github.com/stretchr/testify/mock"
)

func TestInc(t *testing.T) {
	m := new(MockStats)
	m.On("Inc", "test", int64(1), float32(1.0), []interface{}(nil)).Return(nil)
	ctx := stats.WithStats(context.Background(), m)

	stats.Inc(ctx, "test", 1, 1.0)

	m.AssertExpectations(t)
}

func TestDec(t *testing.T) {
	m := new(MockStats)
	m.On("Dec", "test", int64(1), float32(1.0), []interface{}(nil)).Return(nil)
	ctx := stats.WithStats(context.Background(), m)

	stats.Dec(ctx, "test", 1, 1.0)

	m.AssertExpectations(t)
}

func TestGauge(t *testing.T) {
	m := new(MockStats)
	m.On("Gauge", "test", float64(1), float32(1.0), []interface{}(nil)).Return(nil)
	ctx := stats.WithStats(context.Background(), m)

	stats.Gauge(ctx, "test", 1, 1.0)

	m.AssertExpectations(t)
}

func TestTiming(t *testing.T) {
	m := new(MockStats)
	m.On("Timing", "test", time.Second, float32(1.0), []interface{}(nil)).Return(nil)
	ctx := stats.WithStats(context.Background(), m)

	stats.Timing(ctx, "test", time.Second, 1.0)

	m.AssertExpectations(t)
}

func TestClose(t *testing.T) {
	m := new(MockStats)
	m.On("Close").Return(nil)
	ctx := stats.WithStats(context.Background(), m)

	stats.Close(ctx)

	m.AssertExpectations(t)
}

func TestTaggedStats_Inc(t *testing.T) {
	m := new(MockStats)
	m.On("Inc", "test", int64(1), float32(1), []interface{}{"foo", "bar", "global", "foobar"}).Return(nil)
	s := stats.NewTaggedStats(m, "global", "foobar")

	s.Inc("test", 1, 1.0, "foo", "bar")

	m.AssertExpectations(t)
}

func TestTaggedStats_Dec(t *testing.T) {
	m := new(MockStats)
	m.On("Dec", "test", int64(1), float32(1), []interface{}{"foo", "bar", "global", "foobar"}).Return(nil)
	s := stats.NewTaggedStats(m, "global", "foobar")

	s.Dec("test", 1, 1.0, "foo", "bar")

	m.AssertExpectations(t)
}

func TestTaggedStats_Gauge(t *testing.T) {
	m := new(MockStats)
	m.On("Gauge", "test", float64(1), float32(1), []interface{}{"foo", "bar", "global", "foobar"}).Return(nil)
	s := stats.NewTaggedStats(m, "global", "foobar")

	s.Gauge("test", 1.0, 1.0, "foo", "bar")

	m.AssertExpectations(t)
}

func TestTaggedStats_Timing(t *testing.T) {
	m := new(MockStats)
	m.On("Timing", "test", time.Millisecond, float32(1), []interface{}{"foo", "bar", "global", "foobar"}).Return(nil)
	s := stats.NewTaggedStats(m, "global", "foobar")

	s.Timing("test", time.Millisecond, 1.0, "foo", "bar")

	m.AssertExpectations(t)
}

func TestTaggedStats_Close(t *testing.T) {
	m := new(MockStats)
	m.On("Close").Return(nil)
	s := stats.NewTaggedStats(m, "global", "foobar")

	s.Close()

	m.AssertExpectations(t)
}

type MockStats struct {
	mock.Mock
}

func (m *MockStats) Inc(name string, value int64, rate float32, tags ...interface{}) error {
	args := m.Called(name, value, rate, tags)
	return args.Error(0)
}

func (m *MockStats) Dec(name string, value int64, rate float32, tags ...interface{}) error {
	args := m.Called(name, value, rate, tags)
	return args.Error(0)
}

func (m *MockStats) Gauge(name string, value float64, rate float32, tags ...interface{}) error {
	args := m.Called(name, value, rate, tags)
	return args.Error(0)
}

func (m *MockStats) Timing(name string, value time.Duration, rate float32, tags ...interface{}) error {
	args := m.Called(name, value, rate, tags)
	return args.Error(0)
}

func (m *MockStats) Close() error {
	args := m.Called()
	return args.Error(0)
}
