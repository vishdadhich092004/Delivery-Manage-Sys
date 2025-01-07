package response

import (
	"log/slog"
	"runtime"
)

type AppError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      string `json:"error,omitempty"`
	StackTrace string `json:"-"`
}

// captures the stack trace
func captureStackTrace() string {
	stack := make([]byte, 1024)
	n := runtime.Stack(stack, false)
	return string(stack[:n])
}

// a new instance of AppError
func NewAppError(statusCode int, message string, err ...string) *AppError {
	errorMessage := ""
	if len(err) > 0 {
		errorMessage = err[0]
	}

	stack := captureStackTrace()
	slog.Error("AppError", "status", statusCode, "message", message, "error", errorMessage, "stack", stack)

	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Error:      errorMessage,
		StackTrace: stack,
	}
}

// AppError to a standardized API response
func (e *AppError) ToApiResponse() ApiResponseType {
	return ApiResponseType{
		Success: false,
		Message: e.Message,
		Data:    nil,
	}
}
