package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/internal/genproto/staffpb"
	mockgrpc "github.com/namhq1989/vocab-booster-server-admin/internal/mock/grpc"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/application/query"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/dto"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getMeTestSuite struct {
	suite.Suite
	handler             query.GetMeHandler
	mockCtrl            *gomock.Controller
	mockGRPCStaffClient *mockgrpc.MockStaffServiceClient
}

func (s *getMeTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getMeTestSuite) AfterTest(_, _ string) {
	// do nothing
}

func (s *getMeTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockGRPCStaffClient = mockgrpc.NewMockStaffServiceClient(s.mockCtrl)

	staffHub := infrastructure.NewStaffHub(s.mockGRPCStaffClient)
	s.handler = query.NewGetMeHandler(staffHub)
}

func (s *getMeTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getMeTestSuite) Test_1_Success() {
	// mock data
	id := database.NewStringID()

	s.mockGRPCStaffClient.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&staffpb.FindStaffByIDResponse{
			Staff: &staffpb.Staff{
				Id:      id,
				Name:    "Test",
				Email:   "test@gmail.com",
				IsAdmin: true,
			},
		}, nil)

	// call
	ctx := appcontext.New(context.Background())
	resp, err := s.handler.GetMe(ctx, id, dto.GetMeRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), id, resp.ID)
}

func (s *getMeTestSuite) Test_2_Fail_InvalidID() {
	// call
	id := "invalid id"

	s.mockGRPCStaffClient.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	ctx := appcontext.New(context.Background())
	resp, err := s.handler.GetMe(ctx, id, dto.GetMeRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Staff.StaffNotFound, err)
}

//
// END OF CASES
//

func TestGetMeTestSuite(t *testing.T) {
	suite.Run(t, new(getMeTestSuite))
}
