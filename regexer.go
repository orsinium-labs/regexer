package regexer

import "regexp"

type Regex struct {
	rex *regexp.Regexp
}

type stringLiteral string

func New(raw stringLiteral) Regex {
	return Regex{
		rex: regexp.MustCompile(string(raw)),
	}
}

func (r Regex) Bytes(src []byte) Bytes {
	return Bytes{rex: r.rex, src: src}
}

func (r Regex) String(src string) String {
	return String{rex: r.rex, src: src}
}

func (r Regex) Runes(src []rune) Runes {
	return Runes{rex: r.rex, src: src}
}
