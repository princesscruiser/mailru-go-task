package refgraph

import "errors"

type Node struct {
	Id      int
	value   int
	Sum     int
	Childs  []*Node
	Parents []*Node
}

type GenerationMode int

const (
	AcyclicMode = GenerationMode(0)
	CyclicMode  = GenerationMode(1)
)

func GenerateGraph(root *Node, genMode GenerationMode, depth int, childCount int) (map[string]interface{}, error) {
	err := checkArgs(root, genMode, depth, childCount)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func checkArgs(root *Node, genMode GenerationMode, depth int, childCount int) error {
	if root == nil {
		return errors.New("root variable is nil")
	}
	if genMode > 1 {
		return errors.New("incorrect generation mode")
	}
	if depth < 1 {
		return errors.New("incorrect depth variable")
	}

	if childCount < 0 {
		return errors.New("child count set below zero")
	}
	return nil
}
