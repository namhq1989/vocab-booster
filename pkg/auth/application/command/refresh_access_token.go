package command

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/dto"
)

type RefreshAccessTokenHandler struct {
	staffAuthTokenRepository domain.StaffAuthTokenRepository
	jwtRepository            domain.JwtRepository
	staffHub                 domain.StaffHub
}

func NewRefreshAccessTokenHandler(staffAuthTokenRepository domain.StaffAuthTokenRepository, jwtRepository domain.JwtRepository, staffHub domain.StaffHub) RefreshAccessTokenHandler {
	return RefreshAccessTokenHandler{
		staffAuthTokenRepository: staffAuthTokenRepository,
		jwtRepository:            jwtRepository,
		staffHub:                 staffHub,
	}
}

func (h RefreshAccessTokenHandler) RefreshAccessToken(ctx *appcontext.AppContext, req dto.RefreshAccessTokenRequest) (*dto.RefreshAccessTokenResponse, error) {
	ctx.Logger().Info("new refresh access token request", appcontext.Fields{"refreshToken": req.RefreshToken})

	ctx.Logger().Text("find auth token in db")
	authToken, err := h.staffAuthTokenRepository.FindAuthToken(ctx, req.RefreshToken)
	if err != nil {
		ctx.Logger().Error("failed to find auth token", err, appcontext.Fields{})
		return nil, err
	}
	if authToken == nil {
		ctx.Logger().ErrorText("auth token not found")
		return nil, apperrors.Auth.InvalidAuthToken
	}

	ctx.Logger().Info("auth token found, check expiration time", appcontext.Fields{"id": authToken.ID, "expiry": authToken.Expiry.String()})
	if authToken.Expiry.Before(time.Now()) {
		ctx.Logger().Text("auth token is already expired, delete and respond")
		defer func() { _ = h.staffAuthTokenRepository.DeleteAuthToken(ctx, authToken.ID) }()
		return nil, apperrors.Auth.InvalidAuthToken
	}

	ctx.Logger().Info("auth token is still valid, generate new access token", appcontext.Fields{})
	if accessToken, err := h.jwtRepository.GenerateAccessToken(ctx, authToken.StaffID); err != nil {
		ctx.Logger().Error("failed to generate new access token", err, appcontext.Fields{})
		return nil, err
	} else {
		ctx.Logger().Text("done refresh access token request")
		return &dto.RefreshAccessTokenResponse{AccessToken: accessToken}, nil
	}
}
