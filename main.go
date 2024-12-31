package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rudransh-shrivastava/self-space/api"
	"github.com/rudransh-shrivastava/self-space/config"
)

func main() {
	if _, err := os.Stat(config.Envs.BucketPath); os.IsNotExist(err) {
		os.Mkdir(config.Envs.BucketPath, 0755)
	}
	r := mux.NewRouter()
	bucket := api.Bucket{}
	r.HandleFunc("/bucket/{name}/{path:.*}", bucket.Upload).Methods("PUT")
	addr := fmt.Sprintf("%s:%s", config.Envs.PublicHost, config.Envs.Port)
	fmt.Printf("listening on %s", addr)
	http.ListenAndServe(addr, r)
}
