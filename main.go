package main

import (
	"github.com/davecgh/go-spew/spew"
)

type DataSource struct {
	state string
	data  [] string
}

type Node struct {
	left     *Node
	right    *Node
	operator string
	data     *DataSource
}

type Tree struct {
	root *Node
}

func unionData(dataLeft *DataSource, dataRight *DataSource) *DataSource {
	var result = &DataSource{
		state: "ready",
		data: []string{},
	}

	if dataLeft != nil {
		result.data = append(result.data, dataLeft.data...)
	}

	if dataRight != nil {
		result.data = append(result.data, dataRight.data...)
	}

	return result;
}

func processTree(node *Node) *DataSource {
	// var tData []string
	if node.left == nil && node.right == nil {
		return node.data
	} else {

		if node.operator == "union" {
			return unionData(processTree(node.left), processTree(node.right))
		}

		return nil
	}
}

func main() {

	var dataSource1 = &DataSource{
		state: "ready",
		data: []string{
			"first item1",
			"second item1",
		},
	}

	var dataSource2 = &DataSource{
		state: "ready",
		data: []string{
			"first item2",
			"second item2",
		},
	}

	var dataSource3 = &DataSource{
		state: "ready",
		data: []string{
			"first item3",
			"second item3",
		},
	}

	var node1 = &Node{left: nil, right: nil, operator: "", data: dataSource1}
	var node2 = &Node{left: nil, right: nil, operator: "", data: dataSource2}
	var node3 = &Node{left: nil, right: nil, operator: "", data: dataSource3}
	var node4 = &Node{left: node1, right: node2, operator: "union", data: nil}
	var root = &Node{left: node4, right: node3, operator: "union", data: nil}

	spew.Dump(processTree(root))

}
