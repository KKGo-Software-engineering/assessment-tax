package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	httphdl "github.com/wit-switch/assessment-tax/internal/handler/http"
	"github.com/wit-switch/assessment-tax/pkg/errorx"
	"github.com/wit-switch/assessment-tax/pkg/validator"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func request(method, route string, body io.Reader, e *echo.Echo) (int, string) {
	req := httptest.NewRequest(method, route, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	// custom error handler only plays with real request
	e.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}

var _ = Describe("HTTP", func() {
	Describe("bind route", func() {
		var (
			app   *echo.Echo
			route string
		)

		BeforeEach(func() {
			app = echo.New()
			route = "/test"

			app.Validator = httphdl.NewValidator(validator.New())
			app.HTTPErrorHandler = httphdl.HTTPErrorHandler
		})

		When("failed to handler request", func() {
			It("should return 500", func() {
				hdl := func(c echo.Context, _, _ any) ([]string, error) {
					return nil, errorx.New("error")
				}

				bindRoute := httphdl.BindRoute(hdl)
				app.GET(route, bindRoute)

				code, body := request(http.MethodGet, route, nil, app)

				Expect(500).To(Equal(code))
				actual, err := compacJSON(body)
				Expect(err).NotTo(HaveOccurred())

				expectedResp := `{
					"code":"500000",
					"message":"error"
				}`
				expected, err := compacJSON(expectedResp)
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal(expected))
			})
		})

		When("success to handler request", func() {
			Context("with query parser", func() {
				When("failed to parser query", func() {
					It("should return query zero value", func() {
						hdl := func(c echo.Context, query string, _ any) (*ok, error) {
							Expect(query).To(BeEmpty())
							return &ok{OK: true}, nil
						}

						bindRoute := httphdl.BindRoute(hdl, httphdl.WithQueryParser())
						app.GET(route, bindRoute)

						code, body := request(http.MethodGet, route+"?id=1", nil, app)
						Expect(200).To(Equal(code))
						var actual ok
						err := json.Unmarshal([]byte(body), &actual)
						Expect(err).NotTo(HaveOccurred())

						expected := ok{OK: true}
						Expect(actual).To(Equal(expected))
					})
				})

				It("should return query value", func() {
					hdl := func(c echo.Context, query item, _ any) (*ok, error) {
						expected := item{
							ID: "1",
						}
						Expect(query).To(Equal(expected))
						return &ok{OK: true}, nil
					}

					bindRoute := httphdl.BindRoute(hdl, httphdl.WithQueryParser())
					app.GET(route, bindRoute)

					code, body := request(http.MethodGet, route+"?id=1", nil, app)
					Expect(200).To(Equal(code))
					var actual ok
					err := json.Unmarshal([]byte(body), &actual)
					Expect(err).NotTo(HaveOccurred())

					expected := ok{OK: true}
					Expect(actual).To(Equal(expected))
				})
			})

			Context("with body parser", func() {
				When("failed to parser body", func() {
					It("should return body zero value", func() {
						hdl := func(c echo.Context, _ any, body item) (*ok, error) {
							return &ok{OK: true}, nil
						}

						bindRoute := httphdl.BindRoute(hdl, httphdl.WithBodyParser())
						app.POST(route, bindRoute)

						code, body := request(http.MethodPost, route, nil, app)
						Expect(200).To(Equal(code))
						var actual ok
						err := json.Unmarshal([]byte(body), &actual)
						Expect(err).NotTo(HaveOccurred())

						expected := ok{OK: true}
						Expect(actual).To(Equal(expected))
					})
				})

				It("should return body value", func() {
					hdl := func(c echo.Context, _ any, body item) (*ok, error) {
						expected := item{
							ID: "1",
						}
						Expect(body).To(Equal(expected))
						return &ok{OK: true}, nil
					}

					bindRoute := httphdl.BindRoute(hdl, httphdl.WithBodyParser())
					app.POST(route, bindRoute)

					code, body := request(http.MethodPost, route, createBody(item{
						ID: "1",
					}), app)
					Expect(200).To(Equal(code))
					var actual ok
					err := json.Unmarshal([]byte(body), &actual)
					Expect(err).NotTo(HaveOccurred())

					expected := ok{OK: true}
					Expect(actual).To(Equal(expected))
				})
			})
		})

		Context("with body parser and body validator", func() {
			When("failed to parser body", func() {
				It("should return validation error", func() {
					hdl := func(c echo.Context, _ any, body item) (*ok, error) {
						return &ok{OK: true}, nil
					}

					bindRoute := httphdl.BindRoute(hdl, httphdl.WithBodyParser(), httphdl.WithBodyValidator())
					app.POST(route, bindRoute)

					code, body := request(http.MethodPost, route, nil, app)
					Expect(400).To(Equal(code))
					actual, err := compacJSON(body)
					Expect(err).NotTo(HaveOccurred())

					expectedResp := `{
						"code":"400001",
						"message":"validation error",
						"errors":[{"field":"id","message":"id is required"}]
					}`
					expected, err := compacJSON(expectedResp)
					Expect(err).NotTo(HaveOccurred())
					Expect(actual).To(Equal(expected))
				})
			})

			It("should return body value", func() {
				hdl := func(c echo.Context, _ any, body item) (*ok, error) {
					expected := item{
						ID: "1",
					}
					Expect(body).To(Equal(expected))
					return &ok{OK: true}, nil
				}

				bindRoute := httphdl.BindRoute(hdl, httphdl.WithBodyParser(), httphdl.WithBodyValidator())
				app.POST(route, bindRoute)

				code, body := request(http.MethodPost, route, createBody(item{
					ID: "1",
				}), app)
				Expect(200).To(Equal(code))
				var actual ok
				err := json.Unmarshal([]byte(body), &actual)
				Expect(err).NotTo(HaveOccurred())

				expected := ok{OK: true}
				Expect(actual).To(Equal(expected))
			})
		})
	})
})

type ok struct {
	OK bool `json:"ok"`
}
type item struct {
	ID string `json:"id" query:"id" form:"id" validate:"required"`
}

func compacJSON(jsonStr string) ([]byte, error) {
	buffer := bytes.Buffer{}
	if err := json.Compact(&buffer, []byte(jsonStr)); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func createBody(v any) *bytes.Reader {
	body, _ := json.Marshal(v)
	return bytes.NewReader(body)
}
