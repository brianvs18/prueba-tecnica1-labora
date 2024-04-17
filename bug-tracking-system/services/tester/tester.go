package tester

import (
	"prueba-tecnica1-labora/bug-tracking-system/models"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
	"prueba-tecnica1-labora/bug-tracking-system/storage"
)

func CreateBug(description string, taskId string) {
	bug := models.Bug{
		Id:          shared.GenerateId(),
		Description: description,
		Status:      shared.Status[1],
		TaskId:      taskId,
	}
	storage.CreateBug(bug)
}

func UpdateBugStatus(id string, choice int) {
	storage.ChangeBugStatus(id, shared.Status[choice])
}

func GetAllCompletedTasks() {
	storage.GetAllTasksByStatus(2)
}

func GetAllBugs() {
	storage.GetAllBugs()
}
