package test

import "github.com/mrkaspa/iloopsapi/models"

func defaultUser() models.UserLogin {
	return models.UserLogin{Email: "michel.ingesoft@gmail.com", Password: "h1h1h1h1h1h1"}
}

func anotherUser() models.UserLogin {
	return models.UserLogin{Email: "angelbotto@gmail.com", Password: "h1h1h1h1h1h1"}
}

func defaultProject() models.Project {
	return models.Project{Name: "demo"}
}
