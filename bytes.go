package regexer

import (
	"iter"
	"regexp"
)

type BMatch struct {
	Content []byte
	Start   int
	End     int
}

type Bytes struct {
	rex *regexp.Regexp
	src []byte
}

func (b Bytes) Find() iter.Seq[BMatch] {
	return func(yield func(BMatch) bool) {
		shift := 0
		for {
			subSrc := b.src[shift:]
			span := b.rex.FindIndex(subSrc)
			if span == nil {
				return
			}
			spanStart := span[0]
			spanEnd := span[1]
			more := yield(BMatch{
				Content: subSrc[spanStart:spanEnd],
				Start:   shift + spanStart,
				End:     shift + spanEnd,
			})
			if !more {
				return
			}
			shift += spanStart
		}
	}
}
