package models

type Bug struct {
	Id          string
	Description string
	Status      string
	TaskId      string
}

func (b Bug) UpdateStatus(status string) Bug {
	b.Status = status
	return b
}
