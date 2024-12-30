package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const BUCKETS_PATH = "buckets/"
const BUFFER_SIZE = 50

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
		newErrorResponse(w, "name is required", http.StatusBadRequest)
		return
	}
	path, ok := mux.Vars(r)["path"]
	if !ok || path == "" {
		newErrorResponse(w, "path is required", http.StatusBadRequest)
		return
	}

	outFile, err := os.Create(BUCKETS_PATH + path)
	if err != nil {
		newErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	buf := make([]byte, BUFFER_SIZE)
	var totalBytes int64
	for {
		n, readErr := r.Body.Read(buf)
		if readErr == io.EOF {
			fmt.Println("reached end of file at", totalBytes)
			break
		}
		if readErr != nil {
			newErrorResponse(w, readErr.Error(), http.StatusInternalServerError)
			return
		}

		bytesWritten, err := outFile.Write(buf[:n])

		if err != nil {
			newErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
		totalBytes += int64(bytesWritten)
	}

	if totalBytes == 0 {
		newErrorResponse(w, "No data received", http.StatusBadRequest)
		return
	}

	newSuccessResponse(w, fmt.Sprintf("File uploaded successfully to: %s/%s with size %d", name, path, totalBytes))
}

func newErrorResponse(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

func newSuccessResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
