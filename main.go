package main

import (
	"sort"
	"github.com/davecgh/go-spew/spew"
)

type DataSource struct {
	state string
	data  [][]interface{}
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
		data: [][]interface{}{},
	}

	if dataLeft != nil {
		result.data = append(result.data, dataLeft.data...)
	}

	if dataRight != nil {
		result.data = append(result.data, dataRight.data...)
	}

	return result;
}

func joinData(dataLeft *DataSource, dataRight *DataSource) *DataSource {
	var result = &DataSource{
		state: "ready",
		data: [][]interface{}{},
	}

	var index = 0

	for leftIndex := range dataLeft.data {
		var row = dataLeft.data[leftIndex];
	
		
		index = sort.Search(len(dataRight.data), func(searchIndex int) bool {
			return dataRight.data[searchIndex][0] != row[1]
		})
		row = append(row, dataRight.data[index-1][1]) ;

		result.data = append(result.data, row);
	}

	return result;
}

func processTree(node *Node) *DataSource {
	if node.left == nil && node.right == nil {
		return node.data
	} else {

		if node.operator == "union" {
			return unionData(processTree(node.left), processTree(node.right))
		}		
		
		if node.operator == "join" {
			return joinData(processTree(node.left), processTree(node.right))
		}

		return nil
	}
}

func main() {

	var dataSource1 = &DataSource{
		state: "ready",
		data: [][]interface{}{
			{1, 1, "first item1"},
			{2, 2, "second item1"},
		},
	}

	var dataSource2 = &DataSource{
		state: "ready",
		data: [][]interface{}{
			{3, 1, "first item2"},
			{4, 2, "second item2"},
		},
	}

	var dataSource3 = &DataSource{
		state: "ready",
		data: [][]interface{} {
			{5, 1, "first item3"},
			{6, 2, "second item3"},
		},
	}

	var dataSource4 = &DataSource{
		state: "ready",
		data: [][]interface{} {
			{1, "tip"},
			{2, "top"},
		},
	}

	var node1 = &Node{left: nil, right: nil, operator: "", data: dataSource1}
	var node2 = &Node{left: nil, right: nil, operator: "", data: dataSource2}
	var node3 = &Node{left: nil, right: nil, operator: "", data: dataSource3}
	var node4 = &Node{left: node1, right: node2, operator: "union", data: nil}	
	var node5 = &Node{left: node4, right: node3, operator: "union", data: nil}
	var node6 = &Node{left: nil, right: nil, operator: "", data: dataSource4}
	var root = &Node{left: node5, right: node6, operator: "join", data: nil}

	spew.Dump(processTree(root))

}
