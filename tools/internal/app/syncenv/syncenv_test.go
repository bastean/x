package syncenv_test

import (
	"os/exec"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/internal/app/syncenv"
)

type SyncEnvTestSuite struct {
	suite.Default
	SUT string
}

func (s *SyncEnvTestSuite) SetupSuite() {
	s.SUT = "../../../cmd/syncenv"
}

func (s *SyncEnvTestSuite) TestSentinel() {
	s.Equal("syncENV", syncenv.App)
	s.Equal(".env", syncenv.EnvFile)
}

func (s *SyncEnvTestSuite) TestHelp() {
	expected := `                                _______________   _____    __
_____________  _________ __________  ____/___  | / /__ |  / /
__  ___/__  / / /__  __ \_  ___/__  __/   __   |/ / __ | / /
_(__  ) _  /_/ / _  / / // /__  _  /___   _  /|  /  __ |/ /
/____/  _\__, /  /_/ /_/ \___/  /_____/   /_/ |_/   _____/
        /____/

Usage: syncenv [flags]

  -e string
    	Path to ".env" files directory (required)
  -t string
    	Path to ".env" file template (required)
`

	actual, err := exec.Command("go", "run", s.SUT, "-h").CombinedOutput() //nolint:gosec

	s.NoError(err)

	s.Equal(expected, string(actual))
}

func TestAcceptanceSyncEnvSuite(t *testing.T) {
	suite.Run(t, new(SyncEnvTestSuite))
}
