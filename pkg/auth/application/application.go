package application

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/application/command"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/application/query"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/dto"
)

type (
	Commands interface {
		SignInWithGoogle(ctx *appcontext.AppContext, req dto.SignInWithGoogleRequest) (*dto.SignInWithGoogleResponse, error)
		RefreshAccessToken(ctx *appcontext.AppContext, req dto.RefreshAccessTokenRequest) (*dto.RefreshAccessTokenResponse, error)
	}
	Queries interface {
		GetAccessTokenByStaffID(ctx *appcontext.AppContext, req dto.GetAccessTokenByStaffIDRequest) (*dto.GetAccessTokenByStaffIDResponse, error)
		GetMe(ctx *appcontext.AppContext, performerID string, _ dto.GetMeRequest) (*dto.GetMeResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	appCommandHandlers struct {
		command.SignInWithGoogleHandler
		command.RefreshAccessTokenHandler
	}
	appQueryHandler struct {
		query.GetAccessTokenByStaffIDHandler
		query.GetMeHandler
	}
	Application struct {
		appCommandHandlers
		appQueryHandler
	}
)

var _ Instance = (*Application)(nil)

func New(
	staffAuthTokenRepository domain.StaffAuthTokenRepository,
	ssoRepository domain.SSORepository,
	jwtRepository domain.JwtRepository,
	staffHub domain.StaffHub,
) *Application {
	return &Application{
		appCommandHandlers: appCommandHandlers{
			SignInWithGoogleHandler:   command.NewSignInWithGoogleHandler(staffAuthTokenRepository, ssoRepository, jwtRepository, staffHub),
			RefreshAccessTokenHandler: command.NewRefreshAccessTokenHandler(staffAuthTokenRepository, jwtRepository, staffHub),
		},
		appQueryHandler: appQueryHandler{
			GetAccessTokenByStaffIDHandler: query.NewGetAccessTokenByStaffIDHandler(jwtRepository),
			GetMeHandler:                   query.NewGetMeHandler(staffHub),
		},
	}
}
