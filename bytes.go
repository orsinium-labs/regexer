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

			nSubs := len(spans)/2 - 1
			subs := make([]BytesSubMatch, 0, nSubs)
			for i := 2; i < len(spans); i += 2 {
				subStart := spans[i]
				subEnd := spans[i+1]
				sub := BytesSubMatch{
					Content: []byte{},
					Abs: Span{
						Start: shift + subStart,
						End:   shift + subEnd,
					},
					Rel: Span{
						Start: subStart,
						End:   subEnd,
					},
				}
				subs = append(subs, sub)
			}

			spanStart := spans[0]
			spanEnd := spans[1]
			match := BytesMatch{
				Content: subSrc[spanStart:spanEnd],
				Span: Span{
					Start: shift + spanStart,
					End:   shift + spanEnd,
				},
				Sub: subs,
			}
			more := yield(match)
			if !more {
				return
			}
			shift += spanStart
		}
	}
}
