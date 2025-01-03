package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rudransh-shrivastava/self-space/api/apikey"
	"github.com/rudransh-shrivastava/self-space/api/apikeybucketpermission"
	api "github.com/rudransh-shrivastava/self-space/api/bucket"
	"github.com/rudransh-shrivastava/self-space/config"
	"github.com/rudransh-shrivastava/self-space/middleware"
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

	apiKeyStore := apikey.NewAPIKeyStore(a.db)
	apiKeyBucketPermissionStore := apikeybucketpermission.NewAPIKeyBucketPermissionStore(a.db)

	// use middleware
	r.Use(middleware.AuthMiddleware(apiKeyStore, bucketStore, apiKeyBucketPermissionStore))

	r.HandleFunc("/bucket/{bucketName}/{filePath:.*}", bucket.Upload).Methods("PUT")
	r.HandleFunc("/bucket/{bucketName}/{filePath:.*}", bucket.Download).Methods("GET")
	r.HandleFunc("/bucket/{bucketName}/{filePath:.*}", bucket.Delete).Methods("DELETE")

	addr := fmt.Sprintf("%s:%s", config.Envs.PublicHost, config.Envs.Port)

	fmt.Printf("listening on %s \n", addr)
	http.ListenAndServe(addr, r)
}
