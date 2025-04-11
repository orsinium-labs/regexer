package regexer

type Text interface {
	string | []byte
}

type Span struct {
	Start int
	End   int
}

func (s Span) Len() int {
	return s.End - s.Start
}

type Match[T Text] struct {
	Content  []byte
	Span     Span
	rawSpans []int
}

func (m Match[T]) Sub(i int) SubMatch[T] {
	shift := m.Span.Start
	start := m.rawSpans[i*2]
	end := m.rawSpans[i*2+1]
	return SubMatch[T]{
		Content: []byte{},
		Span: Span{
			Start: shift + start,
			End:   shift + end,
		},
	}
}

func (m Match[T]) Subs() []SubMatch[T] {
	spans := m.rawSpans
	shift := m.Span.Start
	nSubs := len(spans)/2 - 1
	subs := make([]SubMatch[T], 0, nSubs)
	for i := 2; i < len(spans); i += 2 {
		subStart := spans[i]
		subEnd := spans[i+1]
		sub := SubMatch[T]{
			Content: []byte{},
			Span: Span{
				Start: shift + subStart,
				End:   shift + subEnd,
			},
		}
		subs = append(subs, sub)
	}
	return subs
}

type SubMatch[T Text] struct {
	Content []byte
	Span    Span
}
