package storage

import (
	"encoding/json"
	"io"
	"os"
)

// بارگذاری داده‌ها از فایل JSON
func LoadJSON(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, data)
}

// ذخیره داده‌ها در فایل JSON
func SaveJSON(filename string, data interface{}) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}
