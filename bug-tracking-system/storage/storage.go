package storage

import (
	"fmt"
	"os"
	"prueba-tecnica1-labora/bug-tracking-system/models"
	"prueba-tecnica1-labora/bug-tracking-system/shared"

	"github.com/olekukonko/tablewriter"
)

var (
	tasks = []models.Task{
		{
			Id:     "3b98053d",
			Name:   "Task one",
			Status: "PENDING",
		},
		{
			Id:     "56aa1d7b",
			Name:   "Task two",
			Status: "PENDING",
		},
	}
	bugs = []models.Bug{
		{
			Id:          "71122e3e",
			Description: "Lorem ipsum dolor sit amet",
			Status:      "PENDING",
			TaskId:      "3b98053d",
		},
		{
			Id:          "d05abacd",
			Description: "Lorem ipsum dolor sit amet",
			Status:      "PENDING",
			TaskId:      "56aa1d7b",
		},
	}
)

func CreateTask(task models.Task) {
	tasks = append(tasks, task)
	fmt.Println("Tarea creada correctamente")
}

func CreateBug(bug models.Bug) {
	bugs = append(bugs, bug)
	fmt.Println("Bug creado correctamente")
}

func GetAllTasks() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Status"})

	for _, task := range tasks {
		taskData := []string{
			task.Id,
			task.Name,
			task.Status,
		}
		table.Append(taskData)
	}

	table.Render()
}

func GetAllTasksByStatus(status int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Status"})

	for _, task := range tasks {
		if task.Status == shared.Status[status] {
			taskData := []string{
				task.Id,
				task.Name,
				task.Status,
			}
			table.Append(taskData)
		}
	}

	table.Render()
}

func GetAllBugs() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Description", "Status", "TaskId"})

	for _, bug := range bugs {
		bugData := []string{
			bug.Id,
			bug.Description,
			bug.Status,
			bug.TaskId,
		}
		table.Append(bugData)
	}

	table.Render()
}

func ChangeTaskStatus(id string, status string) {
	for idx, task := range tasks {
		if task.Id == id {
			taskUpdate := task.UpdateStatus(status)
			tasks[idx] = taskUpdate
			fmt.Println("El estado de la tarea se ha actualizado correctamente")
			return
		}
	}
	fmt.Println("La tarea no existe, intente nuevamente")
}

func ChangeBugStatus(id string, status string) {
	for idx, bug := range bugs {
		if bug.Id == id {
			bugUpdate := bug.UpdateStatus(status)
			bugs[idx] = bugUpdate
			fmt.Println("El estado del bug se ha actualizado correctamente")
		}
	}
	fmt.Println("El bug no existe, intente nuevamente")
}
