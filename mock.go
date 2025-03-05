package test

import "errors"

type User struct {
	ID   int
	Name string
}

type UserService interface {
	GetUserByID(id int) (User, error)
}

func FetchUserDetails(service UserService, id int) (string, error) {
	user, err := service.GetUserByID(id)
	if err != nil {
		return "", errors.New("user not found")
	}
	return "User: " + user.Name, nil
}
