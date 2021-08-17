package inktest

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/hemantjadon/ink"
)

const (
	logMessage = "test log message"
)

// Suite defines a test suite for behaviour driven testing of ink.LogSink.
type Suite struct {
	NewSink       func(out io.Writer) ink.LogSink
	FormatMessage func(msg string) string
	FormatLevel   func(level Level) string
	FormatField   func(field ink.Field) string
	FormatName    func(name string) string
}

// TestMessage tests the formatting of log messages in the output.
func (s Suite) TestMessage(t *testing.T) {
	t.Run(`empty`, func(t *testing.T) {
		t.Parallel()

		msg := ""

		var buf bytes.Buffer
		sink := s.NewSink(&buf)

		sink.Info(msg)
		gotMsg, err := buf.ReadString('\n')
		if err != nil {
			t.Fatalf("read buffer: %v", err)
		}
		if !strings.Contains(gotMsg, s.FormatMessage(msg)) {
			t.Fatalf("got msg = [%v], does not contain = [%v]", strings.TrimSpace(gotMsg), s.FormatMessage(msg))
		}
	})

	t.Run(`non empty`, func(t *testing.T) {
		t.Parallel()

		msg := logMessage

		var buf bytes.Buffer
		sink := s.NewSink(&buf)

		sink.Info(msg)
		gotMsg, err := buf.ReadString('\n')
		if err != nil {
			t.Fatalf("read buffer: %v", err)
		}
		if !strings.Contains(gotMsg, s.FormatMessage(msg)) {
			t.Fatalf("got msg = [%v], does not contain = [%v]", strings.TrimSpace(gotMsg), s.FormatMessage(msg))
		}
	})
}

// TestLevel tests logging at different levels and checks they are formatted
// correctly in the output.
func (s Suite) TestLevel(t *testing.T) {
	t.Run(`debug`, func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		sink := s.NewSink(&buf)

		sink.Debug(logMessage)
		gotMsg, err := buf.ReadString('\n')
		if err != nil {
			t.Fatalf("read buffer: %v", err)
		}
		if !strings.Contains(gotMsg, s.FormatLevel(LevelDebug)) {
			t.Fatalf("got msg = [%v], does not contain = [%v]", strings.TrimSpace(gotMsg), s.FormatLevel(LevelDebug))
		}
	})

	t.Run(`info`, func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		sink := s.NewSink(&buf)

		sink.Info(logMessage)
		gotMsg, err := buf.ReadString('\n')
		if err != nil {
			t.Fatalf("read buffer: %v", err)
		}
		if !strings.Contains(gotMsg, s.FormatLevel(LevelInfo)) {
			t.Fatalf("got msg = [%v], does not contain = [%v]", strings.TrimSpace(gotMsg), s.FormatLevel(LevelInfo))
		}
	})

	t.Run(`error`, func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		sink := s.NewSink(&buf)

		sink.Error(logMessage)
		gotMsg, err := buf.ReadString('\n')
		if err != nil {
			t.Fatalf("read buffer: %v", err)
		}
		if !strings.Contains(gotMsg, s.FormatLevel(LevelError)) {
			t.Fatalf("got msg = [%v], does not contain = [%v]", strings.TrimSpace(gotMsg), s.FormatLevel(LevelError))
		}
	})
}

// TestFields tests the formatting of all kind of fields logged in the output.
func (s Suite) TestFields(t *testing.T) {
	tests := map[string]struct {
		fields []ink.Field
	}{
		"string":        {fields: stringFieldSet()},
		"int64":         {fields: int64FieldSet()},
		"int32":         {fields: int32FieldSet()},
		"int16":         {fields: int16FieldSet()},
		"int8":          {fields: int8FieldSet()},
		"int":           {fields: intFieldSet()},
		"uint64":        {fields: uint64FieldSet()},
		"uint32":        {fields: uint32FieldSet()},
		"uint16":        {fields: uint16FieldSet()},
		"uint8":         {fields: uint8FieldSet()},
		"uint":          {fields: uintFieldSet()},
		"float64":       {fields: float64FieldSet()},
		"float32":       {fields: float32FieldSet()},
		"complex128":    {fields: complex128FieldSet()},
		"complex64":     {fields: complex64FieldSet()},
		"bool":          {fields: boolFieldSet()},
		"time.Time":     {fields: timeFieldSet()},
		"time.Duration": {fields: durationFieldSet()},
		"stringer":      {fields: stringerFieldSet()},
		"object":        {fields: objectFieldSet()},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			sink := s.NewSink(&buf)

			sink.Info(logMessage, test.fields...)

			gotMsg, err := buf.ReadString('\n')
			if err != nil {
				t.Fatalf("read buffer: %v", err)
			}
			for _, field := range test.fields {
				if !strings.Contains(gotMsg, s.FormatField(field)) {
					t.Fatalf("got msg = [%v], does not contain = [%v]", strings.TrimSpace(gotMsg), s.FormatField(field))
				}
			}
		})
	}
}

