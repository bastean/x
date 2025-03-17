package release_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"

	"github.com/bastean/x/tools/pkg/release"
)

type TagTestSuite struct {
	suite.Suite
	SUT  *release.Tag
	doer *release.DoerMock
}

func (s *TagTestSuite) SetupSuite() {
	s.doer = new(release.DoerMock)

	s.SUT = &release.Tag{
		Doer: s.doer,
	}
}

func (s *TagTestSuite) TestLatest() {
	module := release.RandomModuleRelease()

	expected, _, _, _, _ := release.RandomTag(module)

	cmds := []string{"bash", "-c", fmt.Sprintf("git tag --sort -v:refname | grep %s | head -n 1", module.Name)}

	s.doer.Mock.On("Do", cmds).Return(expected)

	actual, err := s.SUT.Latest(module)

	s.NoError(err)

	s.doer.Mock.AssertExpectations(s.T())

	s.Equal(expected, actual)
}

func (s *TagTestSuite) TestCreate() {
	module := release.RandomModuleRelease()

	_, version, _, _, _ := release.RandomTag(module)

	annotate := "v" + version

	message := services.Create.Message()

	cmds := []string{"git", "tag", "-a", annotate, "-m", message}

	s.doer.Mock.On("Do", cmds).Return("")

	s.NoError(s.SUT.Create(annotate, message))

	s.doer.Mock.AssertExpectations(s.T())
}

func (s *TagTestSuite) TestCreateStd() {
	module := release.RandomModuleRelease()

	_, version, _, _, _ := release.RandomTag(module)

	annotate := fmt.Sprintf("%s/v%s", module.Name, version)

	message := fmt.Sprintf("%s %s", module.Name, version)

	cmds := []string{"git", "tag", "-a", annotate, "-m", message}

	s.doer.Mock.On("Do", cmds).Return("")

	s.NoError(s.SUT.CreateStd(module, version))

	s.doer.Mock.AssertExpectations(s.T())
}

func TestUnitTagSuite(t *testing.T) {
	suite.Run(t, new(TagTestSuite))
}
