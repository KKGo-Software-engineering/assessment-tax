package middleware

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	httphdl "github.com/wit-switch/assessment-tax/internal/handler/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (m *Middleware) Logger() echo.MiddlewareFunc {
	return middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		ctx := c.Request().Context()

		var body, resp any
		_ = json.Unmarshal(reqBody, &body)
		_ = json.Unmarshal(resBody, &resp)

		method := c.Request().Method
		path := c.Request().URL.String()

		logMsg := fmt.Sprintf("[*] request method %s, path %s", method, path)
		logParams := []any{
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("status_code", c.Response().Status),
			slog.Any("body", body),
			slog.Any("resp", resp),
			slog.String("ip", c.RealIP()),
			slog.String("start", c.Response().Header().Get("start")),
			slog.String("duration", c.Response().Header().Get("duration")),
		}

		if stackErr, ok := ctx.Value(httphdl.StackCtxKey).(error); ok && stackErr != nil {
			logParams = append(logParams, slog.Any("error", stackErr))
		}

		switch statusCode := c.Response().Status; {
		case statusCode >= http.StatusInternalServerError:
			slog.ErrorContext(ctx, logMsg, logParams...)
		case statusCode >= http.StatusBadRequest:
			slog.WarnContext(ctx, logMsg, logParams...)
		default:
			slog.InfoContext(ctx, logMsg, logParams...)
		}
	})
}
