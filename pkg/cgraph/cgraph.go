package cgraph

type Node struct {
	Id      int
	value   interface{}
	Sum     int
	Childs  []*Node
	Parents []*Node
}

func (n *Node) SetValue(v interface{}) error {
	n.value = v
	return nil
}

func (n *Node) GetValue() (interface{}, error) {
	return n.value, nil
}
