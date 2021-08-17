package inkzap_test

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/contrib/github.com/uber-go/zap/inkzap"
	"github.com/hemantjadon/ink/internal/inktest"
)

func TestNewSink(t *testing.T) {
	t.Parallel()

	t.Run(`nil logger`, func(t *testing.T) {
		t.Parallel()

		sink := inkzap.NewSink(nil)
		if sink != nil {
			t.Fatalf("got sink = %T, want sink = %v", sink, nil)
		}
	})

	t.Run(`non nil logger`, func(t *testing.T) {
		t.Parallel()

		sink := inkzap.NewSink(zap.L())
		if sink == nil {
			t.Fatalf("got sink = %v, want sink = <non-nil>", sink)
		}
	})
}

func TestSink_suite(t *testing.T) {
	t.Parallel()

	suite := inktest.Suite{
		NewSink: func(w io.Writer) ink.LogSink {
			zc := zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				FunctionKey:    zapcore.OmitKey,
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}
			logger := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(zc), noopSync{Writer: w}, zap.DebugLevel))
			return inkzap.NewSink(logger)
		},
		FormatMessage: FormatMessage,
		FormatLevel:   FormatLevel,
		FormatField:   FormatField,
		FormatName:    FormatName,
	}

	suite.TestMessage(t)
	suite.TestLevel(t)
	suite.TestFields(t)
	suite.TestFieldsLogSink(t)
	suite.TestNameLogSink(t)
}

func FormatMessage(msg string) string {
	return fmt.Sprintf("%q:%q", "msg", msg)
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
	return fmt.Sprintf("%q:%q", "level", lvl)
}

func FormatName(name string) string {
	return fmt.Sprintf("%q:%q", "logger", name)
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
		return fmt.Sprintf("%q:%q", field.Key(), formatComplex(field.Complex128(), 128))
	case ink.FieldTypeComplex64:
		return fmt.Sprintf("%q:%q", field.Key(), formatComplex(complex128(field.Complex64()), 64))
	case ink.FieldTypeBool:
		return fmt.Sprintf("%q:%s", field.Key(), strconv.FormatBool(field.Bool()))
	case ink.FieldTypeTime:
		return fmt.Sprintf("%q:%q", field.Key(), field.Time().Format(time.RFC3339Nano))
	case ink.FieldTypeDuration:
		return fmt.Sprintf("%q:%q", field.Key(), field.Duration().String())
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
		case []complex128:
			strs := make([]string, 0, len(val))
			for _, c := range val {
				strs = append(strs, formatComplex(c, 128))
			}
			values = strs
		case []complex64:
			strs := make([]string, 0, len(val))
			for _, c := range val {
				strs = append(strs, formatComplex(complex128(c), 64))
			}
			values = strs
		case []time.Duration:
			strs := make([]string, 0, len(val))
			for _, d := range val {
				strs = append(strs, d.String())
			}
			values = strs
		default:
			values = val
		}
		jsonStr, _ := json.Marshal(values)
		return fmt.Sprintf("%q:%s", field.Key(), string(jsonStr))
	case ink.FieldTypeReflect:
		if field.Reflect() == nil {
			return fmt.Sprintf("%q:%s", field.Key(), "null")
		}
		jsonStr, _ := json.Marshal(field.Reflect())
		return fmt.Sprintf("%q:%s", field.Key(), string(jsonStr))
	default:
		return ""
	}
}

type noopSync struct {
	io.Writer
}

func (noopSync) Sync() error {
	return nil
}

func formatFloat(f float64, bitSize int) string {
	fs := strconv.FormatFloat(f, 'f', -1, bitSize)
	if math.IsNaN(f) || math.IsInf(f, 1) || math.IsInf(f, -1) {
		return `"` + fs + `"`
	}
	return fs
}

func formatComplex(c complex128, bitSize int) string {
	if bitSize != 64 && bitSize != 128 {
		panic("invalid bitSize")
	}
	bitSize >>= 1 // complex128 uses float64 and complex64 uses float32 internally

	r, i := real(c), imag(c)
	rs := strconv.FormatFloat(r, 'f', -1, bitSize)
	is := strconv.FormatFloat(i, 'f', -1, bitSize)

	// Check if imaginary part has a sign. If not, add one.
	if is[0] != '+' && is[0] != '-' {
		is = "+" + is
	}

	return rs + is + "i"
}
