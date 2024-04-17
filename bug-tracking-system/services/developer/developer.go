package developer

import (
	"prueba-tecnica1-labora/bug-tracking-system/shared"
	"prueba-tecnica1-labora/bug-tracking-system/storage"
)

func UpdateTaskStatus(id string, choice int) {
	storage.ChangeTaskStatus(id, shared.Status[choice])
}

func GetAllPendingTasks() {
	storage.GetAllTasksByStatus(1)
}
