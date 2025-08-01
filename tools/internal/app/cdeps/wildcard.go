package cdeps

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	RExWildcard = `{[^{}]+}`
)

var (
	RExWildcardDo = regexp.MustCompile(RExWildcard)
)

func HasWildcard(text string) bool {
	return RExWildcardDo.MatchString(text)
}

func Interpolate(text string, wildcards map[string]string) (string, error) {
	var (
		key, value string
		ok         bool
		err        error
	)

	value = RExWildcardDo.ReplaceAllStringFunc(text, func(wildcard string) string {
		key = strings.Trim(wildcard, "{}")

		if value, ok = wildcards[key]; !ok {
			err = errors.Join(err, fmt.Errorf("%q wildcard unknown in %q", key, text))
		}

		return value
	})

	if err != nil {
		return "", err
	}

	return value, nil
}
