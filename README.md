# regexer

[ [📄 docs](https://pkg.go.dev/github.com/orsinium-labs/regexer) ] [ [🐙 github](https://github.com/orsinium-labs/regexer) ]

Go package with more powerful, flexible, and safe API for regular expressions.

Features:

* Type-safe
* Lazy iteration.
* Supports strings, bytes, and runes as input.
* The same generic API for all inputs.
* Everything possible with stdin regexp: find matches, find submatches, replace, replace with a template.

## Installation

```bash
go get github.com/orsinium-labs/regexer
```

## Usage

Find and print all words in the text and their position.

```go
rex := regexer.New(`\w+`)
input := "never gonna give you up"
matches := rex.String(input).Find()
for match := range matches {
    fmt.Println(match.Span.Start, match.Content)
}
```

The same but for a slice of bytes:

```go
rex := regexer.New(`\w+`)
input := []byte("never gonna give you up")
matches := rex.Bytes(input).Find()
for match := range matches {
    fmt.Println(match.Span.Start, string(match.Content))
}
```

In both cases, `matches` is a lazy iterator. It doesn't require to allocate memory for all matches and if you stop iteration, it will stop scanning the input.

Replacing has very similar API:

```go
rex := regexer.New(`\w+`)
input := "number 42 is the answer"
var result string
matches := rex.String(input).Replace(&result)
for match := range matches {
    template := string(`[$1]`)
    match.ReplaceTemplate(template)
}
fmt.Println(result)
// Output: [number] 42 [is] [the] [answer]
```

Accessing submatches:

```go
rex := regexer.New(`([a-z.]+)@([a-z.]+)`)
input := "my email is mail@example.com, text me"
matches := rex.String(input).Find()
for match := range matches {
    username := match.Subs.At(1).Content
    domain := match.Subs.At(2).Content
    fmt.Printf("username: %s; domain: %s", username, domain)
}
```
