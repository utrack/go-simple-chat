package logger

// Level is a logging level needed for a message to appear in the logs.
type Level uint64

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)
