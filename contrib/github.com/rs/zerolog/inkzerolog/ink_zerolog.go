// Package inkzerolog provides compatibility layer between ink and rs/zerolog
// package. See https://github.com/rs/zerolog for details about zerolog.
package inkzerolog

import (
	"strconv"
	"time"

	"github.com/rs/zerolog"

	"github.com/hemantjadon/ink"
)

// Sink describes a zerolog.Logger based ink.LogSink.
type Sink struct {
	logger *zerolog.Logger
	fields []ink.Field
}

var (
	_ ink.LogSink       = (*Sink)(nil)
	_ ink.FieldsLogSink = (*Sink)(nil)
)

// NewSink wraps the given zerolog.Logger into an ink.LogSink.
func NewSink(logger *zerolog.Logger) *Sink {
	if logger == nil {
		return nil
	}
	sink := Sink{logger: logger}
	return &sink
}

// Debug logs to underlying zerolog.Logger at DEBUG level.
func (s Sink) Debug(msg string, fields ...ink.Field) {
	evt := s.logger.Debug()
	evt = addFieldsToEvent(evt, fields)
	evt.Msg(msg)
}

// Info logs to underlying zerolog.Logger at INFO level.
func (s Sink) Info(msg string, fields ...ink.Field) {
	evt := s.logger.Info()
	evt = addFieldsToEvent(evt, fields)
	evt.Msg(msg)
}

// Error logs to underlying zerolog.Logger at ERROR level.
func (s Sink) Error(msg string, fields ...ink.Field) {
	evt := s.logger.Error()
	evt = addFieldsToEvent(evt, fields)
	evt.Msg(msg)
}

func (s Sink) WithFields(fields ...ink.Field) ink.FieldsLogSink {
	if len(fields) == 0 {
		return s
	}
	sink := s.clone()
	ctx := s.logger.With()
	ctx = addFieldsToContext(ctx, fields)
	logger := ctx.Logger()
	sink.logger = &logger
	sink.fields = make([]ink.Field, 0, len(sink.fields)+len(fields))
	sink.fields = append(sink.fields, s.fields...)
	sink.fields = append(sink.fields, fields...)
	return sink
}

func (s Sink) Fields() []ink.Field {
	fs := make([]ink.Field, 0, len(s.fields))
	fs = append(fs, s.fields...)
	return fs
}

func (s Sink) clone() Sink {
	sink := Sink{logger: s.logger}
	if len(s.fields) == 0 {
		return sink
	}
	fs := make([]ink.Field, 0, len(s.fields))
	fs = append(fs, s.fields...)
	return sink
}

func addFieldsToEvent(evt *zerolog.Event, fields []ink.Field) *zerolog.Event {
	for _, field := range fields {
		key := field.Key()
		switch field.Type() {
		case ink.FieldTypeString:
			evt = evt.Str(key, field.String())
		case ink.FieldTypeInt64:
			evt = evt.Int64(key, field.Int64())
		case ink.FieldTypeInt32:
			evt = evt.Int32(key, field.Int32())
		case ink.FieldTypeInt16:
			evt = evt.Int16(key, field.Int16())
		case ink.FieldTypeInt8:
			evt = evt.Int8(key, field.Int8())
		case ink.FieldTypeInt:
			evt = evt.Int(key, field.Int())
		case ink.FieldTypeUint64:
			evt = evt.Uint64(key, field.Uint64())
		case ink.FieldTypeUint32:
			evt = evt.Uint32(key, field.Uint32())
		case ink.FieldTypeUint16:
			evt = evt.Uint16(key, field.Uint16())
		case ink.FieldTypeUint8:
			evt = evt.Uint8(key, field.Uint8())
		case ink.FieldTypeUint:
			evt = evt.Uint(key, field.Uint())
		case ink.FieldTypeFloat64:
			evt = evt.Float64(key, field.Float64())
		case ink.FieldTypeFloat32:
			evt = evt.Float32(key, field.Float32())
		case ink.FieldTypeComplex128:
			evt = evt.Str(key, strconv.FormatComplex(field.Complex128(), 'f', -1, 128))
		case ink.FieldTypeComplex64:
			evt = evt.Str(key, strconv.FormatComplex(complex128(field.Complex64()), 'f', -1, 64))
		case ink.FieldTypeBool:
			evt = evt.Bool(key, field.Bool())
		case ink.FieldTypeTime:
			evt = evt.Time(key, field.Time())
		case ink.FieldTypeDuration:
			evt = evt.Dur(key, field.Duration())
		case ink.FieldTypeStringer:
			evt = evt.Stringer(key, field.Stringer())
		case ink.FieldTypeSlice:
			switch val := field.Slice().(type) {
			case []string:
				evt = evt.Strs(key, val)
			case []int64:
				evt = evt.Ints64(key, val)
			case []int32:
				evt = evt.Ints32(key, val)
			case []int16:
				evt = evt.Ints16(key, val)
			case []int8:
				evt = evt.Ints8(key, val)
			case []int:
				evt = evt.Ints(key, val)
			case []uint64:
				evt = evt.Uints64(key, val)
			case []uint32:
				evt = evt.Uints32(key, val)
			case []uint16:
				evt = evt.Uints16(key, val)
			case []uint8:
				evt = evt.Uints8(key, val)
			case []uint:
				evt = evt.Uints(key, val)
			case []float64:
				evt = evt.Floats64(key, val)
			case []float32:
				evt = evt.Floats32(key, val)
			case []complex128:
				strs := make([]string, 0, len(val))
				for _, c := range val {
					strs = append(strs, strconv.FormatComplex(c, 'f', -1, 128))
				}
				evt = evt.Strs(key, strs)
			case []complex64:
				strs := make([]string, 0, len(val))
				for _, c := range val {
					strs = append(strs, strconv.FormatComplex(complex128(c), 'f', -1, 64))
				}
				evt = evt.Strs(key, strs)
			case []bool:
				evt = evt.Bools(key, val)
			case []time.Time:
				evt = evt.Times(key, val)
			case []time.Duration:
				evt = evt.Durs(key, val)
			default:
				evt = evt.Interface(key, val)
			}
		case ink.FieldTypeReflect:
			evt = evt.Interface(key, field.Reflect())
		default:
			evt = evt.Interface(key, field.Value())
		}
	}
	return evt
}

