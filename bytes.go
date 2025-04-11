package regexer

import (
	"iter"
	"regexp"
)

type (
	BMatch = Match[[]byte]
	BSub   = Sub[[]byte]
	BSubs  = Subs[[]byte]
)

type Bytes struct {
	rex *regexp.Regexp
	src []byte
}

func (b Bytes) Find() iter.Seq[BMatch] {
	return func(yield func(BMatch) bool) {
		shift := 0
		for {
			subSrc := b.src[shift:]
			spans := b.rex.FindSubmatchIndex(subSrc)
			if spans == nil {
				return
			}
			spanStart := spans[0]
			spanEnd := spans[1]
			matchSrc := subSrc[spanStart:spanEnd]
			match := BMatch{
				Content: matchSrc,
				Span: Span{
					Start: shift + spanStart,
					End:   shift + spanEnd,
				},
				Subs: BSubs{
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
