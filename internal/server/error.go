package server

import (
	"fmt"
	"net/http"
)

type HttpError interface {
	Error() string
	Code() int
}

type httpError struct {
	code    int
	message string
}

func (a httpError) Code() int {
	return a.code
}

func (a httpError) Error() string {
	return a.message
}

// HTTPErrorItemNotFound is a HttpError
func HTTPErrorItemNotFound(format string, msg ...interface{}) HttpError {
	return httpError{
		http.StatusNotFound,
		fmt.Sprintf(format, msg...),
	}
}

// HTTPErrorBadRequest is a HttpError
func HTTPErrorBadRequest(format string, msg ...interface{}) HttpError {
	return httpError{
		http.StatusBadRequest,
		fmt.Sprintf(format, msg...),
	}
}

// HTTPErrorInternal is a HttpError
func HTTPErrorInternal(format string, msg ...interface{}) HttpError {
	return httpError{
		http.StatusInternalServerError,
		fmt.Sprintf(format, msg...),
	}
}

// HTTPErrorUnauthorized is a HttpError
func HTTPErrorUnauthorized(format string, msg ...interface{}) HttpError {
	return httpError{
		http.StatusUnauthorized,
		fmt.Sprintf(format, msg...),
	}
}

// HTTPErrorMethodNotAllowed is a HttpError
func HTTPErrorMethodNotAllowed(format string, msg ...interface{}) HttpError {
	return httpError{
		http.StatusMethodNotAllowed,
		fmt.Sprintf(format, msg...),
	}
}

// HTTPErrorTimeout is a HttpError
func HTTPErrorTimeout(format string, msg ...interface{}) HttpError {
	return httpError{
		http.StatusRequestTimeout,
		fmt.Sprintf(format, msg...),
	}
}

// HTTPErrorExpectationFailed is a HttpError
func HTTPErrorExpectationFailed(format string, msg ...interface{}) HttpError {
	return httpError{
		http.StatusExpectationFailed,
		fmt.Sprintf(format, msg...),
	}
}
