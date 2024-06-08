package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	appjwt "github.com/namhq1989/vocab-booster-server-admin/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/application"
)

type server struct {
	app          application.Instance
	echo         *echo.Echo
	jwt          *appjwt.JWT
	isEnvRelease bool
}

func RegisterServer(_ *appcontext.AppContext, app application.Instance, e *echo.Echo, jwt *appjwt.JWT, isEnvRelease bool) error {
	var s = server{
		app:          app,
		echo:         e,
		jwt:          jwt,
		isEnvRelease: isEnvRelease,
	}

	s.registerAuthRoutes()

	return nil
}
