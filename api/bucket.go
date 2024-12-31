package api

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
}

// TODO: path split and parse
func (b *Bucket) Upload(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)["name"]
	if !ok || name == "" {
		utils.NewErrorResponse(w, "name is required", http.StatusBadRequest)
		return
	}
	path, ok := mux.Vars(r)["path"]
	if !ok || path == "" {
		utils.NewErrorResponse(w, "path is required", http.StatusBadRequest)
		return
	}

	outFile, err := os.Create(config.Envs.BucketPath + path)
	if err != nil {
		utils.NewErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	buf := make([]byte, config.Envs.BufferSize)
	var totalBytes int64
	for {
		n, readErr := r.Body.Read(buf)
		fmt.Printf("the buffer is %v \n", buf)
		if readErr == io.EOF {
			bytesWritten, err := outFile.Write(buf[:n])
			if err != nil {
				utils.NewErrorResponse(w, err.Error(), http.StatusInternalServerError)
			}
			totalBytes += int64(bytesWritten)
			fmt.Println("reached end of file at", totalBytes)
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

	utils.NewSuccessResponse(w, fmt.Sprintf("File uploaded successfully to: %s/%s with size %d", name, path, totalBytes))
}
