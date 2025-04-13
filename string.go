package regexer

import (
	"iter"
	"regexp"
	"strings"
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
			shift += spanEnd
		}
	}
}

func (s String) Replace(res *string) iter.Seq[SReplacement] {
	return func(yield func(SReplacement) bool) {
		prevEnd := 0
		resBuilder := strings.Builder{}
		for match := range s.Find() {
			resBuilder.WriteString(s.src[prevEnd:match.Span.Start])
			ok := yield(SReplacement{
				Match:  match,
				rex:    s.rex,
				src:    s.src[prevEnd:],
				result: &resBuilder,
			})
			if !ok {
				return
			}
			prevEnd = match.Span.End
		}
		resBuilder.WriteString(s.src[prevEnd:])
		*res = resBuilder.String()
	}
}

func (b String) Contains() bool {
	return b.rex.MatchString(b.src)
}

type SReplacement struct {
	Match[string]
	rex    *regexp.Regexp
	src    string
	result *strings.Builder
}

func (r SReplacement) ReplaceLiteral(val string) {
	r.result.WriteString(val)
}

func (r SReplacement) ReplaceTemplate(val string) {
	// TODO: avoid allocations on bytes<->string conversion by using unsafe.
	suffix := r.rex.Expand(nil, []byte(val), []byte(r.src), r.Subs.rawSpans)
	r.result.WriteString(string(suffix))
}

func (r SReplacement) ReplaceFunc(f func(string) string) {
	r.result.WriteString(f(r.Match.Content))
}
