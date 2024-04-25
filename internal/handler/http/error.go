package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/wit-switch/assessment-tax/pkg/errorx"

	"github.com/labstack/echo/v4"
)

var StackCtxKey = struct{}{}

var (
	response = map[errorx.ErrCode]struct {
		message    string
		httpStatus int
	}{
		errorx.CodeValidationFail: {
			message:    msgValidationFail,
			httpStatus: http.StatusBadRequest,
		},
	}
)

func AsErrorResponse[T any](err error) *ResponseError[T] {
	builder := NewResponseErrorBuilder[T]()
	code := errorx.CodeUnknown

	if err == nil {
		return builder.Build()
	}

	if iErr := errorx.IsInternalErr[T](err); iErr != nil {
		code = iErr.Code()
		builder.WithError(iErr).WithParams(iErr.Params())
	}

	resp, ok := response[code]
	if !ok {
		resp.httpStatus = http.StatusInternalServerError
		resp.message = getErrorMessage(err)
	}

	return builder.
		WithMesssage(resp.message).
		WithCode(code.Int()).
		WithHTTPStatus(resp.httpStatus).
		Build()
}

func getErrorMessage(err error) string {
	if err != nil && err.Error() != "" {
		return err.Error()
	}

	return msgUnknown
}

func HTTPErrorHandler(err error, c echo.Context) {
	ctx := context.WithValue(c.Request().Context(), StackCtxKey, err)
	c.SetRequest(c.Request().WithContext(ctx))

	if c.Response().Committed {
		return
	}

	var errHTTP *echo.HTTPError
	if ok := errorx.As(err, &errHTTP); ok {
		var message string
		switch m := errHTTP.Message.(type) {
		case string:
			message = m
		case error:
			message = m.Error()
		}

		resp := NewResponseErrorBuilder[any]().
			WithMesssage(message).
			WithCode(errorx.CodeUnknown.Int()).
			WithHTTPStatus(errHTTP.Code).
			WithError(errHTTP).
			Build()

		if errResp := c.JSON(resp.StatusCode(), resp); errResp != nil {
			slog.ErrorContext(c.Request().Context(), "[!] failed to send response error", slog.Any("err", errResp))
		}

		return
	}

	resp := AsErrorResponse[any](err)
	if errResp := c.JSON(resp.StatusCode(), resp); errResp != nil {
		slog.ErrorContext(c.Request().Context(), "[!] failed to send response error", slog.Any("err", errResp))
	}
}
