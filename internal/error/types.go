// Package error provides error types for PPC
package error

// SrcError represents an error with source context
type SrcError struct {
	Path string
	ID   string
	Msg  string
}

// Error implements the error interface
func (e SrcError) Error() string {
	return e.Msg
}

// New creates a SrcError with path, id, and message
func New(path, id, msg string) SrcError {
	return SrcError{Path: path, ID: id, Msg: msg}
}
