package retry

import (
	"fmt"
)

type IRetryable interface {
	Retryable() bool
}

type retryableError struct {
	wrapped error
}

// MarkRetryable wraps the given error, marking it as retryable.
func MarkRetryable(e error) error {
	return &retryableError{wrapped: e}
}

// Error implements the error interface.
func (e *retryableError) Error() string {
	return fmt.Sprintf("retryable: %s", e.wrapped.Error())
}

// Retryable implements the error interface.
func (e *retryableError) Retryable() bool {
	return true
}
