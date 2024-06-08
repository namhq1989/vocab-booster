package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	appjwt "github.com/namhq1989/vocab-booster-server-admin/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/application"
)

type server struct {
	app  application.Instance
	echo *echo.Echo
	jwt  *appjwt.JWT
}

func RegisterServer(_ *appcontext.AppContext, app application.Instance, e *echo.Echo, jwt *appjwt.JWT) error {
	var s = server{
		app:  app,
		echo: e,
		jwt:  jwt,
	}

	s.registerAuditRoutes()

	return nil
}
