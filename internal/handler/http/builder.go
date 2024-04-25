package http

import (
	"net/http"

	"github.com/wit-switch/assessment-tax/pkg/errorx"
)

type ResponseErrorBuilder[T any] struct {
	message    string
	code       int
	httpStatus int
	err        error
	params     T
}

func NewResponseErrorBuilder[T any]() *ResponseErrorBuilder[T] {
	return &ResponseErrorBuilder[T]{
		message:    msgUnknown,
		code:       errorx.CodeUnknown.Int(),
		httpStatus: http.StatusInternalServerError,
	}
}

func (b *ResponseErrorBuilder[T]) WithMesssage(messsage string) *ResponseErrorBuilder[T] {
	b.message = messsage
	return b
}

func (b *ResponseErrorBuilder[T]) WithCode(code int) *ResponseErrorBuilder[T] {
	b.code = code
	return b
}

func (b *ResponseErrorBuilder[T]) WithHTTPStatus(httpStatus int) *ResponseErrorBuilder[T] {
	b.httpStatus = httpStatus
	return b
}

func (b *ResponseErrorBuilder[T]) WithError(err error) *ResponseErrorBuilder[T] {
	b.err = err
	return b
}

func (b *ResponseErrorBuilder[T]) WithParams(params T) *ResponseErrorBuilder[T] {
	b.params = params
	return b
}

func (b *ResponseErrorBuilder[T]) Build() *ResponseError[T] {
	msg := b.message
	if msg == "" {
		msg = msgUnknown

		if b.code == errorx.CodeValidationFail.Int() {
			msg = msgValidationFail
		} else if b.err != nil {
			msg = b.err.Error()
		}
	}

	return &ResponseError[T]{
		BaseResponse: BaseResponse{
			Code:    ResponseCode(b.httpStatus, b.code),
			Message: msg,
		},
		Errors: b.params,
	}
}
