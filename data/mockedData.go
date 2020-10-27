package mockedData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	dataTypes "../libs"
)

type users []struct {
	id     int    `json:"id"`
	name   string `json:"name"`
	email  string `json:"email"`
	groups []int  `json:"groups"`
}

type groups []struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

func GetUsers() *dataTypes.DataSource {

	// Open our jsonFile
	jsonFile, err := os.Open("./data/users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result users
	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		log.Println(err)
	}

	var source = &dataTypes.DataSource{
		State: "ready",
		Data:  [][]interface{}{},
	}

	for userIndex := range result {
		var row = []interface{}{
			result[userIndex].id,
			result[userIndex].name,
			result[userIndex].email,
			result[userIndex].groups,
		}
		source.Data = append(source.Data, row)
	}
	return source
}

func GetGroups() *dataTypes.DataSource {
	jsonFile, err := os.Open("./data/groups.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result groups
	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		log.Println(err)
	}

	var source = &dataTypes.DataSource{
		State: "ready",
		Data:  [][]interface{}{},
	}

	for userIndex := range result {
		var row = []interface{}{
			result[userIndex].id,
			result[userIndex].name,
		}
		source.Data = append(source.Data, row)
	}
	return source
}
