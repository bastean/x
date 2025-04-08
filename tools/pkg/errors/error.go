package errors

import (
	"regexp"
	"strings"
)

const (
	RExEmbed = `\[.*]`
)

func Extract(err error) string {
	return strings.Trim(regexp.MustCompile(RExEmbed).FindString(err.Error()), "[]")
}
