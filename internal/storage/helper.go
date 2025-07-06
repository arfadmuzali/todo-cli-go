package storage

func getNextId(todos []Todo) int {
	maxId := 0
	for _, todo := range todos {
		if todo.Id > maxId {
			maxId = todo.Id
		}
	}
	return maxId + 1
}
