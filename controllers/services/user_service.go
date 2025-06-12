package services

import (
	"errors"
	"finalExam/controllers/input"
	"finalExam/storage"

	"github.com/google/uuid"
)

const usersFile = "data/users.json"

var Users []input.AddUserInput

// بارگذاری کاربران از فایل
func LoadUsers() error {
	return storage.LoadJSON(usersFile, &Users)
}

// ذخیره کاربران در فایل
func SaveUsers() error {
	return storage.SaveJSON(usersFile, &Users)
}

// اضافه کردن کاربر جدید
func AddUser(username, email string) (input.AddUserInput, error) {
	// چک کردن یکتا بودن نام کاربری
	for _, u := range Users {
		if u.UserName == username {
			return input.AddUserInput{}, errors.New("نام کاربری تکراری است")
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

// لیست تمام کاربران
func ListUsers() []input.AddUserInput {
	return Users
}

// جستجو کاربر بر اساس ID
func FindUserByID(id string) (*input.AddUserInput, error) {
	for i, u := range Users {
		if u.ID == id {
			return &Users[i], nil
		}
	}
	return nil, errors.New("کاربر یافت نشد")
}
