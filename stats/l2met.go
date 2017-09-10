package stats

import (
	"bytes"
	"time"

	"github.com/msales/pkg/log"
	"fmt"
	"strings"
)

// Statsd represents a statsd client.
type L2met struct {
	log    log.Logger
	prefix string
}

// NewStatsd create a Statsd instance.
func NewL2met(l log.Logger, prefix string) Stats {
	return &L2met{
		log:    l,
		prefix: prefix,
	}
}

// Inc increments a count by the value.
func (s L2met) Inc(name string, value int64, rate float32, tags map[string]string) error {
	msg := s.formatL2metMetric(name, fmt.Sprintf("%d", value), "count", tags)
	s.log.Info(msg)

	return nil
}

// Dec decrements a count by the value.
func (s L2met) Dec(name string, value int64, rate float32, tags map[string]string) error {
	msg := s.formatL2metMetric(name, fmt.Sprintf("-%d", value), "count", tags)
	s.log.Info(msg)

	return nil
}

// Gauge measures the value of a metric.
func (s L2met) Gauge(name string, value float64, rate float32, tags map[string]string) error {
	msg := s.formatL2metMetric(name, fmt.Sprintf("%v", value), "measure", tags)
	s.log.Info(msg)

	return nil
}

// Timing sends the value of a Duration.
func (s L2met) Timing(name string, value time.Duration, rate float32, tags map[string]string) error {
	msg := s.formatL2metMetric(name, fmt.Sprintf("%v", value), "measure", tags)
	s.log.Info(msg)

	return nil
}

func (s L2met) formatL2metMetric(name, value, measure string, tags map[string]string) string {
	if s.prefix != "" {
		name = strings.Join([]string{s.prefix, name}, ".")
	}

	var buf bytes.Buffer
	buf.WriteString(formatL2metTags(tags))
	buf.WriteString(measure)
	buf.WriteString("#")
	buf.WriteString(name)
	buf.WriteString("=")
	buf.WriteString(value)

	return buf.String()
}

// formatStatsdTags formats into an InfluxDB style string
func formatL2metTags(tags map[string]string) string {
	if len(tags) == 0 {
		return ""
	}

	var buf bytes.Buffer
	for k, v := range tags {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
		buf.WriteString(" ")
	}

	return buf.String()
}