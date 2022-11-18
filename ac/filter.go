package ac

import (
	"bufio"
	"io"
)

type acFilter struct {
	trie *trie
}

func NewAcFilter() *acFilter {
	return &acFilter{
		trie: NewTrie(),
	}
}

func (a *acFilter) LoadFromFile(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		a.trie.Add(string(line))
	}
	a.trie.InitAC()
	return nil
}

func (a *acFilter) Filter(text string) string {
	return a.trie.Filter(text)
}

func (a *acFilter) Replace(text string, repl rune) string {
	return a.trie.Replace(text, repl)
}

func (a *acFilter) FindIn(text string) (bool, string) {
	return a.trie.FindIn(text)
}

func (a *acFilter) FindAll(text string) []string {
	return a.trie.FindAll(text)
}

func (a *acFilter) Validate(text string) (bool, string) {
	return a.trie.Validate(text)
}
