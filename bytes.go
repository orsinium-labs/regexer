package regexer

import (
	"iter"
	"regexp"
)

type (
	BytesMatch    = Match[[]byte]
	BytesSubMatch = SubMatch[[]byte]
)

type Bytes struct {
	rex *regexp.Regexp
	src []byte
}

func (b Bytes) Find() iter.Seq[BytesMatch] {
	return func(yield func(BytesMatch) bool) {
		shift := 0
		for {
			subSrc := b.src[shift:]
			spans := b.rex.FindSubmatchIndex(subSrc)
			if spans == nil {
				return
			}
			spanStart := spans[0]
			spanEnd := spans[1]
			match := BytesMatch{
				Content: subSrc[spanStart:spanEnd],
				Span: Span{
					Start: shift + spanStart,
					End:   shift + spanEnd,
				},
				rawSpans: spans,
			}
			more := yield(match)
			if !more {
				return
			}
			shift += spanStart
		}
	}
}
