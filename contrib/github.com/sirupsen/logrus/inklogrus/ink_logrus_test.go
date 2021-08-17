package inklogrus_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/contrib/github.com/sirupsen/logrus/inklogrus"
	"github.com/hemantjadon/ink/internal/inktest"
)

func TestNewSink(t *testing.T) {
	t.Parallel()

	t.Run(`nil logger`, func(t *testing.T) {
		t.Parallel()

		sink := inklogrus.NewSink(nil)
		if sink != nil {
			t.Fatalf("got sink = %T, want sink = %v", sink, nil)
		}
	})

	t.Run(`non nil logger`, func(t *testing.T) {
		t.Parallel()

		sink := inklogrus.NewSink(logrus.New())
		if sink == nil {
			t.Fatalf("got sink = %v, want sink = <non-nil>", sink)
		}
	})
}

func TestSink_suite(t *testing.T) {
	t.Parallel()

	suite := inktest.Suite{
		NewSink: func(w io.Writer) ink.LogSink {
			logger := logrus.New()
			logger.SetOutput(w)
			logger.SetFormatter(&logrus.TextFormatter{
				DisableQuote: true,
			})
			logger.SetLevel(logrus.TraceLevel)
			return inklogrus.NewSink(logger)
		},
		FormatMessage: FormatMessage,
		FormatLevel:   FormatLevel,
		FormatField:   FormatField,
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
	return fmt.Sprintf("%s=%s", "msg", msg)
}

func FormatLevel(level inktest.Level) string {
	var lvl string
	switch level {
	case inktest.LevelDebug:
		lvl = "debug"
	case inktest.LevelInfo:
		lvl = "info"
	case inktest.LevelError:
		lvl = "error"
	}
	return fmt.Sprintf("%s=%s", "level", lvl)
}

func FormatField(field ink.Field) string {
	return fmt.Sprintf("%s=%v", field.Key(), field.Value())
}
