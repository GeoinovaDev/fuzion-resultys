package node

// Node ...
type Node struct {
	Kind      string
	Value     string
	Parent    *Node
	Childrens []*Node
}

// New ...
func New() *Node {
	return &Node{
		Kind:      "unknown",
		Childrens: []*Node{},
	}
}

// GetArray ...
func (n *Node) GetArray(indexes ...int) []*Node {
	node := n.Get(indexes...)

	if node == nil {
		return nil
	}

	return node.Childrens
}

// GetValue ...
func (n *Node) GetValue(indexes ...int) string {
	node := n.Get(indexes...)

	if node == nil {
		return ""
	}

	return node.Value
}

// Get ...
func (n *Node) Get(indexes ...int) *Node {
	_node := n

	for i := 0; i < len(indexes); i++ {
		index := indexes[i]

		if _node == nil {
			return nil
		}

		if _node.Childrens == nil {
			return nil
		}

		if index >= len(_node.Childrens) {
			return nil
		}

		_node = _node.Childrens[index]
	}

	return _node
}

// Exist ...
func (n *Node) Exist(indexes ...int) bool {
	_node := n

	for i := 0; i < len(indexes); i++ {
		index := indexes[i]

		if _node.Childrens == nil {
			return false
		}

		if index >= len(_node.Childrens) {
			return false
		}

		_node = _node.Childrens[index]
	}

	return true
}
