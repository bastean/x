package release_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/pkg/release"
)

type CommitTestSuite struct {
	suite.Default
	SUT  *release.Commit
	doer *release.DoerMock
}

func (s *CommitTestSuite) SetupSuite() {
	s.doer = new(release.DoerMock)

	s.SUT = &release.Commit{
		Doer: s.doer,
	}
}

func (s *CommitTestSuite) TestCreate() {
	message := release.Mother().Message()

	cmds := []string{"git", "commit", "--allow-empty", "-m", message}

	s.doer.Mock.On("Do", cmds).Return("")

	s.NoError(s.SUT.Create(message))

	s.doer.Mock.AssertExpectations(s.T())
}

func (s *CommitTestSuite) TestCreateStd() {
	module := release.Mother().RandomModuleRelease()

	latest, version, _, _, _ := release.Mother().RandomTag(module)

	message := "chore(release): " + latest

	cmds := []string{"git", "commit", "--allow-empty", "-m", message}

	s.doer.Mock.On("Do", cmds).Return("")

	s.NoError(s.SUT.CreateStd(module, version))

	s.doer.Mock.AssertExpectations(s.T())
}

func (s *CommitTestSuite) TestReset() {
	cmds := []string{"git", "reset", "--hard", "HEAD^"}

	s.doer.Mock.On("Do", cmds).Return("")

	s.NoError(s.SUT.Reset())

	s.doer.Mock.AssertExpectations(s.T())
}

func TestUnitCommitSuite(t *testing.T) {
	suite.Run(t, new(CommitTestSuite))
}
