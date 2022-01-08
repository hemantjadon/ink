package inkio_test

import (
	"io"
	"testing"

	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/inkio"
	"github.com/hemantjadon/ink/inkio/internal/istrconv"
	"github.com/hemantjadon/ink/internal/inktest"
)

func TestNewSink(t *testing.T) {
	t.Parallel()

	t.Run(`nil writer`, func(t *testing.T) {
		t.Parallel()

		sink := inkio.NewSink(nil)

		if sink != nil {
			t.Fatalf("got sink = %T, want sink = %v", sink, nil)
		}
	})

	t.Run(`non nil writer`, func(t *testing.T) {
		t.Parallel()

		sink := inkio.NewSink(io.Discard)

		if sink == nil {
			t.Fatalf("got sink = %v, want sink = <non-nil>", sink)
		}
	})
}

func TestSink_suite(t *testing.T) {
	t.Parallel()

	t.Run(`suite`, func(t *testing.T) {
		t.Parallel()

		newSink := func(w io.Writer) ink.LogSink {
			return inkio.NewSink(w)
		}

		suite := inktest.Suite{
			NewSink:       newSink,
			FormatMessage: FormatMessage,
			FormatLevel:   FormatLevel,
			FormatField:   FormatField,
		}

		suite.TestLevel(t)
		suite.TestMessage(t)
		suite.TestFields(t)
	})
}

func FormatMessage(msg string) string {
	return msg
}

func FormatLevel(level inktest.Level) string {
	switch level {
	case inktest.LevelDebug:
		return "DEBUG"
	case inktest.LevelInfo:
		return "INFO"
	case inktest.LevelError:
		return "ERROR"
	default:
		return ""
	}
}

func FormatField(field ink.Field) string {
	return istrconv.FormatField(field)
}
