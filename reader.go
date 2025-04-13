package regexer

import (
	"io"
	"iter"
	"regexp"
	"unicode/utf8"
)

type Runes struct {
	rex *regexp.Regexp
	src []rune
}

type runeReader struct {
	inner    []rune
	consumed int
}

func (rr *runeReader) ReadRune() (rune, int, error) {
	if rr.consumed >= len(rr.inner) {
		return 0, 0, io.EOF
	}
	r := rr.inner[rr.consumed]
	rr.consumed += 1
	return r, utf8.RuneLen(r), nil
}

func (b Runes) Find() iter.Seq[RMatch] {
	return func(yield func(RMatch) bool) {
		shift := 0
		for {
			subSrc := b.src[shift:]
			reader := runeReader{inner: subSrc}
			spans := b.rex.FindReaderSubmatchIndex(&reader)
			if spans == nil {
				return
			}
			spanStart := spans[0]
			spanEnd := spans[1]
			matchSrc := subSrc[spanStart:spanEnd]
			match := RMatch{
				Content: matchSrc,
				Span: Span{
					Start: shift + spanStart,
					End:   shift + spanEnd,
				},
				Subs: RSubs{
					shift:    shift,
					content:  matchSrc,
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

// func (b Runes) Replace(res *io.RuneReader) iter.Seq[RReplacement] {
// 	return func(yield func(RReplacement) bool) {
// 		prevEnd := 0
// 		for match := range b.Find() {
// 			*res = append(*res, b.src[prevEnd:match.Span.Start]...)
// 			ok := yield(RReplacement{
// 				Match:  match,
// 				rex:    b.rex,
// 				src:    b.src[prevEnd:],
// 				result: res,
// 			})
// 			if !ok {
// 				return
// 			}
// 			prevEnd = match.Span.End
// 		}
// 		*res = append(*res, b.src[prevEnd:]...)
// 	}
// }

func (b Runes) Contains() bool {
	reader := runeReader{inner: b.src}
	return b.rex.MatchReader(&reader)
}

// type RReplacement struct {
// 	RMatch
// 	rex    *regexp.Regexp
// 	src    io.RuneReader
// 	result *io.RuneReader
// }

// func (r RReplacement) ReplaceLiteral(val io.RuneReader) {
// 	*r.result = append(*r.result, val...)
// }

// func (r RReplacement) ReplaceTemplate(val io.RuneReader) {
// 	*r.result = r.rex.Expand(*r.result, val, r.src, r.Subs.rawSpans)
// }

// func (r RReplacement) ReplaceFunc(f func(io.RuneReader) io.RuneReader) {
// 	*r.result = append(*r.result, f(r.Match.Content)...)
// }
