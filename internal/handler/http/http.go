package http

import (
	"net/http"
	"time"

	"github.com/wit-switch/assessment-tax/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func BindRoute[Query, Body, Response any](
	handler func(echo.Context, Query, Body) (Response, error),
	opts ...OptFunc,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := newOpt(opts...)
		start := time.Now().UTC()

		var query Query
		if opt.queryParser {
			if err := (&echo.DefaultBinder{}).BindQueryParams(c, &query); err != nil {
				return errorx.ErrValidationFail.WithError(err)
			}
		}

		var body Body
		if opt.bodyParser {
			if err := (&echo.DefaultBinder{}).BindBody(c, &body); err != nil {
				return errorx.ErrValidationFail.WithError(err)
			}

			if opt.bodyValidator {
				err := c.Validate(body)
				if err != nil {
					return err
				}
			}
		}

		response, err := handler(c, query, body)
		if err != nil {
			return err
		}

		c.Response().Header().Add("start", start.String())
		c.Response().Header().Add("duration", time.Since(start).String())
		return c.JSON(http.StatusOK, response)
	}
}
