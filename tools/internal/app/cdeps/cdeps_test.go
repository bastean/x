package cdeps_test

import (
	"os/exec"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type CDepsTestSuite struct {
	suite.Default
	SUT string
}

func (s *CDepsTestSuite) SetupSuite() {
	s.SUT = "../../../cmd/cdeps"
}

func (s *CDepsTestSuite) TestHelp() {
	expected := `Usage: cdeps [flags]

  -c string
    	cDeps configuration file (required) (default "cdeps.json")
`

	actual, err := exec.Command("go", "run", s.SUT, "-h").CombinedOutput() //nolint:gosec

	s.NoError(err)

	s.EqualValues(expected, actual)
}

func TestAcceptanceCDepsSuite(t *testing.T) {
	suite.Run(t, new(CDepsTestSuite))
}
