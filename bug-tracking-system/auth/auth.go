package auth

import "prueba-tecnica1-labora/bug-tracking-system/models"

func GetUserByUsernameAndPassword(username string, pass string) models.User {
	for idx, user := range models.Users {
		if user.Username == username && user.Password == pass {
			return models.Users[idx]
		}
	}
	return models.User{}
}
