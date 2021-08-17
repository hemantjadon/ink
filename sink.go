package ink

// LogSink is the core interface which defines the sinks of the logs.
//
// LogSink implementations must be concurrency safe, as sinks should be sharable
// between multiple loggers.
type LogSink interface {
	// Debug logs the message at DEBUG level.
	Debug(msg string, fields ...Field)

	// Info logs the message at INFO level.
	Info(msg string, fields ...Field)

	// Error logs the message at ERROR level.
	Error(msg string, fields ...Field)
}

// NameLogSink defines the sinks which support the named logger.
//
// NameLogSink must be chainable, i.e. if a named sink "logger1" is extended
// with a name "logger2", then the extended logger name must be
// "logger1.logger2".
type NameLogSink interface {
	LogSink

	// WithName gives a new LogSink with given name added to it. It must not
	// modify the existing LogSink but return a new one.
	WithName(name string) NameLogSink

	// Name gets the configured name of the LogSink.
	Name() string
}

// FieldsLogSink defines the sinks which support the logger level sinks.
//
// FieldsLogSink must be chainable, i.e. if a named sink has fields
// "[field1:value1, field2:value2]" is extended with "[field3:value3]", then the
// extended logger must have fields
// "[field1:value1, field2:value2, field3:value3]".
type FieldsLogSink interface {
	LogSink

	// WithFields gives a new LogSink with given fields added to it. It must
	// not modify the current LogSink but return a new one.
	WithFields(fields ...Field) FieldsLogSink

	// Fields gets the configured fields of the LogSink.
	Fields() []Field
}
