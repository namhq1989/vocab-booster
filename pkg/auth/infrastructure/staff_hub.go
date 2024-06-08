package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/genproto/staffpb"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
)

type StaffHub struct {
	client staffpb.StaffServiceClient
}

func NewStaffHub(client staffpb.StaffServiceClient) *StaffHub {
	return &StaffHub{
		client: client,
	}
}

func (r StaffHub) FindOneByEmail(ctx *appcontext.AppContext, email string) (*domain.Staff, error) {
	resp, err := r.client.FindUserByEmail(ctx.Context(), &staffpb.FindStaffByEmailRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	staff := resp.GetStaff()
	if staff == nil {
		return nil, nil
	}
	return &domain.Staff{
		ID:      staff.GetId(),
		Name:    staff.GetName(),
		Email:   staff.GetEmail(),
		IsAdmin: staff.GetIsAdmin(),
	}, nil
}

func (r StaffHub) FindOneByID(ctx *appcontext.AppContext, id string) (*domain.Staff, error) {
	resp, err := r.client.FindUserByID(ctx.Context(), &staffpb.FindStaffByIDRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	staff := resp.GetStaff()
	if staff == nil {
		return nil, nil
	}
	return &domain.Staff{
		ID:      staff.GetId(),
		Name:    staff.GetName(),
		Email:   staff.GetEmail(),
		IsAdmin: staff.GetIsAdmin(),
	}, nil
}
