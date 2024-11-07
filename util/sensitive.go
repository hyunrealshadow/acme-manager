package util

import (
	"regexp"
	"strings"
)

func MaskDSN(dsn string) string {
	re := regexp.MustCompile(`(postgres://)([^:]+):([^@]+)@`)
	return re.ReplaceAllString(dsn, `$1******:******@`)
}

func MakeSensitive(info string) string {
	if len(info) <= 4 {
		return info[:1] + strings.Repeat("*", len(info)-1)
	} else {
		return info[:2] + strings.Repeat("*", len(info)-4) + info[len(info)-2:]
	}
}
