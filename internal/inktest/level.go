package inktest

// Level defines a log Level.
type Level string

const (
	// LevelDebug is log Level for DEBUG logs.
	LevelDebug Level = "DEBUG"

	// LevelInfo is log Level for INFO logs.
	LevelInfo Level = "INFO"

	// LevelError is log Level for ERROR logs.
	LevelError Level = "ERROR"
)

func (l Level) String() string {
	return string(l)
}