// TestFieldsLogSink tests the behaviour of ink.LogSink which is ink.FieldsLogSink.
func (s Suite) TestFieldsLogSink(t *testing.T) {
	sk := s.NewSink(io.Discard)
	_, ok := sk.(ink.FieldsLogSink)
	if !ok {
		t.Fatalf("sink given by NewSink is of type %T which does not implements ink.FieldsLogSink", sk)
	}

	t.Run(`set then get fields`, func(t *testing.T) {
		t.Parallel()

		t.Run(`empty fields`, func(t *testing.T) {
			sk := s.NewSink(io.Discard)

			sink := sk.(ink.FieldsLogSink)

			sink = sink.WithFields()

			gotFields := sink.Fields()
			if len(gotFields) != 0 {
				t.Fatalf("got len(gotFields) = %d, want len(gotFields) = %d", len(gotFields), 0)
			}
		})

		t.Run(`set once`, func(t *testing.T) {
			sk := s.NewSink(io.Discard)

			sink := sk.(ink.FieldsLogSink)

			fields := []ink.Field{
				ink.String("sink_field_1", "str1"),
				ink.Int64("sink_field_2", 1),
			}

			sink = sink.WithFields(fields...)

			gotFields := sink.Fields()
			if len(gotFields) != len(fields) {
				t.Fatalf("len(gotFields) = %d, len(fields) = %d", len(gotFields), len(fields))
			}
			for i := 0; i < len(fields); i++ {
				gotField := gotFields[i]
				wantField := fields[i]

				equalField(t, gotField, wantField)
			}
		})

		t.Run(`set multiple times`, func(t *testing.T) {
			sk := s.NewSink(io.Discard)

			sink := sk.(ink.FieldsLogSink)

			fields1 := []ink.Field{
				ink.String("sink_field_1", "str1"),
				ink.Int64("sink_field_2", 1),
			}

			sink = sink.WithFields(fields1...)

			fields2 := []ink.Field{
				ink.Bool("sink_field_3", true),
				ink.Time("sink_field_4", time.Now()),
			}

			sink = sink.WithFields(fields2...)

			allFields := append(fields1, fields2...)

			gotFields := sink.Fields()
			if len(gotFields) != len(allFields) {
				t.Fatalf("len(gotFields) = %d, len(allFields) = %d", len(gotFields), len(allFields))
			}
			for i := 0; i < len(allFields); i++ {
				gotField := gotFields[i]
				wantField := allFields[i]

				equalField(t, gotField, wantField)
			}
		})

	})

	t.Run(`log output`, func(t *testing.T) {
		t.Parallel()

		t.Run(`without additional fields`, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			sk := s.NewSink(&buf)

			sink := sk.(ink.FieldsLogSink)

			fields := []ink.Field{
				ink.String("sink_field_1", "str1"),
				ink.Int64("sink_field_2", 1),
			}

			sink = sink.WithFields(fields...)

			sink.Info("test message")

			gotMsg, err := buf.ReadString('\n')
			if err != nil {
				t.Fatalf("read buffer: %v", err)
			}
			for _, field := range fields {
				if !strings.Contains(gotMsg, s.FormatField(field)) {
					t.Fatalf("got msg = [%v], does not contain = [%v]", strings.TrimSpace(gotMsg), s.FormatField(field))
				}
			}
		})

		t.Run(`with additional fields`, func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			sk := s.NewSink(&buf)

			sink := sk.(ink.FieldsLogSink)

			fields := []ink.Field{
				ink.String("sink_field_1", "str1"),
				ink.Int64("sink_field_2", 1),
			}

			sink = sink.WithFields(fields...)

			sink.Info(logMessage)

			gotMsg, err := buf.ReadString('\n')
			if err != nil {
				t.Fatalf("read buffer: %v", err)
			}
			for _, field := range fields {
				if !strings.Contains(gotMsg, s.FormatField(field)) {
					t.Fatalf("got msg = [%v], does not contain = [%v]", strings.TrimSpace(gotMsg), s.FormatField(field))
				}
			}
		})

		t.Run(`field`, func(t *testing.T) {
			t.Parallel()

			tests := map[string]struct {
				fields []ink.Field
			}{
				"string":        {fields: stringFieldSet()},
				"int64":         {fields: int64FieldSet()},
				"int32":         {fields: int32FieldSet()},
				"int16":         {fields: int16FieldSet()},
				"int8":          {fields: int8FieldSet()},
				"int":           {fields: intFieldSet()},
				"uint64":        {fields: uint64FieldSet()},
				"uint32":        {fields: uint32FieldSet()},
				"uint16":        {fields: uint16FieldSet()},
				"uint8":         {fields: uint8FieldSet()},
				"uint":          {fields: uintFieldSet()},
				"float64":       {fields: float64FieldSet()},
				"float32":       {fields: float32FieldSet()},
				"complex128":    {fields: complex128FieldSet()},
				"complex64":     {fields: complex64FieldSet()},
				"bool":          {fields: boolFieldSet()},
				"time.Time":     {fields: timeFieldSet()},
				"time.Duration": {fields: durationFieldSet()},
				"stringer":      {fields: stringerFieldSet()},
				"object":        {fields: objectFieldSet()},
			}

			for name, test := range tests {
				test := test
				t.Run(name, func(t *testing.T) {
					t.Parallel()

					var buf bytes.Buffer
					sk := s.NewSink(&buf)

					sink := sk.(ink.FieldsLogSink)

					sink = sink.WithFields(test.fields...)

					sink.Info(logMessage)

					gotMsg, err := buf.ReadString('\n')
					if err != nil {
						t.Fatalf("read buffer: %v", err)
					}
					for _, field := range test.fields {
						if !strings.Contains(gotMsg, s.FormatField(field)) {
							t.Fatalf("got msg = [%v], does not contain = [%v]", strings.TrimSpace(gotMsg), s.FormatField(field))
						}
					}
				})
			}
		})
	})
}

