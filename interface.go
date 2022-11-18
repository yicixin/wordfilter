package wordfilter

import "io"

type WordFilter interface {
	LoadFromFile(rd io.Reader) error // load the words from the file

	Filter(text string) string
	Replace(text string, repl rune) string
	FindIn(text string) (bool, string)
	FindAll(text string) []string
	Validate(text string) (bool, string)
}
