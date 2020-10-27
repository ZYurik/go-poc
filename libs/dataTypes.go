package dataTypes

type DataType [][]interface{}

// type DataSource struct {
// 	State string
// 	Data  *DataType
// }

type DataSource struct {
	State string
	Data  [][]interface{}
}
type Node struct {
	Left     *Node
	Right    *Node
	Operator string
	Data     *DataSource
}

type Tree struct {
	Root *Node
}

type Tables []struct {
	TableName string
	Alias     string
}
