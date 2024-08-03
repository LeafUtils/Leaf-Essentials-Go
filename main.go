package main

import (
	"fmt"
	"leafmcbe/database"
	"strconv"
)

func main() {
	db := database.NewDatabase()

	doc := db.InsertDocument(map[string]interface{}{"key1": "value1"})
	fmt.Println(strconv.FormatInt(doc.ID, 10))
	str, _ := db.ExportToJSON()
	fmt.Println(str)
}
