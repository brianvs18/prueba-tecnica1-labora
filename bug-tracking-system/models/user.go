package models

import "prueba-tecnica1-labora/bug-tracking-system/shared"

type User struct {
	Id       string
	Username string
	Password string
	Role     string
}

var (
	Users = [3]User{
		{
			Id:       shared.GenerateId(),
			Username: "admin",
			Password: "root",
			Role:     "MANAGER",
		},
		{
			Id:       shared.GenerateId(),
			Username: "dev",
			Password: "develop",
			Role:     "DEVELOPER",
		},
		{
			Id:       shared.GenerateId(),
			Username: "qa",
			Password: "tester",
			Role:     "TESTER",
		},
	}
)
