package regexer

import "regexp"

type String struct {
	rex *regexp.Regexp
	src string
}
