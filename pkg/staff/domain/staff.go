package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
)

type StaffRepository interface {
	CreateStaff(ctx *appcontext.AppContext, staff Staff) error
	FindStaffByEmail(ctx *appcontext.AppContext, email string) (*Staff, error)
	CountByEmail(ctx *appcontext.AppContext, email string) (int64, error)
}

type Staff struct {
	ID        string
	Name      string
	Email     string
	Role      StaffRole
	Status    StaffStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStaff(name, email, role, status string) (*Staff, error) {
	if name == "" {
		return nil, apperrors.Common.InvalidName
	}

	if email == "" {
		return nil, apperrors.Common.InvalidEmail
	}

	dRole := ToStaffRole(role)
	if !dRole.IsValid() {
		return nil, apperrors.Common.InvalidRole
	}

	dStatus := ToStaffStatus(status)
	if !dStatus.IsValid() {
		return nil, apperrors.Common.InvalidStatus
	}

	return &Staff{
		ID:        database.NewStringID(),
		Name:      name,
		Email:     email,
		Role:      dRole,
		Status:    dStatus,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