func addFieldsToContext(ctx zerolog.Context, fields []ink.Field) zerolog.Context {
	for _, field := range fields {
		key := field.Key()
		switch field.Type() {
		case ink.FieldTypeString:
			ctx = ctx.Str(key, field.String())
		case ink.FieldTypeInt64:
			ctx = ctx.Int64(key, field.Int64())
		case ink.FieldTypeInt32:
			ctx = ctx.Int32(key, field.Int32())
		case ink.FieldTypeInt16:
			ctx = ctx.Int16(key, field.Int16())
		case ink.FieldTypeInt8:
			ctx = ctx.Int8(key, field.Int8())
		case ink.FieldTypeInt:
			ctx = ctx.Int(key, field.Int())
		case ink.FieldTypeUint64:
			ctx = ctx.Uint64(key, field.Uint64())
		case ink.FieldTypeUint32:
			ctx = ctx.Uint32(key, field.Uint32())
		case ink.FieldTypeUint16:
			ctx = ctx.Uint16(key, field.Uint16())
		case ink.FieldTypeUint8:
			ctx = ctx.Uint8(key, field.Uint8())
		case ink.FieldTypeUint:
			ctx = ctx.Uint(key, field.Uint())
		case ink.FieldTypeFloat64:
			ctx = ctx.Float64(key, field.Float64())
		case ink.FieldTypeFloat32:
			ctx = ctx.Float32(key, field.Float32())
		case ink.FieldTypeComplex128:
			ctx = ctx.Str(key, strconv.FormatComplex(field.Complex128(), 'f', -1, 128))
		case ink.FieldTypeComplex64:
			ctx = ctx.Str(key, strconv.FormatComplex(complex128(field.Complex64()), 'f', -1, 64))
		case ink.FieldTypeBool:
			ctx = ctx.Bool(key, field.Bool())
		case ink.FieldTypeTime:
			ctx = ctx.Time(key, field.Time())
		case ink.FieldTypeDuration:
			ctx = ctx.Dur(key, field.Duration())
		case ink.FieldTypeStringer:
			ctx = ctx.Stringer(key, field.Stringer())
		case ink.FieldTypeSlice:
			switch val := field.Slice().(type) {
			case []string:
				ctx = ctx.Strs(key, val)
			case []int64:
				ctx = ctx.Ints64(key, val)
			case []int32:
				ctx = ctx.Ints32(key, val)
			case []int16:
				ctx = ctx.Ints16(key, val)
			case []int8:
				ctx = ctx.Ints8(key, val)
			case []int:
				ctx = ctx.Ints(key, val)
			case []uint64:
				ctx = ctx.Uints64(key, val)
			case []uint32:
				ctx = ctx.Uints32(key, val)
			case []uint16:
				ctx = ctx.Uints16(key, val)
			case []uint8:
				ctx = ctx.Uints8(key, val)
			case []uint:
				ctx = ctx.Uints(key, val)
			case []float64:
				ctx = ctx.Floats64(key, val)
			case []float32:
				ctx = ctx.Floats32(key, val)
			case []complex128:
				strs := make([]string, 0, len(val))
				for _, c := range val {
					strs = append(strs, strconv.FormatComplex(c, 'f', -1, 128))
				}
				ctx = ctx.Strs(key, strs)
			case []complex64:
				strs := make([]string, 0, len(val))
				for _, c := range val {
					strs = append(strs, strconv.FormatComplex(complex128(c), 'f', -1, 64))
				}
				ctx = ctx.Strs(key, strs)
			case []bool:
				ctx = ctx.Bools(key, val)
			case []time.Time:
				ctx = ctx.Times(key, val)
			case []time.Duration:
				ctx = ctx.Durs(key, val)
			default:
				ctx = ctx.Interface(key, val)
			}
		case ink.FieldTypeReflect:
			ctx = ctx.Interface(key, field.Reflect())
		default:
			ctx = ctx.Interface(key, field.Value())
		}
	}
	return ctx
}
