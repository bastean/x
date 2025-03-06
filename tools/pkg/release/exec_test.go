package release_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/x/tools/pkg/release"
)

type ExecTestSuite struct {
	suite.Suite
	SUT  release.Doer
	file string
}

func (s *ExecTestSuite) SetupSuite() {
	s.file = "ignore.test"
	s.SUT = new(release.Exec)
}

func (s *ExecTestSuite) SetupTest() {
	_ = os.Remove(s.file)
}

func (s *ExecTestSuite) TestDo() {
	cmds := []string{"touch", s.file}

	_, err := s.SUT.Do(cmds...)

	s.NoError(err)

	s.FileExists(s.file)
}

func (s *ExecTestSuite) TearDownTest() {
	_ = os.Remove(s.file)
}

func TestIntegrationExecSuite(t *testing.T) {
	suite.Run(t, new(ExecTestSuite))
}
