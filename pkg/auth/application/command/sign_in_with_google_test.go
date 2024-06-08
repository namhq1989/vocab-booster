package command_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type signInWithGoogleTestSuite struct {
	suite.Suite
}

func (s *signInWithGoogleTestSuite) Test_1_Success() {
	s.T().Skip("not implemented yet")
}

func (s *signInWithGoogleTestSuite) Test_2_Fail() {
	s.T().Skip("not implemented yet")
}

func TestSignInWithGoogleTestSuite(t *testing.T) {
	suite.Run(t, new(signInWithGoogleTestSuite))
}
