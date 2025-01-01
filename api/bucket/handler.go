package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rudransh-shrivastava/self-space/config"
	"github.com/rudransh-shrivastava/self-space/utils"
)

type Bucket struct{}

func (b *Bucket) Upload(w http.ResponseWriter, r *http.Request) {
	bucketName, ok := mux.Vars(r)["bucketName"]
	if !ok || bucketName == "" {
		utils.NewErrorResponse(w, "bucket name is required", http.StatusBadRequest)
		return
	}

	fullFilePath, ok := mux.Vars(r)["filePath"]
	if !ok || fullFilePath == "" {
		utils.NewErrorResponse(w, "file path is required", http.StatusBadRequest)
		return
	}

	pathArray := strings.Split(fullFilePath, "/")
	fileName := pathArray[len(pathArray)-1]
	filePath := strings.Join(pathArray[:len(pathArray)-1], "/")

	utils.CreateDirectoryIfNotExists(config.Envs.BucketPath + bucketName + "/" + filePath)

	finalFilePath := config.Envs.BucketPath + bucketName + "/" + filePath + "/" + fileName

	outFile, err := os.Create(finalFilePath)
	if err != nil {
		utils.NewErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer outFile.Close()

	buf := make([]byte, config.Envs.BufferSize)
	var totalBytes int64

	// steam the file to disk
	for {
		n, readErr := r.Body.Read(buf)
		if readErr == io.EOF {
			bytesWritten, err := outFile.Write(buf[:n])
			if err != nil {
				utils.NewErrorResponse(w, err.Error(), http.StatusInternalServerError)
			}
			totalBytes += int64(bytesWritten)
			break
		}
		if readErr != nil {
			utils.NewErrorResponse(w, readErr.Error(), http.StatusInternalServerError)
			return
		}

		bytesWritten, err := outFile.Write(buf[:n])

		if err != nil {
			utils.NewErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
		totalBytes += int64(bytesWritten)
	}

	if totalBytes == 0 {
		utils.NewErrorResponse(w, "No data received", http.StatusBadRequest)
		return
	}

	utils.NewSuccessResponse(w, fmt.Sprintf("File uploaded successfully to: %s/%s with size %d", bucketName, fullFilePath, totalBytes))
}
