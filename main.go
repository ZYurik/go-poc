package main

import (
	"sort"

	connectors "./connectors"
	dataTypes "./libs"
	"github.com/davecgh/go-spew/spew"
)

func unionData(dataLeft *dataTypes.DataSource, dataRight *dataTypes.DataSource) *dataTypes.DataSource {
	var result = &dataTypes.DataSource{
		State: "ready",
		Data:  [][]interface{}{},
	}

	if dataLeft != nil {
		result.Data = append(result.Data, dataLeft.Data...)
	}

	if dataRight != nil {
		result.Data = append(result.Data, dataRight.Data...)
	}

	return result
}

func joinData(dataLeft *dataTypes.DataSource, dataRight *dataTypes.DataSource) *dataTypes.DataSource {
	var result = &dataTypes.DataSource{
		State: "ready",
		Data:  [][]interface{}{},
	}

	var index = 0

	for leftIndex := range dataLeft.Data {
		var row = dataLeft.Data[leftIndex]

		index = sort.Search(len(dataRight.Data), func(searchIndex int) bool {
			return dataRight.Data[searchIndex][0] != row[1]
		})
		row = append(row, dataRight.Data[index-1][1])

		result.Data = append(result.Data, row)
	}

	return result
}

func processTree(node *dataTypes.Node) *dataTypes.DataSource {
	if node.Left == nil && node.Right == nil {
		return node.Data
	} else {
		if node.Operator == "union" {
			return unionData(processTree(node.Left), processTree(node.Right))
		}

		if node.Operator == "join" {
			return joinData(processTree(node.Left), processTree(node.Right))
		}

		return nil
	}
}

func main() {
	var dataSource1 = connectors.GetData("select * from table1")
	var dataSource2 = connectors.GetData("select * from table2")
	var dataSource3 = connectors.GetData("select * from table3")
	var dataSource4 = connectors.GetData("select * from table4")
	var node1 = &dataTypes.Node{Left: nil, Right: nil, Operator: "", Data: dataSource1}
	var node2 = &dataTypes.Node{Left: nil, Right: nil, Operator: "", Data: dataSource2}
	var node3 = &dataTypes.Node{Left: nil, Right: nil, Operator: "", Data: dataSource3}
	var node4 = &dataTypes.Node{Left: node1, Right: node2, Operator: "union", Data: nil}
	var node5 = &dataTypes.Node{Left: node4, Right: node3, Operator: "union", Data: nil}
	var node6 = &dataTypes.Node{Left: nil, Right: nil, Operator: "", Data: dataSource4}
	var root = &dataTypes.Node{Left: node5, Right: node6, Operator: "join", Data: nil}

	spew.Dump(processTree(root))
	spew.Dump(connectors.GetData("select * from users"))
	spew.Dump(connectors.GetData("select * from groups"))
}
