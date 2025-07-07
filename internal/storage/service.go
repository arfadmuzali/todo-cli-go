package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

func GetDataFilePath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("Failed get storage file: " + err.Error())
	}

	appDir := filepath.Join(configDir, "todo")
	os.MkdirAll(appDir, 0755)

	return filepath.Join(appDir, "todos.json")
}

func LoadTodos() ([]Todo, error) {
	pathFile := GetDataFilePath()
	if _, err := os.Stat(pathFile); os.IsNotExist(err) {
		empty := []Todo{}
		jsonData, _ := json.Marshal(empty)
		os.WriteFile(pathFile, jsonData, 0644)
		return empty, nil
	}

	rawTodos, err := os.ReadFile(pathFile)
	if err != nil {
		return nil, err
	}
	var todos []Todo
	json.Unmarshal(rawTodos, &todos)
	return todos, nil
}

func AddTodo(name string) error {
	pathFile := GetDataFilePath()
	if _, err := os.Stat(pathFile); os.IsNotExist(err) {
		empty := []Todo{}
		jsonData, _ := json.Marshal(empty)
		os.WriteFile(pathFile, jsonData, 0644)
		return nil
	}

	rawTodos, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}

	var todos []Todo
	json.Unmarshal(rawTodos, &todos)
	todos = append(todos, Todo{Id: getNextId(todos), Description: name, Status: "to-do", CreatedAt: time.Now()})
	return SaveTodos(todos)
}

func DeleteTodo(id int) error {
	pathFile := GetDataFilePath()

	rawTodos, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}

	var todos []Todo
	json.Unmarshal(rawTodos, &todos)

	found := false
	var updateTodo []Todo
	for _, value := range todos {
		if value.Id == id {
			found = true
			continue
		}
		updateTodo = append(updateTodo, value)
	}

	if !found {
		return errors.New("Can't find todo")
	}

	return SaveTodos(updateTodo)
}

func DeleteAll() error {
	pathFile := GetDataFilePath()

	_, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}
	return SaveTodos([]Todo{})
}

func UpdateTodoDescription(id int, newContent string) error {
	pathFile := GetDataFilePath()

	rawTodos, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}

	var todos []Todo
	json.Unmarshal(rawTodos, &todos)

	err = errors.New("Cant find todo")
	var updateTodo []Todo
	for _, value := range todos {
		if value.Id == id {
			updateTodo = append(updateTodo, Todo{Id: value.Id, Description: newContent, CreatedAt: value.CreatedAt, Status: value.Status})
			err = nil
		} else {
			updateTodo = append(updateTodo, value)
		}
	}

	if err != nil {
		return err
	}

	return SaveTodos(updateTodo)
}

func UpdateTodosStatus(id int) error {
	pathFile := GetDataFilePath()

	rawTodos, err := os.ReadFile(pathFile)
	if err != nil {
		return err
	}

	var todos []Todo
	json.Unmarshal(rawTodos, &todos)

	err = errors.New("Cant find todo")
	var updateTodo []Todo
	for _, value := range todos {
		var newStatus string
		if value.Id == id {
			err = nil
			if value.Status == "to-do" {
				newStatus = "done"
			} else {
				newStatus = "to-do"
			}
			updateTodo = append(updateTodo, Todo{Id: value.Id, Description: value.Description, CreatedAt: value.CreatedAt, Status: newStatus})
		} else {
			updateTodo = append(updateTodo, value)
		}
	}

	if err != nil {
		return err
	}

	return SaveTodos(updateTodo)
}

func SaveTodos(todos []Todo) error {

	jsonData, err := json.Marshal(todos)

	if err != nil {
		return err
	}

	return os.WriteFile(GetDataFilePath(), jsonData, 0644)

}
