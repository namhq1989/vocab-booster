package grpc_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/internal/genproto/authpb"
	"github.com/namhq1989/vocab-booster-server-admin/internal/genproto/staffpb"
	mockgrpc "github.com/namhq1989/vocab-booster-server-admin/internal/mock/grpc"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/grpc"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type isAdminTestSuite struct {
	suite.Suite
	handler             grpc.IsAdminHandler
	mockCtrl            *gomock.Controller
	mockGRPCStaffClient *mockgrpc.MockStaffServiceClient
}

func (s *isAdminTestSuite) SetupSuite() {
	s.setupApplication()
}

func (*isAdminTestSuite) AfterTest(_, _ string) {
	// do nothing
}

func (s *isAdminTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockGRPCStaffClient = mockgrpc.NewMockStaffServiceClient(s.mockCtrl)

	staffHub := infrastructure.NewStaffHub(s.mockGRPCStaffClient)
	s.handler = grpc.NewIsAdminHandler(staffHub)
}

func (s *isAdminTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *isAdminTestSuite) Test_1_Success() {
	// mock data
	id := database.NewStringID()
	s.mockGRPCStaffClient.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&staffpb.FindStaffByIDResponse{Staff: &staffpb.Staff{
			Id:      id,
			Name:    "Test",
			Email:   "test@gmail.com",
			IsAdmin: true,
		}}, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.IsAdmin(ctx, &authpb.IsAdminRequest{
		Id: id,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), true, resp.GetIsAdmin())
}

func (s *isAdminTestSuite) Test_2_Fail_NotAdmin() {
	// mock data
	id := database.NewStringID()
	s.mockGRPCStaffClient.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&staffpb.FindStaffByIDResponse{Staff: &staffpb.Staff{
			Id:      id,
			Name:    "Test",
			Email:   "test@gmail.com",
			IsAdmin: false,
		}}, nil)

	// call
	resp, err := s.handler.IsAdmin(appcontext.NewGRPC(context.Background()), &authpb.IsAdminRequest{
		Id: id,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.NotEqual(s.T(), true, resp.GetIsAdmin())
}

func (s *isAdminTestSuite) Test_2_Fail_InvalidUserId() {
	// mock data
	id := "invalid id"

	s.mockGRPCStaffClient.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&staffpb.FindStaffByIDResponse{Staff: nil}, nil)

	resp, err := s.handler.IsAdmin(appcontext.NewGRPC(context.Background()), &authpb.IsAdminRequest{
		Id: id,
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Staff.StaffNotFound, err)
}

//
// END OF CASES
//

func TestIsAdminTestSuite(t *testing.T) {
	suite.Run(t, new(isAdminTestSuite))
}
