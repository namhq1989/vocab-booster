package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
)

type StaffAuthTokenRepository interface {
	CreateAuthToken(ctx *appcontext.AppContext, token AuthToken) error
	DeleteAuthToken(ctx *appcontext.AppContext, tokenID string) error
	FindAuthToken(ctx *appcontext.AppContext, refreshToken string) (*AuthToken, error)
}

type Tokens struct {
	AccessToken        string
	RefreshToken       string
	AccessTokenExpiry  time.Time
	RefreshTokenExpiry time.Time
}

type AuthToken struct {
	ID           string
	StaffID      string
	RefreshToken string
	Expiry       time.Time
	CreatedAt    time.Time
}

func NewAuthToken(staffID, refreshToken string, expiry time.Time) (*AuthToken, error) {
	if staffID == "" {
		return nil, apperrors.Staff.InvalidStaffID
	}

	if refreshToken == "" {
		return nil, apperrors.Auth.InvalidRefreshToken
	}

	if expiry.IsZero() {
		return nil, apperrors.Auth.InvalidExpiry
	}

	return &AuthToken{
		ID:           database.NewStringID(),
		StaffID:      staffID,
		RefreshToken: refreshToken,
		Expiry:       expiry,
		CreatedAt:    time.Now(),
	}, nil
}
