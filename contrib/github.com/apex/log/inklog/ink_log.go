// Package inklog provides compatibility layer between ink and apex/log
// package. See https://github.com/apex/log for details about apex log.
package inklog

import (
	apexlog "github.com/apex/log"

	"github.com/hemantjadon/ink"
)

// Sink describes a apexlog.Logger based ink.LogSink.
type Sink struct {
	logger apexlog.Interface
	fields []ink.Field
}

var (
	_ ink.LogSink       = (*Sink)(nil)
	_ ink.FieldsLogSink = (*Sink)(nil)
)

// NewSink wraps the given apexlog.Interface into an ink.LogSink.
func NewSink(logger apexlog.Interface) *Sink {
	if logger == nil {
		return nil
	}
	sink := Sink{logger: logger}
	return &sink
}

// Debug logs to underlying apexlog.Logger at DEBUG level.
func (l Sink) Debug(msg string, fields ...ink.Field) {
	l.logger.WithFields(convertFields(fields)).Debug(msg)
}

// Info logs to underlying apexlog.Logger at INFO level.
func (l Sink) Info(msg string, fields ...ink.Field) {
	l.logger.WithFields(convertFields(fields)).Info(msg)
}

// Error logs to underlying apexlog.Logger at ERROR level.
func (l Sink) Error(msg string, fields ...ink.Field) {
	l.logger.WithFields(convertFields(fields)).Error(msg)
}

func (s Sink) WithFields(fields ...ink.Field) ink.FieldsLogSink {
	if len(fields) == 0 {
		return s
	}
	sink := s.clone()
	sink.logger = s.logger.WithFields(convertFields(fields))
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

func convertFields(fields []ink.Field) apexlog.Fields {
	lfs := make(apexlog.Fields, len(fields))
	for _, field := range fields {
		key := field.Key()
		value := field.Value()
		lfs[key] = value
	}
	return lfs
}
