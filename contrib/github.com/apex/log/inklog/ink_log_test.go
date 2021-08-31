package inklog_test

import (
	"fmt"
	"io"
	"testing"

	apexlog "github.com/apex/log"
	apextext "github.com/apex/log/handlers/text"

	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/contrib/github.com/apex/log/inklog"
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

		logger := &apexlog.Logger{}
		sink := inklog.NewSink(logger)
		if sink == nil {
			t.Fatalf("got sink = %v, want sink = <non-nil>", sink)
		}
	})
}

func TestSink_suite(t *testing.T) {
	t.Parallel()

	suite := inktest.Suite{
		NewSink: func(w io.Writer) ink.LogSink {
			logger := apexlog.Logger{
				Handler: apextext.New(w),
				Level:   apexlog.DebugLevel,
			}
			return inklog.NewSink(&logger)
		},
		FormatMessage: FormatMessage,
		FormatLevel:   FormatLevel,
		FormatField:   FormatFieldInfoText,
	}

	suite.TestMessage(t)
	suite.TestLevel(t)
	suite.TestFields(t)
	suite.TestFieldsLogSink(t)
}

func FormatMessage(msg string) string {
	if len(msg) == 0 {
		return ""
	}
	return msg
}

func FormatLevel(level inktest.Level) string {
	var lvl string
	switch level {
	case inktest.LevelDebug:
		lvl = "DEBUG"
	case inktest.LevelInfo:
		lvl = "INFO"
	case inktest.LevelError:
		lvl = "ERROR"
	}
	return lvl
}

func FormatFieldInfoText(field ink.Field) string {
	return fmt.Sprintf("\033[%dm%s\033[0m=%v", 34, field.Key(), field.Value())
}
