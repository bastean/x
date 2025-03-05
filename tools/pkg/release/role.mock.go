package release

import (
	"github.com/stretchr/testify/mock"
)

type DoerMock struct {
	mock.Mock
}

func (m *DoerMock) Do(list ...string) (string, error) {
	args := m.Called(list)
	return args.Get(0).(string), nil
}
