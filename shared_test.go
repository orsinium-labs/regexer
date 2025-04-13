package regexer_test

import (
	"fmt"

	"github.com/orsinium-labs/regexer"
)

func ExampleSubs_At() {
	rex := regexer.New(`([a-z.]+)@([a-z.]+)`)
	input := "my email is mail@example.com, text me"
	matches := rex.String(input).Find()
	for match := range matches {
		username := match.Subs.At(1).Content
		domain := match.Subs.At(2).Content
		fmt.Printf("username: %s; domain: %s", username, domain)
	}
	//Output: username: mail; domain: example.com
}

func ExampleSubs_Slice() {
	rex := regexer.New(`([a-z.]+)@([a-z.]+)`)
	input := "my email is mail@example.com, text me"
	matches := rex.String(input).Find()
	for match := range matches {
		subs := match.Subs.Slice()
		username := subs[1].Content
		domain := subs[2].Content
		fmt.Printf("username: %s; domain: %s", username, domain)
	}
	//Output: username: mail; domain: example.com
}

func ExampleSubs_Iter() {
	rex := regexer.New(`([a-z.]+)@([a-z.]+)`)
	input := "my email is mail@example.com, text me"
	matches := rex.String(input).Find()
	for match := range matches {
		for sub := range match.Subs.Iter() {
			fmt.Println(sub.Content)
		}
	}
	//Output:
	// mail@example.com
	// mail
	// example.com
}
