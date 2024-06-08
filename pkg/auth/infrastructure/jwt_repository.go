package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	appjwt "github.com/namhq1989/vocab-booster-server-admin/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
)

type JwtRepository struct {
	jwt appjwt.JWTInterface
}

func NewJwtRepository(jwt appjwt.JWTInterface) JwtRepository {
	return JwtRepository{
		jwt: jwt,
	}
}

func (r JwtRepository) GenerateAccessToken(ctx *appcontext.AppContext, userID string) (string, error) {
	return r.jwt.GenerateAccessToken(ctx, userID)
}

func (r JwtRepository) GenerateTokens(ctx *appcontext.AppContext, userID string) (*domain.Tokens, error) {
	result, err := r.jwt.GenerateTokens(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Tokens{
		AccessToken:        result.AccessToken,
		RefreshToken:       result.RefreshToken,
		AccessTokenExpiry:  result.AccessTokenExpiry,
		RefreshTokenExpiry: result.RefreshTokenExpiry,
	}, nil
}
