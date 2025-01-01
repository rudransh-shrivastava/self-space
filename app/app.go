package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/rudransh-shrivastava/self-space/api/bucket"
	"github.com/rudransh-shrivastava/self-space/config"
	"github.com/rudransh-shrivastava/self-space/utils"
	"gorm.io/gorm"
)

type ApiServer struct {
	db *gorm.DB
}

func NewApiServer(db *gorm.DB) *ApiServer {
	return &ApiServer{db: db}
}

func (a *ApiServer) Start() {
	// create buckets storage directory
	utils.CreateDirectoryIfNotExists(config.Envs.BucketPath)

	r := mux.NewRouter()
	bucketStore := api.NewBucketStore(a.db)
	bucket := api.Bucket{BucketStore: bucketStore}

	r.HandleFunc("/bucket/{bucketName}/{filePath:.*}", bucket.Upload).Methods("PUT")
	addr := fmt.Sprintf("%s:%s", config.Envs.PublicHost, config.Envs.Port)

	fmt.Printf("listening on %s \n", addr)
	http.ListenAndServe(addr, r)
}
