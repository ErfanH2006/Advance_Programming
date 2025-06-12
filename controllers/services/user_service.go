package services

import (
	"errors"
	"finalExam/controllers/input"
	"finalExam/storage"

	"github.com/google/uuid"
)

const usersFile = "data/users.json"

var Users []input.AddUserInput

func LoadUsers() error {
	return storage.LoadJSON(usersFile, &Users)
}

func SaveUsers() error {
	return storage.SaveJSON(usersFile, &Users)
}

func AddUser(username, email string) (input.AddUserInput, error) {
	for _, u := range Users {
		if u.UserName == username {
			return input.AddUserInput{}, errors.New("the username is duplicate")
		}
	}

	user := input.AddUserInput{
		ID:       uuid.New().String(),
		UserName: username,
		Email:    email,
	}

	Users = append(Users, user)
	err := SaveUsers()
	return user, err
}

func ListUsers() []input.AddUserInput {
	return Users
}

func FindUserByID(id string) (*input.AddUserInput, error) {
	for i, u := range Users {
		if u.ID == id {
			return &Users[i], nil
		}
	}
	return nil, errors.New("user not found")
}
