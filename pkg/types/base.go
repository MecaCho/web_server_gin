package types

import (
	"fmt"
)

type ErrorRespI interface {
	Error() string
}

type ErrorResponse struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func NewErrorResponse(code int, msg string) ErrorRespI {
	return &ErrorResponse{fmt.Sprintf("BWS.%d", code), msg}
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("error_code = %s; error_msg = %s", e.ErrorCode, e.ErrorMessage)
}

// NewNotFoundResp ...
func NewNotFoundResp(code int, resourceID int64, messages ...string) (err ErrorRespI) {
	var message string
	if len(messages) == 0 {
		message = fmt.Sprintf("Resource not found, id : (%d)", resourceID)
	} else {
		for _, msg := range messages {
			message = message + msg
		}
	}
	return NewErrorResponse(code, message)
}

// NewInternalErrorResp ...
func NewInternalErrorResp(code int, messages ...string) (err ErrorRespI) {
	var message string

	for _, msg := range messages {
		message = message + msg
	}
	return NewErrorResponse(code, fmt.Sprintf("internal error : %s", message))
}

// BadRequestResp ...
// The client request contains a syntax error or mistake, unable to complete your request
//
// swagger:response
type BadRequestResp struct {
	// The client request contains a syntax error or mistake, unable to complete your request
	//
	// in:body
	Body ErrorResponse
}

// InternalError ...
// Server error, server error in processing requests
//
// swagger:response
type InternalError struct {
	// Server error, server error in processing requests
	//
	// in:body
	Body ErrorResponse
}
