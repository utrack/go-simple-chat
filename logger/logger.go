//Package logger provides common abstraction level over various
//logging mechanisms.
package logger

// Logger represents an active logging object that restricts the messages
// based on their logging level and generates output somewhere.
type Logger func(Level, string, ...interface{})
