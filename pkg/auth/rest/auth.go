package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-admin/internal/utils/validation"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/dto"
)

func (s server) registerAuthRoutes() {
	g := s.echo.Group("/api/auth")

	g.POST("/sign-in-with-google", func(c echo.Context) error {
		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.SignInWithGoogleRequest)
		)

		resp, err := s.app.SignInWithGoogle(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.SignInWithGoogleRequest](next)
	})

	g.POST("/refresh-access-token", func(c echo.Context) error {
		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.RefreshAccessTokenRequest)
		)

		resp, err := s.app.RefreshAccessToken(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.RefreshAccessTokenRequest](next)
	})

	g.GET("/access-token", func(c echo.Context) error {
		if s.isEnvRelease {
			return httprespond.R404(c, nil, nil)
		}

		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.GetAccessTokenByStaffIDRequest)
		)

		resp, err := s.app.GetAccessTokenByStaffID(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetAccessTokenByStaffIDRequest](next)
	})

	g.GET("/me", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetMeRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetMe(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetMeRequest](next)
	})
}
