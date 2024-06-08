package command_test

import (
	"context"
	"testing"
	"time"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	mockauth "github.com/namhq1989/vocab-booster-server-admin/internal/mock/domain/auth"
	mockgrpc "github.com/namhq1989/vocab-booster-server-admin/internal/mock/grpc"
	mockjwt "github.com/namhq1989/vocab-booster-server-admin/internal/mock/jwt"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/application/command"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/dto"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type refreshAccessTokenTestSuite struct {
	suite.Suite
	handler                      command.RefreshAccessTokenHandler
	mockCtrl                     *gomock.Controller
	mockJwt                      *mockjwt.MockJWTInterface
	mockGRPCStaffClient          *mockgrpc.MockStaffServiceClient
	mockStaffAuthTokenRepository *mockauth.MockStaffAuthTokenRepository
}

func (s *refreshAccessTokenTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *refreshAccessTokenTestSuite) AfterTest(_, _ string) {
	// do nothing
}

func (s *refreshAccessTokenTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockStaffAuthTokenRepository = mockauth.NewMockStaffAuthTokenRepository(s.mockCtrl)
	s.mockJwt = mockjwt.NewMockJWTInterface(s.mockCtrl)
	s.mockGRPCStaffClient = mockgrpc.NewMockStaffServiceClient(s.mockCtrl)

	staffHub := infrastructure.NewStaffHub(s.mockGRPCStaffClient)
	jwtRepository := infrastructure.NewJwtRepository(s.mockJwt)
	s.handler = command.NewRefreshAccessTokenHandler(s.mockStaffAuthTokenRepository, jwtRepository, staffHub)
}

func (s *refreshAccessTokenTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *refreshAccessTokenTestSuite) Test_1_Success() {
	// mock data
	s.mockStaffAuthTokenRepository.EXPECT().
		FindAuthToken(gomock.Any(), gomock.Any()).
		Return(&domain.RefreshToken{
			ID:      database.NewStringID(),
			StaffID: database.NewStringID(),
			Token:   "refresh_token",
			Expiry:  time.Now().Add(time.Second * 1000),
		}, nil)

	s.mockJwt.EXPECT().
		GenerateAccessToken(gomock.Any(), gomock.Any()).
		Return("access_token", nil)

	// call
	ctx := appcontext.New(context.Background())
	resp, err := s.handler.RefreshAccessToken(ctx, dto.RefreshAccessTokenRequest{
		RefreshToken: "refresh_token",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "access_token", resp.AccessToken)
}

func (s *refreshAccessTokenTestSuite) Test_2_Fail_TokenExpired() {
	// mock data
	s.mockStaffAuthTokenRepository.EXPECT().
		FindAuthToken(gomock.Any(), gomock.Any()).
		Return(&domain.RefreshToken{
			ID:      database.NewStringID(),
			StaffID: database.NewStringID(),
			Token:   "refresh_token",
			Expiry:  time.Now().Add(time.Second * -1000),
		}, nil)

	s.mockStaffAuthTokenRepository.EXPECT().
		DeleteAuthToken(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.New(context.Background())
	resp, err := s.handler.RefreshAccessToken(ctx, dto.RefreshAccessTokenRequest{
		RefreshToken: "refresh_token",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Auth.InvalidAuthToken, err)
}

func (s *refreshAccessTokenTestSuite) Test_2_Fail_TokenNotFound() {
	// mock data
	s.mockStaffAuthTokenRepository.EXPECT().
		FindAuthToken(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	// call
	ctx := appcontext.New(context.Background())
	resp, err := s.handler.RefreshAccessToken(ctx, dto.RefreshAccessTokenRequest{
		RefreshToken: "refresh_token",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Auth.InvalidAuthToken, err)
}

//
// END OF CASES
//

func TestRefreshAccessTokenTestSuite(t *testing.T) {
	suite.Run(t, new(refreshAccessTokenTestSuite))
}
