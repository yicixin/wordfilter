package ac

type nodeQueue struct {
	queue []*Node
}

func newNodeQueue() *nodeQueue {
	return &nodeQueue{
		queue: make([]*Node, 0),
	}
}

func (n *nodeQueue) Push(node ...*Node) {
	n.queue = append(n.queue, node...)
}

func (n *nodeQueue) Pop() *Node {
	if len(n.queue) == 0 {
		return nil
	} else {
		t := n.queue[0]
		n.queue[0] = nil // release the reference
		n.queue = n.queue[1:]
		return t
	}
}

func (n *nodeQueue) IsEmpty() bool {
	return len(n.queue) == 0
}
