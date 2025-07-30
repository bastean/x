package release_test

import (
	"os/exec"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type ReleaseTestSuite struct {
	suite.Default
	SUT string
}

func (s *ReleaseTestSuite) SetupSuite() {
	s.SUT = "../../../cmd/release"
}

func (s *ReleaseTestSuite) TestHelp() {
	expected := `Usage: release [flags]

  -f	First Release (default: false)
  -i string
    	Increment "patch", "minor" or "major" (optional: if "-f" is used)
  -m string
    	Module name (required)
`

	actual, err := exec.Command("go", "run", s.SUT, "-h").CombinedOutput() //nolint:gosec

	s.NoError(err)

	s.EqualValues(expected, actual)
}

func TestAcceptanceReleaseSuite(t *testing.T) {
	suite.Run(t, new(ReleaseTestSuite))
}
