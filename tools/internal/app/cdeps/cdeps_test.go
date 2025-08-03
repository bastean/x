package cdeps_test

import (
	"os/exec"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/internal/app/cdeps"
)

type CDepsTestSuite struct {
	suite.Default
	SUT string
}

func (s *CDepsTestSuite) SetupSuite() {
	s.SUT = "../../../cmd/cdeps"
}

func (s *CDepsTestSuite) TestSentinel() {
	s.Equal("cDeps", cdeps.App)
}

func (s *CDepsTestSuite) TestHelp() {
	expected := `       ________
__________  __ \_____ ________ ________
_  ___/__  / / /_  _ \___  __ \__  ___/
/ /__  _  /_/ / /  __/__  /_/ /_(__  )
\___/  /_____/  \___/ _  .___/ /____/
                      /_/

Usage: cdeps [flags]

  -c string
    	cDeps configuration file (required) (default "cdeps.json")
`

	actual, err := exec.Command("go", "run", s.SUT, "-h").CombinedOutput() //nolint:gosec

	s.NoError(err)

	s.Equal(expected, string(actual))
}

func TestAcceptanceCDepsSuite(t *testing.T) {
	suite.Run(t, new(CDepsTestSuite))
}
