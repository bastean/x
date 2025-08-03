package log_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/internal/pkg/log"
)

type LogTestSuite struct {
	suite.Default
}

func (s *LogTestSuite) TestSentinel() {
	s.Equal("speed", log.Font)
}

func TestUnitLogSuite(t *testing.T) {
	suite.Run(t, new(LogTestSuite))
}
