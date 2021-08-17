// Package inkzap provides compatibility layer between ink and uber-go/zap
// package. See https://github.com/uber-go/zap for details about zap.
package inkzap

import (
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/hemantjadon/ink"
)

// Sink describes a zap.Logger based ink.LogSink.
type Sink struct {
	logger *zap.Logger
	fields []ink.Field
	name   string
}

var (
	_ ink.LogSink       = (*Sink)(nil)
	_ ink.FieldsLogSink = (*Sink)(nil)
	_ ink.NameLogSink   = (*Sink)(nil)
)

// NewSink wraps the given zap.Logger into an ink.LogSink.
func NewSink(logger *zap.Logger) *Sink {
	if logger == nil {
		return nil
	}
	sink := Sink{logger: logger}
	return &sink
}

// Debug logs to underlying zap.Logger at DEBUG level.
func (s Sink) Debug(msg string, fields ...ink.Field) {
	s.logger.Debug(msg, convertFields(fields)...)
}

// Info logs to underlying zap.Logger at INFO level.
func (s Sink) Info(msg string, fields ...ink.Field) {
	s.logger.Info(msg, convertFields(fields)...)
}

// Error logs to underlying zap.Logger at ERROR level.
func (s Sink) Error(msg string, fields ...ink.Field) {
	s.logger.Error(msg, convertFields(fields)...)
}

func (s Sink) WithFields(fields ...ink.Field) ink.FieldsLogSink {
	if len(fields) == 0 {
		return s
	}
	sink := s.clone()
	sink.logger = s.logger.With(convertFields(fields)...)
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

func (s Sink) WithName(name string) ink.NameLogSink {
	if len(name) == 0 {
		return s
	}
	sink := s.clone()
	sink.logger = s.logger.Named(name)
	if len(s.name) == 0 {
		sink.name = name
	} else {
		sink.name = strings.Join([]string{s.name, name}, ".")
	}
	return sink
}

func (s Sink) Name() string {
	return s.name
}

func (s Sink) clone() Sink {
	sink := Sink{logger: s.logger, name: s.name}
	if len(s.fields) == 0 {
		return sink
	}
	fs := make([]ink.Field, 0, len(s.fields))
	fs = append(fs, s.fields...)
	return sink
}

func convertFields(fields []ink.Field) []zap.Field {
	zfs := make([]zap.Field, 0, len(fields))
	for _, field := range fields {
		key := field.Key()
		switch field.Type() {
		case ink.FieldTypeString:
			zfs = append(zfs, zap.String(key, field.String()))
		case ink.FieldTypeInt64:
			zfs = append(zfs, zap.Int64(key, field.Int64()))
		case ink.FieldTypeInt32:
			zfs = append(zfs, zap.Int32(key, field.Int32()))
		case ink.FieldTypeInt16:
			zfs = append(zfs, zap.Int16(key, field.Int16()))
		case ink.FieldTypeInt8:
			zfs = append(zfs, zap.Int8(key, field.Int8()))
		case ink.FieldTypeInt:
			zfs = append(zfs, zap.Int(key, field.Int()))
		case ink.FieldTypeUint64:
			zfs = append(zfs, zap.Uint64(key, field.Uint64()))
		case ink.FieldTypeUint32:
			zfs = append(zfs, zap.Uint32(key, field.Uint32()))
		case ink.FieldTypeUint16:
			zfs = append(zfs, zap.Uint16(key, field.Uint16()))
		case ink.FieldTypeUint8:
			zfs = append(zfs, zap.Uint8(key, field.Uint8()))
		case ink.FieldTypeUint:
			zfs = append(zfs, zap.Uint(key, field.Uint()))
		case ink.FieldTypeFloat64:
			zfs = append(zfs, zap.Float64(key, field.Float64()))
		case ink.FieldTypeFloat32:
			zfs = append(zfs, zap.Float32(key, field.Float32()))
		case ink.FieldTypeComplex128:
			zfs = append(zfs, zap.Complex128(key, field.Complex128()))
		case ink.FieldTypeComplex64:
			zfs = append(zfs, zap.Complex64(key, field.Complex64()))
		case ink.FieldTypeBool:
			zfs = append(zfs, zap.Bool(key, field.Bool()))
		case ink.FieldTypeTime:
			zfs = append(zfs, zap.Time(key, field.Time()))
		case ink.FieldTypeDuration:
			zfs = append(zfs, zap.Duration(key, field.Duration()))
		case ink.FieldTypeStringer:
			zfs = append(zfs, zap.Stringer(key, field.Stringer()))
		case ink.FieldTypeSlice:
			switch val := field.Slice().(type) {
			case []string:
				zfs = append(zfs, zap.Strings(key, val))
			case []int64:
				zfs = append(zfs, zap.Int64s(key, val))
			case []int32:
				zfs = append(zfs, zap.Int32s(key, val))
			case []int16:
				zfs = append(zfs, zap.Int16s(key, val))
			case []int8:
				zfs = append(zfs, zap.Int8s(key, val))
			case []int:
				zfs = append(zfs, zap.Ints(key, val))
			case []uint64:
				zfs = append(zfs, zap.Uint64s(key, val))
			case []uint32:
				zfs = append(zfs, zap.Uint32s(key, val))
			case []uint16:
				zfs = append(zfs, zap.Uint16s(key, val))
			case []uint8:
				zfs = append(zfs, zap.Uint8s(key, val))
			case []uint:
				zfs = append(zfs, zap.Uints(key, val))
			case []float64:
				zfs = append(zfs, zap.Float64s(key, val))
			case []float32:
				zfs = append(zfs, zap.Float32s(key, val))
			case []complex128:
				zfs = append(zfs, zap.Complex128s(key, val))
			case []complex64:
				zfs = append(zfs, zap.Complex64s(key, val))
			case []bool:
				zfs = append(zfs, zap.Bools(key, val))
			case []time.Time:
				zfs = append(zfs, zap.Times(key, val))
			case []time.Duration:
				zfs = append(zfs, zap.Durations(key, val))
			default:
				zfs = append(zfs, zap.Reflect(key, val))
			}
		case ink.FieldTypeReflect:
			zfs = append(zfs, zap.Any(key, field.Reflect()))
		default:
			zfs = append(zfs, zap.Reflect(key, field.Value()))
		}
	}
	return zfs
}
