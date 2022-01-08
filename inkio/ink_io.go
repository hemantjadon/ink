// Package inkio defines basic implementation of ink.LogSink wrapping basic
// io.Writer.
package inkio

import (
	"fmt"
	"io"
	"sync"

	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/inkio/internal/istrconv"
	"github.com/hemantjadon/ink/internal/pool"
)

// Sink describes a basic ink.LogSink which writes the log lines to configured
// io.Writer.
type Sink struct {
	mu     sync.Mutex
	writer io.Writer
}

var _ ink.LogSink = (*Sink)(nil)

// NewSink creates a new Sink with the given io.Writer.
func NewSink(writer io.Writer) *Sink {
	if writer == nil {
		return nil
	}
	sink := Sink{writer: writer}
	return &sink
}

// Debug writes message and fields to writer with DEBUG prefix.
func (l *Sink) Debug(msg string, fields ...ink.Field) {
	l.mu.Lock()
	l.output(debugPrefix, msg, fields...)
	l.mu.Unlock()
}

// Info writes message and fields to writer with INFO prefix.
func (l *Sink) Info(msg string, fields ...ink.Field) {
	l.mu.Lock()
	l.output(infoPrefix, msg, fields...)
	l.mu.Unlock()
}

// Error writes message and fields to writer with ERROR prefix.
func (l *Sink) Error(msg string, fields ...ink.Field) {
	l.mu.Lock()
	l.output(errorPrefix, msg, fields...)
	l.mu.Unlock()
}

func (l *Sink) output(prefix, msg string, fields ...ink.Field) {
	msg = format(prefix, msg, fields)
	_, _ = fmt.Fprintln(l.writer, msg)
}

func format(prefix, msg string, fields []ink.Field) string {
	sb := pool.GetStringsBuilder()
	defer pool.PutStringsBuilder(sb)

	sb.WriteString(prefix)
	sb.WriteString(space)
	sb.WriteString(msg)

	for _, field := range fields {
		sb.WriteString(space)
		istrconv.AppendBuilderField(sb, field)
	}

	return sb.String()
}

const (
	debugPrefix = "DEBUG"
	infoPrefix  = "INFO"
	errorPrefix = "ERROR"
	space       = " "
)
