package services

import (
	"finalExam/controllers/input"
	"finalExam/storage"
)

const categoriesFile = "data/categories.json"

var Categories []input.Category

func LoadCategories() error {
	return storage.LoadJSON(categoriesFile, &Categories)
}

func SaveCategories() error {
	return storage.SaveJSON(categoriesFile, &Categories)
}

func LoadDefaultCategories(userID string) {
	defaults := []string{"Food", "Transportation", "Bills", "Shop", "Entertaiment", "other"}
	for _, c := range defaults {
		Categories = append(Categories, input.Category{Name: c, UserId: userID})
	}
	SaveCategories()
}

func AddCategory(userID, name string) error {
	for _, c := range Categories {
		if c.Name == name && c.UserId == userID {
		}
	}

	Categories = append(Categories, input.Category{Name: name, UserId: userID})
	return SaveCategories()
}

func GetUserCategories(userID string) []input.Category {
	var res []input.Category
	for _, c := range Categories {
		if c.UserId == userID {
			res = append(res, c)
		}
	}
	return res
}
