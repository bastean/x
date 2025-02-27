package example_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/stretchr/testify/suite"

	"github.com/bastean/x/example"
)

type PrintTestSuite struct {
	suite.Suite
}

func (s *PrintTestSuite) TestSprint() {
	expected := services.Create.Message()

	actual := example.Sprint(expected)

	s.Equal(expected, actual)
}

func TestUnitPrintSuite(t *testing.T) {
	suite.Run(t, new(PrintTestSuite))
}
