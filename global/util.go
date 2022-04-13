package global

import "strings"

type Utils struct {
}

func (Utils) Eq(a, b string) bool {
	if strings.Compare(a, b) == 0 {
		return true
	}
	return false
}
