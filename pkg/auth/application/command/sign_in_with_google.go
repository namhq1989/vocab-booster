package command

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/dto"
)

type SignInWithGoogleHandler struct {
	staffAuthTokenRepository domain.StaffAuthTokenRepository
	ssoRepository            domain.SSORepository
	jwtRepository            domain.JwtRepository
	staffHub                 domain.StaffHub
}

func NewSignInWithGoogleHandler(staffAuthTokenRepository domain.StaffAuthTokenRepository, ssoRepository domain.SSORepository, jwtRepository domain.JwtRepository, staffHub domain.StaffHub) SignInWithGoogleHandler {
	return SignInWithGoogleHandler{
		staffAuthTokenRepository: staffAuthTokenRepository,
		jwtRepository:            jwtRepository,
		ssoRepository:            ssoRepository,
		staffHub:                 staffHub,
	}
}

func (h SignInWithGoogleHandler) SignInWithGoogle(ctx *appcontext.AppContext, req dto.SignInWithGoogleRequest) (*dto.SignInWithGoogleResponse, error) {
	ctx.Logger().Info("new sign in with Google request", appcontext.Fields{"token": req.Token})

	ctx.Logger().Text("get staff data with Google token")
	googleUser, err := h.ssoRepository.GetUserDataWithGoogleToken(ctx, req.Token)
	if err != nil {
		ctx.Logger().Error("failed to get staff data with Google token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Info("find staff with email in db", appcontext.Fields{"email": googleUser.Email})
	staff, err := h.staffHub.FindOneByEmail(ctx, googleUser.Email)
	if err != nil {
		ctx.Logger().Error("failed to find staff by email", err, appcontext.Fields{})
		return nil, err
	}
	if staff == nil {
		ctx.Logger().ErrorText("staff not found")
		return nil, apperrors.Staff.StaffNotFound
	}

	ctx.Logger().Info("staff found, generate token", appcontext.Fields{"staffID": staff.ID})
	generatedTokens, err := h.jwtRepository.GenerateTokens(ctx, staff.ID)
	if err != nil {
		ctx.Logger().Error("failed to generate tokens", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist refresh token to database")
	err = h.staffAuthTokenRepository.CreateAuthToken(ctx, domain.RefreshToken{
		StaffID: staff.ID,
		Token:   generatedTokens.RefreshToken,
		Expiry:  generatedTokens.RefreshTokenExpiry,
	})
	if err != nil {
		ctx.Logger().Error("failed to persist refresh token", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("generate response's tokens data")
	tokens := &domain.Tokens{
		AccessToken:  generatedTokens.AccessToken,
		RefreshToken: generatedTokens.RefreshToken,
	}

	ctx.Logger().Text("done sign in with Google request")
	return &dto.SignInWithGoogleResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
