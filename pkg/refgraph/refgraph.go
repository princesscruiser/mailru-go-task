package refgraph

import (
	"math/rand"
	"time"
)

type Node struct {
	Id      int
	Value   int
	Sum     int
	Childs  []*Node
	Parents []*Node
}

type GenSettings struct {
	Cycles    bool
	MaxDepth  uint
	MaxChilds uint
	NodeCount uint
}

var (
	lastId int
)

func Generate(root *Node, genSettings GenSettings) error {
	if err := checkArgs(root, genSettings); err != nil {
		return err
	}

	lastId = 0

	if err := generateRecursively(root, int(genSettings.MaxChilds), int(genSettings.MaxDepth), int(genSettings.NodeCount)); err != nil {
		return err
	}
	return nil
}

func checkArgs(root *Node, genSettings GenSettings) error {
	return nil
}

func generateRecursively(parentNode *Node, maxChilds, depthLeft, nodesLeft int) error {
	rand.Seed(time.Now().UTC().UnixNano())
	if nodesLeft > 0 {
		chCount := rand.Intn(maxChilds + 1)
		childs := make([]*Node, 0, chCount)
		for i := 0; i < chCount; i++ {
			rand.Seed(time.Now().UTC().UnixNano())
			value := rand.Intn(100)
			parents := make([]*Node, 0, 1)
			parents = append(parents, parentNode)
			child := &Node{
				Id:      lastId,
				Value:   value,
				Parents: parents,
			}
			childs = append(childs, child)
			lastId++
		}
		nodesLeft -= chCount
		depthLeft--
		if depthLeft > 0 {
			for _, v := range childs {
				generateRecursively(v, maxChilds, depthLeft, nodesLeft)
			}
		}
	}
	return nil
}
