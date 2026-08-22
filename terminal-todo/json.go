package main

import (
	"encoding/json"
	"log"
	"os"
)

func SaveTodos(todos []Todo) { // function to save todos to a json file
	data, err := json.MarshalIndent(todos, "", "  ") // marshal todos from slice into json, Marshal indent takes 3 args
	if err != nil {
		log.Fatal(err)
		return
	}
	err = os.WriteFile("todos.json", data, 0000) // os.writefile takes in a file path, data, and unix permissions, though dosent apply on windows so we use 0000!
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Todos have saved successfully.") // success message

}
