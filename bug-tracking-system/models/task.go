package models

type Task struct {
	Id     string
	Name   string
	Status string
}

func (t Task) UpdateStatus(status string) Task {
	t.Status = status
	return t
}
