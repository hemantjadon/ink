package ink_test

import (
	"bytes"
	"testing"

	"github.com/hemantjadon/ink/inkio"

	"github.com/hemantjadon/ink"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	t.Run("with sink", func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		sink := inkio.NewSink(&buf)
		logger := ink.NewLogger(ink.WithSink(sink))

		logger.Info("test log message")

		got, err := buf.ReadString('\n')
		if err != nil {
			t.Fatalf("unable to read buffer: %v", err)
		}
		if len(got) == 0 {
			t.Fatalf("unable to log message")
		}
	})

	t.Run("with name", func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		sink := inkio.NewSink(&buf)
		logger := ink.NewLogger(ink.WithSink(sink), ink.WithName("test_logger"))

		logger.Info("test log message")

		got, err := buf.ReadString('\n')
		if err != nil {
			t.Fatalf("unable to read buffer: %v", err)
		}
		if len(got) == 0 {
			t.Fatalf("unable to log message")
		}
	})
}
