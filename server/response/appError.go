package response

import (
	"log/slog"
	"runtime"
)

type AppError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      string `json:"error,omitempty"`
}

// captureStackTrace captures the stack trace
func captureStackTrace() string {
	stack := make([]byte, 1024)
	n := runtime.Stack(stack, false)
	return string(stack[:n])
}

func NewAppError(statusCode int, message string, err ...string) *AppError {
	stack := captureStackTrace()
	slog.Error(stack)
	error := ""
	if len(err) > 0 {
		error = err[0]
	}
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Error:      error,
	}
}
