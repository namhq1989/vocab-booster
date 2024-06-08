package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/utils/sso"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
)

type SSORepository struct {
	googleClientID     string
	googleClientSecret string
}

func NewSSORepository(googleClientID, googleClientSecret string) SSORepository {
	return SSORepository{
		googleClientID:     googleClientID,
		googleClientSecret: googleClientSecret,
	}
}

func (r SSORepository) GetUserDataWithGoogleToken(ctx *appcontext.AppContext, token string) (*domain.SSOGoogleUser, error) {
	googleUser, err := sso.GetUserDataWithGoogleToken(ctx, r.googleClientID, token)
	if err != nil {
		return nil, err
	}

	return &domain.SSOGoogleUser{
		ID:    googleUser.ID,
		Email: googleUser.Email,
		Name:  googleUser.Name,
	}, nil
}
