package validation

import (
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-admin/internal/utils/httprespond"
)

func ValidateHTTPPayload[T any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			req T
		)

		if err := c.Bind(&req); err != nil {
			return httprespond.R400(c, err, echo.Map{})
		}

		if v := validate.Struct(req); !v.Validate() {
			return httprespond.R400(c, v.Errors.OneError(), echo.Map{})
		}

		// assign to context
		c.Set("req", req)
		return next(c)
	}
}

func ValidateHTTPFormData[T any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			req T
		)

		file, err := c.FormFile("file")
		if err != nil {
			return httprespond.R400(c, err, echo.Map{})
		}

		if err = c.Bind(&req); err != nil {
			return httprespond.R400(c, err, echo.Map{})
		}

		if v := validate.Struct(req); !v.Validate() {
			return httprespond.R400(c, v.Errors.OneError(), echo.Map{})
		}

		// assign to context
		c.Set("req", req)
		c.Set("file", file)
		return next(c)
	}
}
