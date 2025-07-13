package release

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mock"
)

type DoerMock struct {
	mock.Default
}

func (m *DoerMock) Do(list ...string) (string, error) {
	args := m.Called(list)
	return args.Get(0).(string), nil
}
