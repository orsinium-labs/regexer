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
	Content []byte
	Span    Span
	Sub     []SubMatch[T]
}

type SubMatch[T Text] struct {
	Content []byte
	Abs     Span
	Rel     Span
}
