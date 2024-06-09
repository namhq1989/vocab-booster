package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/staff/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Staff struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Role      string             `bson:"role"`
	Status    string             `bson:"status"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func (m Staff) ToDomain() domain.Staff {
	return domain.Staff{
		ID:        m.ID.Hex(),
		Name:      m.Name,
		Email:     m.Email,
		Role:      domain.ToStaffRole(m.Role),
		Status:    domain.ToStaffStatus(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (Staff) FromDomain(staff domain.Staff) (*Staff, error) {
	id, err := database.ObjectIDFromString(staff.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	return &Staff{
		ID:        id,
		Name:      staff.Name,
		Email:     staff.Email,
		Role:      staff.Role.String(),
		Status:    staff.Status.String(),
		CreatedAt: staff.CreatedAt,
		UpdatedAt: staff.UpdatedAt,
	}, nil
}
