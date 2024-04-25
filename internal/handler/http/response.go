package http

import (
	"fmt"
	"strconv"
)

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (b *BaseResponse) StatusCode() int {
	idx := 3
	suf := b.Code
	if len(suf) < 3 {
		idx = len(suf)
	}

	code := suf[:idx]
	i, err := strconv.ParseInt(code, 0, 32)
	if err != nil {
		return 500
	}

	return int(i)
}

type ResponseError[T any] struct {
	BaseResponse
	Errors T `json:"errors,omitempty"`
}

func (e *ResponseError[T]) Error() string {
	return e.Message
}

func ResponseCode(httpStatus, code int) string {
	return fmt.Sprintf("%d%03d", httpStatus, code)
}
