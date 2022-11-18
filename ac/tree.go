package ac

// trie
type trie struct {
	Root *Node
}

// NewTrie new a trie
func NewTrie() *trie {
	return &trie{
		Root: NewRootNode(0),
	}
}

// InitAC init the Aho-Corasick automaton, BFS,set fail pointer
func (tree *trie) InitAC() {
	queue := newNodeQueue()
	for _, v := range tree.Root.children {
		// root's all children's fail pointer point to root
		v.fail = tree.Root
		queue.Push(v)
	}
	for !queue.IsEmpty() {
		p := queue.Pop()
		failTo := p.fail
		for _, v := range p.children {
			// push the child node to queue
			queue.Push(v)
			for {
				if failTo == nil {
					v.fail = tree.Root
					break
				}
				if failTo.IsChild(v.character) && failTo.children[v.character] != p {
					v.fail = failTo.children[v.character]
					break
				} else {
					failTo = failTo.fail
				}
			}
		}
	}
}

// Add add words to the trie
func (tree *trie) Add(words ...string) {
	for _, word := range words {
		tree.add(word)
	}
}

func (tree *trie) add(word string) {
	var current = tree.Root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if current.IsChild(r) {
			current = current.children[r]
		} else {
			current.children[r] = tree.batchNewNode(runes[position:], len(runes))
			return
		}
	}
	current.isPathEnd = true
	current.count = len(runes)
}

func (tree *trie) batchNewNode(rs []rune, count int) *Node {
	if len(rs) == 1 {
		return &Node{
			isRootNode: false,
			isPathEnd:  true,
			character:  rs[0],
			children:   make(map[rune]*Node),
			count:      count,
		}
	}
	n := &Node{
		isRootNode: false,
		isPathEnd:  false,
		character:  rs[0],
		children:   make(map[rune]*Node),
	}
	n.children[rs[1]] = tree.batchNewNode(rs[1:], count)
	return n
}

// Del delete words from the trie
func (tree *trie) Del(words ...string) {
	for _, word := range words {
		tree.del(word)
	}
}

func (tree *trie) del(word string) {
	var current = tree.Root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if next, ok := current.children[r]; !ok {
			return
		} else {
			current = next
		}

		if position == len(runes)-1 {
			current.SoftDel()
		}
	}
}

// Validate validate the text, return true if the text is valid, otherwise return false and the first invalid word
func (tree *trie) Validate(text string) (bool, string) {
	var (
		current = tree.Root
		runes   = []rune(text)
	)

	for position := 0; position < len(runes); position++ {
		if current.IsChild(runes[position]) {
			r := runes[position]
			if current.children[r].IsPathEnd() {
				return false, string(runes[position-current.children[r].count+1 : position+1])
			}
			current = current.children[r]
		} else {
			if current.fail != nil {
				current = current.fail
				position--
			}
		}
	}
	return true, ""
}

// Replace replace the sensitive words with the given character
func (tree *trie) Replace(text string, character rune) string {
	var (
		current = tree.Root
		runes   = []rune(text)
	)

	for position := 0; position < len(runes); position++ {
		if current.IsChild(runes[position]) {
			r := runes[position]
			if current.children[r].IsPathEnd() {
				// 修改
				for i := position - current.children[r].count + 1; i <= position; i++ {
					runes[i] = character
				}
			}
			current = current.children[r]
		} else {
			if current.fail != nil {
				current = current.fail
				position--
			}
		}
	}
	return string(runes)
}

// Filter filter the sensitive words
func (tree *trie) Filter(text string) string {
	var (
		current     = tree.Root
		runes       = []rune(text)
		length      = len(runes)
		left        = 0
		resultRunes = make([]rune, 0, length)
	)

	for position := 0; position < len(runes); position++ {
		if current.IsChild(runes[position]) {
			r := runes[position]
			if current.children[r].IsPathEnd() {
				resultRunes = append(resultRunes, runes[left:position-current.children[r].count+1]...)
				left = position + 1
			}
			current = current.children[r]
		} else {
			if current.fail != nil {
				current = current.fail
				position--
			}
		}
	}
	return string(resultRunes)
}

// FindIn find the sensitive words in the text
func (tree *trie) FindIn(text string) (bool, string) {
	validated, first := tree.Validate(text)
	return !validated, first
}

// FindAll find all the sensitive words in the text
func (tree *trie) FindAll(text string) []string {
	var (
		current = tree.Root
		runes   = []rune(text)
		matches []string
	)
	for position := 0; position < len(runes); position++ {
		if current.IsChild(runes[position]) {
			r := runes[position]
			child := current.children[r]
			if child.IsPathEnd() {
				matches = append(matches, string(runes[position-child.count+1:position+1]))
				for {
					if child.fail != nil {
						child = child.fail
						if child.IsPathEnd() {
							matches = append(matches, string(runes[position-child.count+1:position+1]))
						} else {
							break
						}
					} else {
						break
					}
				}
			}
			current = current.children[r]
		} else {
			if current.fail != nil {
				current = current.fail
				position--
			}
		}
	}

	return matches
}
