package auth

import "github.com/google/uuid"

type User struct {
	Id       string
	Username string
	Password string
	Role     string
}

var (
	Users = [3]User{
		{
			Id:       generateId(),
			Username: "admin",
			Password: "root",
			Role:     "MANAGER",
		},
		{
			Id:       generateId(),
			Username: "dev",
			Password: "develop",
			Role:     "DEVELOPER",
		},
		{
			Id:       generateId(),
			Username: "qa",
			Password: "tester",
			Role:     "TESTER",
		},
	}
)

func generateId() string {
	return uuid.New().String()[:8]
}
