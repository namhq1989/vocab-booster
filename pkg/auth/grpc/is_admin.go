package grpc

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/genproto/authpb"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
)

type IsAdminHandler struct {
	staffHub domain.StaffHub
}

func NewIsAdminHandler(staffHub domain.StaffHub) IsAdminHandler {
	return IsAdminHandler{
		staffHub: staffHub,
	}
}

func (h IsAdminHandler) IsAdmin(ctx *appcontext.AppContext, req *authpb.IsAdminRequest) (*authpb.IsAdminResponse, error) {
	ctx.Logger().Info("new check is_admin request", appcontext.Fields{"id": req.GetId()})

	ctx.Logger().Text("find staff by id with grpc")
	staff, err := h.staffHub.FindOneByID(ctx, req.GetId())
	if err != nil {
		ctx.Logger().Error("failed to find staff with grpc", err, appcontext.Fields{})
		return nil, err
	}
	if staff == nil {
		ctx.Logger().ErrorText("staff not found")
		return nil, apperrors.Staff.StaffNotFound
	}

	ctx.Logger().Text("done check is_admin request")
	return &authpb.IsAdminResponse{IsAdmin: staff.IsAdmin}, nil
}
