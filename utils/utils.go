package utils

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func NewErrorResponse(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

func NewSuccessResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

// format: foo/bar/baz NOT /foo/bar/baz
func CreateDirectoryIfNotExists(path string) {
	paths := strings.Split(path, "/")
	createDir(paths)
}

func createDir(paths []string) {
	for i := 0; i < len(paths); i++ {
		dirPath := strings.Join(paths[:i+1], "/")
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			os.Mkdir(dirPath, 0755)
		}
	}
}

func GenerateAPIKey() (string, error) {
	keyLength := 32

	key := make([]byte, keyLength)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	apiKey := base64.URLEncoding.EncodeToString(key)
	return apiKey, nil
}

func HashKey(key string) (string, error) {
	hashedKey, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedKey), nil
}

func DeleteDirectory(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
