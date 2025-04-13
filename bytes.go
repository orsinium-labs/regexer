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
			match := BMatch{
				Content: subSrc[spanStart:spanEnd],
				Span: Span{
					Start: shift + spanStart,
					End:   shift + spanEnd,
				},
				Subs: BSubs{
					shift:    shift,
					content:  subSrc,
					rawSpans: spans,
				},
			}
			more := yield(match)
			if !more {
				return
			}
			shift += spanEnd
		}
	}
}

func (b Bytes) Replace(res *[]byte) iter.Seq[BReplacement] {
	return func(yield func(BReplacement) bool) {
		prevEnd := 0
		for match := range b.Find() {
			*res = append(*res, b.src[prevEnd:match.Span.Start]...)
			ok := yield(BReplacement{
				Match:  match,
				rex:    b.rex,
				src:    b.src[prevEnd:],
				result: res,
			})
			if !ok {
				return
			}
			prevEnd = match.Span.End
		}
		*res = append(*res, b.src[prevEnd:]...)
	}
}

func (b Bytes) Contains() bool {
	return b.rex.Match(b.src)
}

type BReplacement struct {
	Match[[]byte]
	rex    *regexp.Regexp
	src    []byte
	result *[]byte
}

func (r BReplacement) ReplaceLiteral(val []byte) {
	*r.result = append(*r.result, val...)
}

func (r BReplacement) ReplaceTemplate(val []byte) {
	*r.result = r.rex.Expand(*r.result, val, r.src, r.Subs.rawSpans)
}

func (r BReplacement) ReplaceFunc(f func([]byte) []byte) {
	*r.result = append(*r.result, f(r.Match.Content)...)
}
