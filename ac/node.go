package ac

// Node the node of the trie tree
type Node struct {
	isRootNode bool
	isPathEnd  bool           // Whether it is the end of a matching path
	character  rune           // the character of the node
	children   map[rune]*Node // the children of the node
	fail       *Node          // the fail pointer in the Aho-Corasick automaton
	count      int            // the length of the matched word
}

// NewNode new a node
func NewNode(character rune) *Node {
	return &Node{
		character: character,
		children:  make(map[rune]*Node, 0),
	}
}

// NewRootNode new a root node
func NewRootNode(character rune) *Node {
	return &Node{
		isRootNode: true,
		character:  character,
		children:   make(map[rune]*Node, 0),
	}
}

// IsLeafNode whether the node is a leaf node
func (node *Node) IsLeafNode() bool {
	return len(node.children) == 0
}

// IsRootNode whether the node is a root node
func (node *Node) IsRootNode() bool {
	return node.isRootNode
}

// IsPathEnd whether the node is the end of a matching path
func (node *Node) IsPathEnd() bool {
	return node.isPathEnd
}

// SoftDel soft delete the node
func (node *Node) SoftDel() {
	node.isPathEnd = false
}

// IsChild whether the rune r is the child of the node
func (node *Node) IsChild(r rune) bool {
	if node.children == nil {
		return false
	} else {
		_, found := node.children[r]
		return found
	}
}
