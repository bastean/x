package errs_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/internal/pkg/errs"
)

type SentinelTestSuite struct {
	suite.Default
}

func (s *SentinelTestSuite) TestSentinel() {
	s.EqualError(errs.ErrRequiredFlags, "please define required flags")
}

func TestUnitSentinelSuite(t *testing.T) {
	suite.Run(t, new(SentinelTestSuite))
}
