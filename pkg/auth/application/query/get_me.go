package query

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/dto"
)

type GetMeHandler struct {
	staffHub domain.StaffHub
}

func NewGetMeHandler(staffHub domain.StaffHub) GetMeHandler {
	return GetMeHandler{
		staffHub: staffHub,
	}
}

func (h GetMeHandler) GetMe(ctx *appcontext.AppContext, performerID string, _ dto.GetMeRequest) (*dto.GetMeResponse, error) {
	ctx.Logger().Info("new get me request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("find staff in db")
	staff, err := h.staffHub.FindOneByID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to find staff", err, appcontext.Fields{})
		return nil, err
	}
	if staff == nil {
		ctx.Logger().ErrorText("staff not found")
		return nil, apperrors.Staff.StaffNotFound
	}

	ctx.Logger().Text("done get me request")
	return &dto.GetMeResponse{
		ID:   staff.ID,
		Name: staff.Name,
	}, nil
}
