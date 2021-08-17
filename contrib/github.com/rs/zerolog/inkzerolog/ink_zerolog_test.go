package inkzerolog_test

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/rs/zerolog"

	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/contrib/github.com/rs/zerolog/inkzerolog"
	"github.com/hemantjadon/ink/internal/inktest"
)

func TestNewSink(t *testing.T) {
	t.Parallel()

	t.Run(`nil logger`, func(t *testing.T) {
		t.Parallel()

		sink := inkzerolog.NewSink(nil)
		if sink != nil {
			t.Fatalf("got sink = %T, want sink = %v", sink, nil)
		}
	})

	t.Run(`non nil logger`, func(t *testing.T) {
		t.Parallel()

		logger := zerolog.New(io.Discard)
		sink := inkzerolog.NewSink(&logger)
		if sink == nil {
			t.Fatalf("got sink = %v, want sink = <non-nil>", sink)
		}
	})
}

func TestSink_suite(t *testing.T) {
	t.Parallel()

	suite := inktest.Suite{
		NewSink: func(w io.Writer) ink.LogSink {
			logger := zerolog.New(w)
			logger.Level(zerolog.TraceLevel)
			return inkzerolog.NewSink(&logger)
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
	return fmt.Sprintf("%q:%q", "message", msg)
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
	return fmt.Sprintf("%q:%q", "level", lvl)
}

func FormatField(field ink.Field) string {
	switch field.Type() {
	case ink.FieldTypeString:
		return fmt.Sprintf("%q:%q", field.Key(), field.String())
	case ink.FieldTypeInt64:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatInt(field.Int64(), 10))
	case ink.FieldTypeInt32:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatInt(int64(field.Int32()), 10))
	case ink.FieldTypeInt16:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatInt(int64(field.Int16()), 10))
	case ink.FieldTypeInt8:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatInt(int64(field.Int8()), 10))
	case ink.FieldTypeInt:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatInt(int64(field.Int()), 10))
	case ink.FieldTypeUint64:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatUint(field.Uint64(), 10))
	case ink.FieldTypeUint32:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatUint(uint64(field.Uint32()), 10))
	case ink.FieldTypeUint16:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatUint(uint64(field.Uint16()), 10))
	case ink.FieldTypeUint8:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatUint(uint64(field.Uint8()), 10))
	case ink.FieldTypeUint:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatUint(uint64(field.Uint()), 10))
	case ink.FieldTypeFloat64:
		return fmt.Sprintf("%q:%s", field.Key(), formatFloat(field.Float64(), 64))
	case ink.FieldTypeFloat32:
		return fmt.Sprintf("%q:%s", field.Key(), formatFloat(float64(field.Float32()), 32))
	case ink.FieldTypeComplex128:
		return fmt.Sprintf("%q:%q", field.Key(), strconv.FormatComplex(field.Complex128(), 'f', -1, 128))
	case ink.FieldTypeComplex64:
		return fmt.Sprintf("%q:%q", field.Key(), strconv.FormatComplex(complex128(field.Complex64()), 'f', -1, 64))
	case ink.FieldTypeBool:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatBool(field.Bool()))
	case ink.FieldTypeTime:
		return fmt.Sprintf("%q:%q", field.Key(), field.Time().Format(time.RFC3339))
	case ink.FieldTypeDuration:
		return fmt.Sprintf("%q:%d", field.Key(), field.Duration().Milliseconds())
	case ink.FieldTypeStringer:
		return fmt.Sprintf("%q:%q", field.Key(), field.Stringer().String())
	case ink.FieldTypeSlice:
		if field.Slice() == nil || reflect.ValueOf(field.Slice()).IsNil() {
			return fmt.Sprintf("%q:%s", field.Key(), "[]")
		}
		var values interface{}
		switch val := field.Slice().(type) {
		case []uint8:
			ints := make([]int64, 0, len(val))
			for _, u8 := range val {
				ints = append(ints, int64(u8))
			}
			values = ints
		case []float64:
			vals := make([]interface{}, 0, len(val))
			for _, f := range val {
				switch {
				case math.IsNaN(f), math.IsInf(f, 1), math.IsInf(f, -1):
					fs := strconv.FormatFloat(f, 'f', -1, 64)
					vals = append(vals, fs)
				default:
					vals = append(vals, f)
				}
			}
			values = vals
		case []float32:
			vals := make([]interface{}, 0, len(val))
			for _, f := range val {
				switch {
				case math.IsNaN(float64(f)), math.IsInf(float64(f), 1), math.IsInf(float64(f), -1):
					fs := strconv.FormatFloat(float64(f), 'f', -1, 64)
					vals = append(vals, fs)
				default:
					vals = append(vals, f)
				}
			}
			values = vals
		case []complex128:
			strs := make([]string, 0, len(val))
			for _, c := range val {
				strs = append(strs, strconv.FormatComplex(c, 'f', -1, 128))
			}
			values = strs
		case []complex64:
			strs := make([]string, 0, len(val))
			for _, c := range val {
				strs = append(strs, strconv.FormatComplex(complex128(c), 'f', -1, 64))
			}
			values = strs
		case []time.Time:
			strs := make([]string, 0, len(val))
			for _, t := range val {
				strs = append(strs, t.Format(time.RFC3339))
			}
			values = strs
		case []time.Duration:
			ints := make([]int64, 0, len(val))
			for _, d := range val {
				ints = append(ints, d.Milliseconds())
			}
			values = ints
		default:
			values = val
		}
		jsonStr, _ := json.Marshal(values)
		return fmt.Sprintf("%q:%s", field.Key(), string(jsonStr))
	case ink.FieldTypeReflect:
		if field.Reflect() == nil {
			return fmt.Sprintf("%q:%s", field.Key(), "null")
		}
		return ""
	default:
		return ""
	}
}

func formatFloat(f float64, bitSize int) string {
	fs := strconv.FormatFloat(f, 'f', -1, bitSize)
	if math.IsNaN(f) || math.IsInf(f, 1) || math.IsInf(f, -1) {
		return `"` + fs + `"`
	}
	return fs
}
