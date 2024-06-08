package query

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/dto"
)

type GetAccessTokenByStaffIDHandler struct {
	jwtRepository domain.JwtRepository
}

func NewGetAccessTokenByStaffIDHandler(jwtRepository domain.JwtRepository) GetAccessTokenByStaffIDHandler {
	return GetAccessTokenByStaffIDHandler{
		jwtRepository: jwtRepository,
	}
}

func (h GetAccessTokenByStaffIDHandler) GetAccessTokenByStaffID(ctx *appcontext.AppContext, req dto.GetAccessTokenByStaffIDRequest) (*dto.GetAccessTokenByStaffIDResponse, error) {
	ctx.Logger().Info("new get access token by staff id request", appcontext.Fields{"staffID": req.StaffID})

	ctx.Logger().Text("validate staff id")
	if !database.IsValidObjectID(req.StaffID) {
		ctx.Logger().Error("invalid staff id", nil, appcontext.Fields{})
		return nil, apperrors.Staff.InvalidStaffID
	}

	ctx.Logger().Text("generate new access token")
	token, err := h.jwtRepository.GenerateAccessToken(ctx, req.StaffID)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Info("done get access token by staff id request", appcontext.Fields{"token": token})
	return &dto.GetAccessTokenByStaffIDResponse{AccessToken: token}, nil
}