// TestNameLogSink tests the behaviour of ink.LogSink which is ink.NameLogSink.
func (s Suite) TestNameLogSink(t *testing.T) {
	sk := s.NewSink(io.Discard)
	_, ok := sk.(ink.NameLogSink)
	if !ok {
		t.Fatalf("sink given by NewSink is of type %T which does not implements ink.NameLogSink", sk)
	}

	t.Run(`set then get name`, func(t *testing.T) {
		t.Parallel()

		t.Run(`empty fields`, func(t *testing.T) {
			t.Parallel()

			sk := s.NewSink(io.Discard)

			sink := sk.(ink.NameLogSink)
			sink = sink.WithName("")

			gotName := sink.Name()
			if len(gotName) != 0 {
				t.Fatalf("got len(gotName) = %d, want len(gotName) = %d", len(gotName), 0)
			}
		})

		t.Run(`set once`, func(t *testing.T) {
			t.Parallel()

			sk := s.NewSink(io.Discard)

			sink := sk.(ink.NameLogSink)

			name := "alpha"
			sink = sink.WithName(name)

			gotName := sink.Name()
			if gotName != name {
				t.Fatalf("got name = %s, name = %s", gotName, name)
			}
		})

		t.Run(`set multiple times`, func(t *testing.T) {
			t.Parallel()

			sk := s.NewSink(io.Discard)

			sink := sk.(ink.NameLogSink)

			name1 := "alpha"
			sink = sink.WithName(name1)

			name2 := "beta"
			sink = sink.WithName(name2)

			gotName := sink.Name()
			if !strings.Contains(gotName, name1) {
				t.Fatalf("got name = %s, does not contain name = %s", gotName, name1)
			}
			if !strings.Contains(gotName, name2) {
				t.Fatalf("got name = %s, does not contain name = %s", gotName, name2)
			}
		})

	})

	t.Run(`log output`, func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		sk := s.NewSink(&buf)

		sink := sk.(ink.NameLogSink)

		name := "alpha"
		sink = sink.WithName(name)

		sink.Info(logMessage)

		gotMsg, err := buf.ReadString('\n')
		if err != nil {
			t.Fatalf("read buffer: %v", err)
		}
		if !strings.Contains(gotMsg, s.FormatName(name)) {
			t.Fatalf("got msg = %s, does not contain name field = %s", gotMsg, s.FormatName(name))
		}
	})
}

func equalField(t testing.TB, got, want ink.Field) {
	t.Helper()

	gotKey := got.Key()
	wantKey := want.Key()
	if gotKey != wantKey {
		t.Fatalf("got key = %s, want key = %s", gotKey, wantKey)
	}

	gotType := got.Type()
	wantType := want.Type()
	if gotType != wantType {
		t.Fatalf("got type = %d, want type = %d", gotType, wantType)
	}

	gotValue := got.Value()
	wantValue := want.Value()
	if gotValue != wantValue {
		t.Fatalf("got value = %s, want value = %s", gotValue, wantValue)
	}
}
