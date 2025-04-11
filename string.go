package regexer

import (
	"iter"
	"regexp"
)

type (
	SMatch = Match[string]
	SSub   = Sub[string]
	SSubs  = Subs[string]
)

type String struct {
	rex *regexp.Regexp
	src string
}

func (b String) Find() iter.Seq[SMatch] {
	return func(yield func(SMatch) bool) {
		shift := 0
		for {
			subSrc := b.src[shift:]
			spans := b.rex.FindStringSubmatchIndex(subSrc)
			if spans == nil {
				return
			}
			spanStart := spans[0]
			spanEnd := spans[1]
			matchSrc := subSrc[spanStart:spanEnd]
			match := SMatch{
				Content: matchSrc,
				Span: Span{
					Start: shift + spanStart,
					End:   shift + spanEnd,
				},
				Subs: SSubs{
					shift:    shift,
					content:  matchSrc,
					rawSpans: spans,
				},
			}
			more := yield(match)
			if !more {
				return
			}
			shift += spanStart
		}
	}
}

func (b String) Match() bool {
	return b.rex.MatchString(b.src)
}
