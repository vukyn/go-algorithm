package utils

import (
	"os"
	"path/filepath"
)

// Write text to a file.
// Create if not exist, or overwrite the existing file.
//
// Example:
//
//	WriteFile("Hello World", "temp/output.txt")
func WriteFile(input string, filePath string) error {
	data := []byte(input)
	dir, _ := filepath.Split(filePath)

	if _, err := os.Stat(dir); err == nil {
		os.Remove(filePath)
	} else {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	if err := os.WriteFile(filePath, data, 0); err != nil {
		return err
	}
	return nil
}

// Create empty directory if not exist.
//
// Example:
//
//	CreateFolder(""temp/output")
func CreateFolder(folderPath string) error {
	err := os.MkdirAll(folderPath, os.ModePerm)
	return err
}