package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rudransh-shrivastava/self-space/utils"
)

const BUCKETS_PATH = "buckets/"
const BUFFER_SIZE = 9

type Bucket struct {
}

func main() {
	// create buckets folder
	if _, err := os.Stat(BUCKETS_PATH); os.IsNotExist(err) {
		os.Mkdir(BUCKETS_PATH, 0755)
	}
	r := mux.NewRouter()
	bucket := &Bucket{}
	r.HandleFunc("/bucket/{name}/{path:.*}", bucket.upload).Methods("PUT")
	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", r)

}

// TODO: path split and parse
func (b *Bucket) upload(w http.ResponseWriter, r *http.Request) {
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

	outFile, err := os.Create(BUCKETS_PATH + path)
	if err != nil {
		utils.NewErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	buf := make([]byte, BUFFER_SIZE)
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
