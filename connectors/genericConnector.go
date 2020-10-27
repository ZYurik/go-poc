package connectors

import (
	mockedData "../data"
	dataTypes "../libs"
)

func GetData(query string) *dataTypes.DataSource {
	if query == "select * from users" {
		return mockedData.GetUsers()
	} else if query == "select * from groups " {
		return mockedData.GetGroups()
	} else if query == "select * from table1" {
		return &dataTypes.DataSource{
			State: "ready",
			Data: [][]interface{}{
				{1, 1, "first item2"},
				{2, 2, "second item2"},
			},
		}
	} else if query == "select * from table2" {
		return &dataTypes.DataSource{
			State: "ready",
			Data: [][]interface{}{
				{3, 1, "first item2"},
				{4, 2, "second item2"},
			},
		}
	} else if query == "select * from table3" {
		return &dataTypes.DataSource{
			State: "ready",
			Data: [][]interface{}{
				{5, 1, "first item3"},
				{6, 2, "second item3"},
			},
		}

	} else if query == "select * from table4" {
		return &dataTypes.DataSource{
			State: "ready",
			Data: [][]interface{}{
				{1, "tip"},
				{2, "top"},
			},
		}
	}

	return nil
}
