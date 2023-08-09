package storage

import (
	"github.com/google/uuid"
	"os"
)

func SaveString(data []byte) (string, error) {
	fileName := uuid.NewString()

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	_, err = file.Write(data)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func GetString(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return data, nil
}
