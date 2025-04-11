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
	Content T
	Span    Span
	Subs    Subs[T]
}

type Subs[T Text] struct {
	content  T
	shift    int
	rawSpans []int
}

func (s Subs[T]) At(i int) Sub[T] {
	start := s.rawSpans[i*2]
	end := s.rawSpans[i*2+1]
	return Sub[T]{
		Content: s.content[start:end],
		Span: Span{
			Start: s.shift + start,
			End:   s.shift + end,
		},
	}
}

func (s Subs[T]) Slice() []Sub[T] {
	spans := s.rawSpans
	nSubs := len(spans)/2 - 1
	subs := make([]Sub[T], 0, nSubs)
	for i := 2; i < len(spans); i += 2 {
		subStart := spans[i]
		subEnd := spans[i+1]
		sub := Sub[T]{
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

type Sub[T Text] struct {
	Content T
	Span    Span
}
