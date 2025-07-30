package syncenv_test

import (
	"os/exec"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type SyncEnvTestSuite struct {
	suite.Default
	SUT string
}

func (s *SyncEnvTestSuite) SetupSuite() {
	s.SUT = "../../../cmd/syncenv"
}

func (s *SyncEnvTestSuite) TestHelp() {
	expected := `Usage: syncenv [flags]

  -e string
    	Path to ".env" files directory (required)
  -t string
    	Path to ".env" file template (required)
`

	actual, err := exec.Command("go", "run", s.SUT, "-h").CombinedOutput() //nolint:gosec

	s.NoError(err)

	s.EqualValues(expected, actual)
}

func TestAcceptanceSyncEnvSuite(t *testing.T) {
	suite.Run(t, new(SyncEnvTestSuite))
}
