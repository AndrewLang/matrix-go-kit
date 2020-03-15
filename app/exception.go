package app

import (
	"time"
)

// Exception describe an exception
type Exception struct {
	Message        string
	StackTrace     string
	InnerException *Exception
	Date           time.Time
}

func (e Exception) String() string {
	return e.Message
}

func (e Exception) Error() string {
	return e.Message
}

// NewException create
func NewException(message string) *Exception {
	return &Exception{
		Message: message,
		Date:    time.Now(),
	}
}

// NewExceptionWithInner create
func NewExceptionWithInner(message string, innerException *Exception) *Exception {
	return &Exception{
		Message:        message,
		Date:           time.Now(),
		InnerException: innerException,
	}
}
