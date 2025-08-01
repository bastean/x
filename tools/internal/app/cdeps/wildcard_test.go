package cdeps_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"

	"github.com/bastean/x/tools/internal/app/cdeps"
)

type WildcardTestSuite struct {
	suite.Default
}

func (s *WildcardTestSuite) TestSentinel() {
	s.Equal(`{[^{}]+}`, cdeps.RExWildcard)
}

func (s *WildcardTestSuite) TestHasWildcardWithValidValue() {
	s.True(cdeps.HasWildcard(cdeps.Mother().WildcardValid()))
}

func (s *WildcardTestSuite) TestHasWildcardWithInvalidValue() {
	s.False(cdeps.HasWildcard(cdeps.Mother().WildcardInvalid()))
}

func (s *WildcardTestSuite) TestInterpolateWithValidValue() {
	wildcards := cdeps.Mother().WildcardsValid()

	key := cdeps.Mother().RandomMapKey(wildcards).(string)

	actual, err := cdeps.Interpolate(cdeps.Mother().WildcardNew(key), wildcards)

	s.NoError(err)

	expected := wildcards[key]

	s.Equal(expected, actual)
}

func (s *WildcardTestSuite) TestInterpolateErrWildcardUnknown() {
	wildcards := cdeps.Mother().WildcardsInvalid()

	wildcard := cdeps.Mother().Word()

	text := fmt.Sprintf("%s/%s", cdeps.Mother().WildcardNew(wildcard), cdeps.Mother().Word())

	_, actual := cdeps.Interpolate(text, wildcards)

	expected := errors.Join(fmt.Errorf("%q wildcard unknown in %q", wildcard, text))

	s.Equal(expected, actual)
}

func TestUnitWildcardSuite(t *testing.T) {
	suite.Run(t, new(WildcardTestSuite))
}
