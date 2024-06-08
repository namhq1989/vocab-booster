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
}

func (m AuthToken) ToDomain() domain.RefreshToken {
	return domain.RefreshToken{
		ID:      m.ID.Hex(),
		StaffID: m.StaffID.Hex(),
		Token:   m.RefreshToken,
		Expiry:  m.Expiry,
	}
}

func (AuthToken) FromDomain(token domain.RefreshToken) (*AuthToken, error) {
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
		RefreshToken: token.Token,
		Expiry:       token.Expiry,
	}, nil
}
