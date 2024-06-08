package worker_test

import (
	"testing"

	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/worker"
	"github.com/stretchr/testify/suite"
)

type newStaffCreatedTestSuite struct {
	suite.Suite
	worker worker.Worker
}

// func (s *newUserCreatedTestSuite) SetupSuite() {
// 	s.setupApplication()
// }
//
// func (s *newUserCreatedTestSuite) AfterTest(_, _ string) {
// 	// do nothing
// }
// func (s *newUserCreatedTestSuite) setupApplication() {
// 	auditRepository := infrastructure.NewAuditRepository(s.dbInstance)
// 	s.worker = worker.New(s.queueInstance, auditRepository)
// }

//
// CASES
//

// func (s *newUserCreatedTestSuite) Test_1_Success() {
// 	err := s.worker.NewUserCreated(appcontext.NewWorker(context.Background()), domain.QueueNewUserCreatedAuditLog{
// 		ActorID:   database.NewStringID(),
// 		UserID:    database.NewStringID(),
// 		SourceIp:  "1.1.1.1",
// 		CreatedAt: time.Now(),
// 	})
//
// 	assert.Nil(s.T(), err)
// }
//
// func (s *newUserCreatedTestSuite) Test_2_Fail_InvalidActorID() {
// 	err := s.worker.NewUserCreated(appcontext.NewWorker(context.Background()), domain.QueueNewUserCreatedAuditLog{
// 		ActorID:   "invalid_id",
// 		UserID:    database.NewStringID(),
// 		SourceIp:  "1.1.1.1",
// 		CreatedAt: time.Now(),
// 	})
//
// 	assert.NotNil(s.T(), err)
// 	assert.Equal(s.T(), err, apperrors.Common.InvalidID)
// }

//
// END OF CASES
//

func TestNewStaffCreatedTestSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(newStaffCreatedTestSuite))
}
