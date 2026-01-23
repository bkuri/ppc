// Package error provides error types for PPC
package error

// SrcError represents an error with source context
type SrcError struct {
	Path string
	ID   string
	Msg  string
	Hint string
	Line int
}

// Error implements the error interface
func (e SrcError) Error() string {
	return e.Msg
}

// New creates a SrcError with path, id, and message
func New(path, id, msg string) SrcError {
	return SrcError{Path: path, ID: id, Msg: msg}
}

// NewAtLine creates a SrcError with line number
func NewAtLine(path, id, msg string, line int) SrcError {
	return SrcError{Path: path, ID: id, Msg: msg, Line: line}
}

// NewWithHint creates a SrcError with a hint
func NewWithHint(path, id, msg, hint string) SrcError {
	return SrcError{Path: path, ID: id, Msg: msg, Hint: hint}
}
