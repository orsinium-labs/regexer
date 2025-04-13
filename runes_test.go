package regexer_test

import (
	"fmt"
	"unicode"

	"github.com/orsinium-labs/regexer"
)

func ExampleRunes_Find() {
	rex := regexer.New(`[a-z]+`)
	input := []rune("never gonna give you up")
	matches := rex.Runes(input).Find()
	for match := range matches {
		fmt.Println(match.Span.Start, string(match.Content))
	}
	//Output:
	// 0 never
	// 6 gonna
	// 12 give
	// 17 you
	// 21 up
}

func ExampleRunes_Contains() {
	rex := regexer.New(`[0-9]+`)
	input := []rune("number 42 is the answer")
	contains := rex.Runes(input).Contains()
	if contains {
		fmt.Println("the rune slice contains a regexp match")
	}
	//Output: the rune slice contains a regexp match
}

func ExampleRunes_Replace() {
	rex := regexer.New(`\w+`)
	input := []rune("number 42 is the answer")
	var result []rune
	matches := rex.Runes(input).Replace(&result)
	for match := range matches {
		first := unicode.ToUpper(match.Content[0])
		newVal := append([]rune{first}, match.Content[1:]...)
		match.ReplaceLiteral(newVal)
	}
	fmt.Println(string(result))
	//Output: Number 42 Is The Answer
}

func ExampleRReplacement_ReplaceLiteral() {
	rex := regexer.New(`\w+`)
	input := []rune("number 42 is the answer")
	var result []rune
	matches := rex.Runes(input).Replace(&result)
	for match := range matches {
		first := unicode.ToUpper(match.Content[0])
		newVal := append([]rune{first}, match.Content[1:]...)
		match.ReplaceLiteral(newVal)
	}
	fmt.Println(string(result))
	//Output: Number 42 Is The Answer
}

func ExampleRReplacement_ReplaceTemplate() {
	rex := regexer.New(`(is|the)`)
	input := []rune("number 42 is the answer")
	var result []rune
	matches := rex.Runes(input).Replace(&result)
	for match := range matches {
		template := []rune(`[$1]`)
		match.ReplaceTemplate(template)
	}
	fmt.Println(string(result))
	//Output: number 42 [is] [the] answer
}

func ExampleRReplacement_ReplaceFunc() {
	rex := regexer.New(`[a-z]+`)
	input := []rune("number 42 is the answer")
	var result []rune
	matches := rex.Runes(input).Replace(&result)
	for match := range matches {
		match.ReplaceFunc(func(b []rune) []rune {
			first := unicode.ToUpper(b[0])
			return append([]rune{first}, b[1:]...)
		})
	}
	fmt.Println(string(result))
	//Output: Number 42 Is The Answer
}
