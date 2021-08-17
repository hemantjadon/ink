package inklog_test

import (
	"io"
	"log"
	"testing"

	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/contrib/log/inklog"
	"github.com/hemantjadon/ink/contrib/log/inklog/internal/istrconv"
	"github.com/hemantjadon/ink/internal/inktest"
)

func TestNewSink(t *testing.T) {
	t.Parallel()

	t.Run(`nil logger`, func(t *testing.T) {
		t.Parallel()

		sink := inklog.NewSink(nil)

		if sink != nil {
			t.Fatalf("got sink = %T, want sink = %v", sink, nil)
		}
	})

	t.Run(`non nil logger`, func(t *testing.T) {
		t.Parallel()

		sink := inklog.NewSink(log.Default())

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
			logger := log.New(w, "", log.LstdFlags)
			return inklog.NewSink(logger)
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
