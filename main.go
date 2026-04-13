package main 

import (
	"fmt"
)

type Todo struct {
	Id int `json:"id"`
	Task string `json:"task"`
}

func addTodo(task string) {
	// --- SQL STEP ---
	sqliteId := 1
	
	// VECTOR DB STEP
	payload := map[string]interface{}{
		"ids": []int{sqliteId},
		"documents": []string{task},
		"metadatas": []map[string]interface{}{
			{"id": sqliteId},
		},
	}

	sendToChroma("/add", payload)
	fmt.Println("SAVED TO BOTH CHROMA AND SQLITE")
}

func SearchTodo(query string) string {
	serchPayload := map[string]interface{}{
		"query_texts": []string{query},
		"n_results": 1,
	}

	response := sendToChroma("/search", serchPayload)

	sqliteId := response["ids"][0].(int)

	var task string

	task = "Buy Milk"

	return task
}

func sendToChroma(endPoint string, payload interface{}) map[string]interface{} {
	return map[string]interface{}{
		"MetaData": []map[string]interface{}{
			{"sqlite_id": 1}
		}
	}
}


func main () {
	addTodo("buy Milk")

	results := SearchTodo("i need something for the fridge")

	fmt.Println("hellow world")
}