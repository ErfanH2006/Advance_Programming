package services

import (
	"finalExam/controllers/input"
	"finalExam/storage"
)

const categoriesFile = "data/categories.json"

var Categories []input.Category

// بارگذاری دسته‌بندی‌ها
func LoadCategories() error {
	return storage.LoadJSON(categoriesFile, &Categories)
}

// ذخیره دسته‌بندی‌ها
func SaveCategories() error {
	return storage.SaveJSON(categoriesFile, &Categories)
}

// دسته‌بندی‌های پیش‌فرض
func LoadDefaultCategories(userID string) {
	defaults := []string{"Food", "Transportation", "Bills", "Shop", "Entertaiment", "other"}
	for _, c := range defaults {
		Categories = append(Categories, input.Category{Name: c, UserId: userID})
	}
	SaveCategories()
}

// افزودن دسته‌بندی سفارشی
func AddCategory(userID, name string) error {
	// بررسی تکراری نبودن دسته‌بندی
	for _, c := range Categories {
		if c.Name == name && c.UserId == userID {
		}
	}

	Categories = append(Categories, input.Category{Name: name, UserId: userID})
	return SaveCategories()
}

// مشاهده دسته‌بندی‌های کاربر
func GetUserCategories(userID string) []input.Category {
	var res []input.Category
	for _, c := range Categories {
		if c.UserId == userID {
			res = append(res, c)
		}
	}
	return res
}
