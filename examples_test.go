package regexer_test

import (
	"fmt"

	"github.com/orsinium-labs/regexer"
)

func ExampleString_Find() {
	rex := regexer.New(`[a-z]+`)
	input := "never gonna give you up"
	matches := rex.String(input).Find()
	for match := range matches {
		_ = match
		fmt.Println(match.Span.Start, match.Content)
	}
	//Output:
	// 0 never
	// 6 gonna
	// 12 give
	// 17 you
	// 21 up
}

func ExampleString_Contains() {
	rex := regexer.New(`[0-9]+`)
	input := "number 42 is the answer"
	contains := rex.String(input).Contains()
	if contains {
		fmt.Println("the string contains a regexp match")
	}
	//Output: the string contains a regexp match
}

func ExampleBytes_Find() {
	rex := regexer.New(`[a-z]+`)
	input := []byte("never gonna give you up")
	matches := rex.Bytes(input).Find()
	for match := range matches {
		_ = match
		fmt.Println(match.Span.Start, string(match.Content))
	}
	//Output:
	// 0 never
	// 6 gonna
	// 12 give
	// 17 you
	// 21 up
}

func ExampleBytes_Contains() {
	rex := regexer.New(`[0-9]+`)
	input := []byte("number 42 is the answer")
	contains := rex.Bytes(input).Contains()
	if contains {
		fmt.Println("the byte slice contains a regexp match")
	}
	//Output: the byte slice contains a regexp match
}
