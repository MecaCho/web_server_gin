package types

import (
	"fmt"
	"strconv"
)

type ErrorRespI interface {
	Error() string
}

type ErrorResponse struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func NewErrorResponse(code int, msg string) ErrorRespI {
	return &ErrorResponse{strconv.FormatInt(int64(code), 10), msg}
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("error_code = %s; error_msg = %s", e.ErrorCode, e.ErrorMessage)
}
