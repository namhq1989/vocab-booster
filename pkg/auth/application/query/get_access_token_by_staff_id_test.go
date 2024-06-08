package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	mockjwt "github.com/namhq1989/vocab-booster-server-admin/internal/mock/jwt"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/application/query"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/dto"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getAccessTokenByStaffIDTestSuite struct {
	suite.Suite
	handler  query.GetAccessTokenByStaffIDHandler
	mockCtrl *gomock.Controller
	mockJwt  *mockjwt.MockJWTInterface
}

func (s *getAccessTokenByStaffIDTestSuite) SetupSuite() {
	s.setupApplication()
}

func (*getAccessTokenByStaffIDTestSuite) AfterTest(_, _ string) {
	// do nothing
}

func (s *getAccessTokenByStaffIDTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockJwt = mockjwt.NewMockJWTInterface(s.mockCtrl)

	jwtRepository := infrastructure.NewJwtRepository(s.mockJwt)
	s.handler = query.NewGetAccessTokenByStaffIDHandler(jwtRepository)
}

func (s *getAccessTokenByStaffIDTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getAccessTokenByStaffIDTestSuite) Test_1_Success() {
	// mock data
	s.mockJwt.EXPECT().
		GenerateAccessToken(gomock.Any(), gomock.Any()).
		Return("access_token", nil)

	// call
	ctx := appcontext.New(context.Background())
	resp, err := s.handler.GetAccessTokenByStaffID(ctx, dto.GetAccessTokenByStaffIDRequest{StaffID: database.NewStringID()})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "access_token", resp.AccessToken)
}

func (s *getAccessTokenByStaffIDTestSuite) Test_2_Fail_InvalidID() {
	// call
	ctx := appcontext.New(context.Background())
	resp, err := s.handler.GetAccessTokenByStaffID(ctx, dto.GetAccessTokenByStaffIDRequest{StaffID: "invalid id"})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Staff.InvalidStaffID, err)
}

//
// END OF CASES
//

func TestGetAccessTokenByStaffIDTestSuite(t *testing.T) {
	suite.Run(t, new(getAccessTokenByStaffIDTestSuite))
}
