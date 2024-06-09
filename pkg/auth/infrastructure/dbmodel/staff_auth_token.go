package dbmodel

import (
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthToken struct {
	ID           primitive.ObjectID `bson:"_id"`
	StaffID      primitive.ObjectID `bson:"staffId"`
	RefreshToken string             `bson:"refreshToken"`
	Expiry       time.Time          `bson:"expiry"`
	CreatedAt    time.Time          `bson:"createdAt"`
}

func (m AuthToken) ToDomain() domain.AuthToken {
	return domain.AuthToken{
		ID:           m.ID.Hex(),
		StaffID:      m.StaffID.Hex(),
		RefreshToken: m.RefreshToken,
		Expiry:       m.Expiry,
		CreatedAt:    m.CreatedAt,
	}
}

func (AuthToken) FromDomain(token domain.AuthToken) (*AuthToken, error) {
	id, err := database.ObjectIDFromString(token.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidID
	}

	staffID, err := database.ObjectIDFromString(token.StaffID)
	if err != nil {
		return nil, apperrors.Staff.InvalidStaffID
	}

	return &AuthToken{
		ID:           id,
		StaffID:      staffID,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
		CreatedAt:    token.CreatedAt,
	}, nil
}
