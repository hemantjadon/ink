// Package inklog provides compatibility layer between ink and standard
// library's log package. See https://pkg.go.dev/log for details about log.
package inklog

import (
	"fmt"
	"log"
	"sync"

	"github.com/hemantjadon/ink"
	"github.com/hemantjadon/ink/contrib/log/inklog/internal/istrconv"
	"github.com/hemantjadon/ink/internal/pool"
)

// Sink describes a standard library based ink.LogSink.
type Sink struct {
	mu          sync.Mutex
	debugLogger *log.Logger
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

var _ ink.LogSink = (*Sink)(nil)

// NewSink wraps the given log.Logger into an ink.LogSink.
func NewSink(logger *log.Logger) *Sink {
	if logger == nil {
		return nil
	}

	writer := logger.Writer()
	prefix := logger.Prefix()
	flags := logger.Flags()

	sink := Sink{
		debugLogger: log.New(writer, fmt.Sprintf("%s %s ", prefix, debugPrefix), flags),
		infoLogger:  log.New(writer, fmt.Sprintf("%s %s ", prefix, infoPrefix), flags),
		errorLogger: log.New(writer, fmt.Sprintf("%s %s ", prefix, errorPrefix), flags),
	}

	return &sink
}

// Debug logs the message and fields to underlying log.Logger with DEBUG prefix.
func (s *Sink) Debug(msg string, fields ...ink.Field) {
	s.mu.Lock()
	s.debugLogger.Print(format(msg, fields))
	s.mu.Unlock()
}

// Info logs the message and fields to underlying log.Logger with INFO prefix.
func (s *Sink) Info(msg string, fields ...ink.Field) {
	s.mu.Lock()
	s.infoLogger.Print(format(msg, fields))
	s.mu.Unlock()
}

// Error logs the message and fields to underlying log.Logger with ERROR prefix.
func (s *Sink) Error(msg string, fields ...ink.Field) {
	s.mu.Lock()
	s.errorLogger.Print(format(msg, fields))
	s.mu.Unlock()
}

func format(msg string, fields []ink.Field) string {
	if len(fields) == 0 {
		return msg
	}

	sb := pool.GetStringsBuilder()
	defer pool.PutStringsBuilder(sb)

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
