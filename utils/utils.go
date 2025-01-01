package utils

import (
	"net/http"
	"os"
	"strings"
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
