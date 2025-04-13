package regexer

import "iter"

// The same as [Match] but for runes.
//
// Because the compiler can't infer the core type of [text] if we extend it
// with a slice of runes.
type RMatch struct {
	// The full match text.
	Content []rune
	// The range of the match in the original text.
	Span Span
	// Matches for sub-patterns.
	Subs RSubs
}

// Matches for sub-patterns.
type RSubs struct {
	content  []rune
	shift    int
	rawSpans []int
}

func (s RSubs) Len() int {
	return len(s.rawSpans)/2 - 1
}

func (s RSubs) At(i int) RSub {
	start := s.rawSpans[i*2]
	end := s.rawSpans[i*2+1]
	return RSub{
		Content: s.content[start:end],
		Span: Span{
			Start: s.shift + start,
			End:   s.shift + end,
		},
	}
}

func (s RSubs) Slice() []RSub {
	spans := s.rawSpans
	nSubs := len(spans)/2 - 1
	subs := make([]RSub, 0, nSubs)
	for i := 2; i < len(spans); i += 2 {
		subStart := spans[i]
		subEnd := spans[i+1]
		sub := RSub{
			Content: s.content[subStart:subEnd],
			Span: Span{
				Start: s.shift + subStart,
				End:   s.shift + subEnd,
			},
		}
		subs = append(subs, sub)
	}
	return subs
}

func (s RSubs) Iter() iter.Seq[RSub] {
	return func(yield func(RSub) bool) {
		spans := s.rawSpans
		for i := 2; i < len(spans); i += 2 {
			subStart := spans[i]
			subEnd := spans[i+1]
			sub := RSub{
				Content: s.content[subStart:subEnd],
				Span: Span{
					Start: s.shift + subStart,
					End:   s.shift + subEnd,
				},
			}
			more := yield(sub)
			if !more {
				return
			}
		}
	}
}

type RSub = Sub[[]rune]
