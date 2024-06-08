package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-admin/internal/utils/httprespond"
)

func (s server) registerAuditRoutes() {
	g := s.echo.Group("/api/audit")

	g.GET("", func(c echo.Context) error {
		return httprespond.R200(c, echo.Map{"success": true})
	})
}
