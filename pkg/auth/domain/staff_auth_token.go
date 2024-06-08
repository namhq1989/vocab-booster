package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
)

type StaffAuthTokenRepository interface {
	CreateAuthToken(ctx *appcontext.AppContext, token RefreshToken) error
	DeleteAuthToken(ctx *appcontext.AppContext, tokenID string) error
	FindAuthToken(ctx *appcontext.AppContext, refreshToken string) (*RefreshToken, error)
}

type Tokens struct {
	AccessToken        string
	RefreshToken       string
	AccessTokenExpiry  time.Time
	RefreshTokenExpiry time.Time
}

type RefreshToken struct {
	ID      string
	StaffID string
	Token   string
	Expiry  time.Time
}
