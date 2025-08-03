package release_test

import (
	"os/exec"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/internal/app/release"
)

type ReleaseTestSuite struct {
	suite.Default
	SUT string
}

func (s *ReleaseTestSuite) SetupSuite() {
	s.SUT = "../../../cmd/release"
}

func (s *ReleaseTestSuite) TestSentinel() {
	s.Equal("Release", release.App)
}

func (s *ReleaseTestSuite) TestHelp() {
	expected := `________       ______
___  __ \_____ ___  /_____ ______ ______________
__  /_/ /_  _ \__  / _  _ \_  __ ` + "`" + `/__  ___/_  _ \
_  _, _/ /  __/_  /  /  __// /_/ / _(__  ) /  __/
/_/ |_|  \___/ /_/   \___/ \__,_/  /____/  \___/

Usage: release [flags]

  -f	First Release (default: false)
  -i string
    	Increment "patch", "minor" or "major" (optional: if "-f" is used)
  -m string
    	Module name (required)
`

	actual, err := exec.Command("go", "run", s.SUT, "-h").CombinedOutput() //nolint:gosec

	s.NoError(err)

	s.Equal(expected, string(actual))
}

func TestAcceptanceReleaseSuite(t *testing.T) {
	suite.Run(t, new(ReleaseTestSuite))
}
