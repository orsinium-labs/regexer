package regexer_test

import (
	"bytes"
	"fmt"
	"strings"

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

func ExampleString_Replace() {
	rex := regexer.New(`(is|the)`)
	input := "number 42 is the answer"
	var result string
	matches := rex.String(input).Replace(&result)
	for match := range matches {
		newVal := strings.ToUpper(match.Content)
		match.ReplaceLiteral(newVal)
	}
	fmt.Println(string(result))
	//Output: number 42 IS THE answer
}

func ExampleSReplacement_ReplaceLiteral() {
	rex := regexer.New(`(is|the)`)
	input := "number 42 is the answer"
	var result string
	matches := rex.String(input).Replace(&result)
	for match := range matches {
		newVal := strings.ToUpper(match.Content)
		match.ReplaceLiteral(newVal)
	}
	fmt.Println(string(result))
	//Output: number 42 IS THE answer
}

func ExampleSReplacement_ReplaceTemplate() {
	rex := regexer.New(`(is|the)`)
	input := "number 42 is the answer"
	var result string
	matches := rex.String(input).Replace(&result)
	for match := range matches {
		template := string(`[$1]`)
		match.ReplaceTemplate(template)
	}
	fmt.Println(string(result))
	//Output: number 42 [is] [the] answer
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

func ExampleBytes_Replace() {
	rex := regexer.New(`(is|the)`)
	input := []byte("number 42 is the answer")
	var result []byte
	matches := rex.Bytes(input).Replace(&result)
	for match := range matches {
		newVal := bytes.ToUpper(match.Content)
		match.ReplaceLiteral(newVal)
	}
	fmt.Println(string(result))
	//Output: number 42 IS THE answer
}

func ExampleBReplacement_ReplaceLiteral() {
	rex := regexer.New(`(is|the)`)
	input := []byte("number 42 is the answer")
	var result []byte
	matches := rex.Bytes(input).Replace(&result)
	for match := range matches {
		newVal := bytes.ToUpper(match.Content)
		match.ReplaceLiteral(newVal)
	}
	fmt.Println(string(result))
	//Output: number 42 IS THE answer
}

func ExampleBReplacement_ReplaceTemplate() {
	rex := regexer.New(`(is|the)`)
	input := []byte("number 42 is the answer")
	var result []byte
	matches := rex.Bytes(input).Replace(&result)
	for match := range matches {
		template := []byte(`[$1]`)
		match.ReplaceTemplate(template)
	}
	fmt.Println(string(result))
	//Output: number 42 [is] [the] answer
}
