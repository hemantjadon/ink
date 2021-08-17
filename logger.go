package ink

import (
	"strings"
)

// Logger is the core logger struct.
type Logger struct {
	// sink of the logger, to which all the logs are written.
	sink LogSink

	// name of the logger, will be non-empty if and only if sink does not have
	// support for named loggers.
	name string

	// fields of the logger, will be non-nil if and only if sink does not have
	// support for loggers with logger level fields.
	fields []Field
}

// NewLogger creates a new Logger with given options.
func NewLogger(opts ...Option) *Logger {
	var logger Logger
	for _, opt := range opts {
		opt(&logger)
	}
	return &logger
}

// Option provide the ability to customize the Logger being created or extended.
type Option func(l *Logger)

// WithName gives an Option to add the name to the logger.
//
// If the sink of the Logger supports named loggers by implementing NameLogSink
// interface, then WithName uses it, otherwise name is added to the fields while
// logging.
func WithName(name string) Option {
	return func(l *Logger) {
		currentName := l.getName()
		currentFields := l.getFields()
		newSink := l.sink
		if nls, ok := newSink.(NameLogSink); ok {
			newSink = nls.WithName(name)
			l.name = ""
		} else {
			if len(currentName) != 0 {
				l.name = strings.Join([]string{currentName, name}, ".")
			} else {
				l.name = name
			}
		}
		if fls, ok := newSink.(FieldsLogSink); ok {
			newSink = fls.WithFields(currentFields...)
			l.fields = nil
		} else {
			l.fields = currentFields
		}
		l.sink = newSink
	}
}

// WithFields gives an Option to add the fields to the logger.
//
// If the sink for the Logger supports logger level fields by implementing
// FieldsLogSink interface then WithFields uses it, otherwise fields are added
// to the collection of existing fields.
func WithFields(fields ...Field) Option {
	return func(l *Logger) {
		currentName := l.getName()
		currentFields := l.getFields()
		newSink := l.sink
		if fls, ok := newSink.(FieldsLogSink); ok {
			newSink = fls.WithFields(fields...)
			l.fields = nil
		} else {
			if len(currentFields) != 0 {
				fs := make([]Field, 0, len(currentFields)+len(fields))
				fs = append(fs, currentFields...)
				fs = append(fs, fields...)
				l.fields = fs
			} else {
				l.fields = fields
			}
		}
		if nls, ok := newSink.(NameLogSink); ok {
			newSink = nls.WithName(currentName)
			l.name = ""
		} else {
			l.name = currentName
		}
		l.sink = newSink
	}
}

// WithSink gives an Option to change the sink of the logger.
//
// It applies current name and fields to the new sink as well.
//
// If the sink supports named loggers by implementing NameLogSink interface,
// then WithSink uses it for applying current name to the sink.
//
// If the sink supports logger level fields by implementing FieldsLogSink
// interface then WithSink uses it for applying current fields to the sink
func WithSink(sink LogSink) Option {
	return func(l *Logger) {
		currentName := l.getName()
		currentFields := l.getFields()
		newSink := sink
		if nls, ok := newSink.(NameLogSink); ok {
			sink = nls.WithName(currentName)
			l.name = ""
		} else {
			l.name = currentName
		}
		if fls, ok := newSink.(FieldsLogSink); ok {
			sink = fls.WithFields(currentFields...)
			l.fields = nil
		} else {
			l.fields = currentFields
		}
		l.sink = newSink
	}
}

// Extend creates a new Logger similar to the logger, with different options.
//
// The sink, name and fields of the extended logger remains same as the current,
// Logger unless overridden by any given option.
func (l Logger) Extend(opts ...Option) *Logger {
	logger := l.clone()
	for _, opt := range opts {
		opt(&logger)
	}
	return &logger
}

// Debug logs the message at DEBUG level using LogSink's Debug method.
func (l Logger) Debug(msg string, fields ...Field) {
	if l.sink == nil {
		return
	}
	l.sink.Debug(msg, l.merge(fields)...)
}

// Info logs the message at INFO level using LogSink's Info method.
func (l Logger) Info(msg string, fields ...Field) {
	if l.sink == nil {
		return
	}
	l.sink.Info(msg, l.merge(fields)...)
}

// Error logs the message at DEBUG level using LogSink's Error method.
func (l Logger) Error(msg string, fields ...Field) {
	if l.sink == nil {
		return
	}
	l.sink.Error(msg, l.merge(fields)...)
}

// getName gets the configured name of the Logger.
// If current sink supports named loggers by implementing NameLogSink interface,
// then name is taken from the sink, otherwise logger's name is returned.
func (l Logger) getName() string {
	name := l.name
	if len(name) == 0 {
		if nl, ok := l.sink.(NameLogSink); ok {
			name = nl.Name()
		}
	}
	return name
}

// getFields gets the configured fields of the Logger
//
// If current sink supports logger level fields by implementing FieldsLogSink
// interface, then name is taken from the sink, otherwise logger's name is
// returned.
func (l Logger) getFields() []Field {
	fs := l.fields
	if len(fs) == 0 {
		if fl, ok := l.sink.(FieldsLogSink); ok {
			fs = fl.Fields()
		}
	}
	return fs
}

// clone returns a new logger with same sink, name and fields.
func (l Logger) clone() Logger {
	logger := Logger{sink: l.sink, name: l.name}
	if len(l.fields) == 0 {
		return logger
	}
	logger.fields = make([]Field, 0, len(l.fields))
	logger.fields = append(logger.fields, l.fields...)
	return logger
}

// merge combines the given fields with logger's fields and name.
//
// If Logger l has no fields and no name configured then given fields are
// returned.
//
// If Logger l has either fields or name or both configured then given fields
// are merged with configured fields and na me, allocating a new slice.
func (l Logger) merge(fields []Field) []Field {
	switch {
	case len(l.fields) == 0 && len(l.name) == 0:
		return fields
	case len(l.fields) != 0 && len(l.name) == 0:
		fs := make([]Field, 0, len(l.fields)+len(fields))
		fs = append(fs, l.fields...)
		fs = append(fs, fields...)
		return fs
	case len(l.fields) == 0 && len(l.name) != 0:
		fs := make([]Field, 0, len(fields)+1)
		fs = append(fs, String("logger", l.name))
		fs = append(fs, fields...)
		return fs
	case len(l.fields) != 0 && len(l.name) != 0:
		fs := make([]Field, 0, len(l.fields)+len(fields)+1)
		fs = append(fs, String("logger", l.name))
		fs = append(fs, l.fields...)
		fs = append(fs, fields...)
		return fs
	default:
		return nil
	}
}
