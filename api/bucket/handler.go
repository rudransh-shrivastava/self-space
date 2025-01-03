package bucket

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rudransh-shrivastava/self-space/config"
	"github.com/rudransh-shrivastava/self-space/utils"
)

type Bucket struct {
	BucketStore *BucketStore
}

func (b *Bucket) Upload(w http.ResponseWriter, r *http.Request) {
	bucketName := mux.Vars(r)["bucketName"]
	filePath := r.Header.Get("filePath")
	fileName := r.Header.Get("fileName")

	utils.CreateDirectoryIfNotExists(config.Envs.BucketPath + bucketName + "/" + filePath)

	fullFilePath := filePath + "/" + fileName
	finalFilePath := config.Envs.BucketPath + bucketName + "/" + fullFilePath

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

	utils.NewSuccessResponse(w, fmt.Sprintf("File uploaded successfully \nbucket: %s \npath: %s \nsize: %d", bucketName, fullFilePath, totalBytes))
}

func (b *Bucket) Download(w http.ResponseWriter, r *http.Request) {
	bucketName := mux.Vars(r)["bucketName"]
	filePath := r.Header.Get("filePath")
	fileName := r.Header.Get("fileName")

	fullFilePath := config.Envs.BucketPath + bucketName + "/" + filePath + "/" + fileName

	file, err := os.Open(fullFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			utils.NewErrorResponse(w, "File not found", http.StatusNotFound)
		} else {
			utils.NewErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	defer file.Close()

	// Get file information to set headers
	fileInfo, err := file.Stat()
	if err != nil {
		utils.NewErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if fileInfo.IsDir() {
		utils.NewErrorResponse(w, "Requested path is a directory, not a file", http.StatusBadRequest)
		return
	}
	// Set headers
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name()))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// Stream the file to the response writer
	buf := make([]byte, config.Envs.BufferSize)
	for {
		n, readErr := file.Read(buf)
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			utils.NewErrorResponse(w, readErr.Error(), http.StatusInternalServerError)
			return
		}

		_, writeErr := w.Write(buf[:n])
		if writeErr != nil {
			utils.NewErrorResponse(w, writeErr.Error(), http.StatusInternalServerError)
			return
		}
	}
}
