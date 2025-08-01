package cdeps

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type m struct {
	*mother.Mother
}

func (m *m) WildcardNew(wildcard string) string {
	return fmt.Sprintf("{%s}", wildcard)
}

func (m *m) WildcardValid() string {
	return m.WildcardNew(m.Word())
}

func (m *m) WildcardInvalid() string {
	return m.Word()
}

func (m *m) WildcardsValid() map[string]string {
	return map[string]string{
		m.Word(): m.Word(),
		m.Word(): m.Word(),
		m.Word(): m.Word(),
	}
}

func (m *m) WildcardsInvalid() map[string]string {
	return map[string]string{}
}

var Mother = mother.New[m]
