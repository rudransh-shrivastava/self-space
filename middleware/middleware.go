package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rudransh-shrivastava/self-space/api/apikey"
	"github.com/rudransh-shrivastava/self-space/api/apikeybucketpermission"
	"github.com/rudransh-shrivastava/self-space/api/bucket"
	"github.com/rudransh-shrivastava/self-space/utils"
)

// validate api key, bucket name and permissions
// set headers for fileName and filePath
func AuthMiddleware(apiKeyStore *apikey.APIKeyStore, bucketStore *bucket.BucketStore, apiKeyBucketPermissionStore *apikeybucketpermission.APIKeyBucketPermissionStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// validate api key
			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				utils.NewErrorResponse(w, "API Key is required", http.StatusUnauthorized)
				return
			}
			dbApiKey, err := apiKeyStore.FindAPIKeyByKey(apiKey)
			if err != nil {
				fmt.Println(err)
				utils.NewErrorResponse(w, "API Key is invalid", http.StatusUnauthorized)
				return
			}
			// validate bucket name
			bucketName, ok := mux.Vars(r)["bucketName"]
			if !ok || bucketName == "" {
				utils.NewErrorResponse(w, "bucket name is required", http.StatusBadRequest)
				return
			}
			dbBucket, err := bucketStore.FindBucketByName(bucketName)
			if err != nil {
				utils.NewErrorResponse(w, "bucket does not exist", http.StatusNotFound)
				return
			}
			// validate permissions
			var permission string
			switch r.Method {
			case http.MethodGet:
				permission = "READ"
			case http.MethodPut:
				permission = "WRITE"
			case http.MethodDelete:
				permission = "DELETE"
			default:
				utils.NewErrorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
			}

			hasPermission := false

			hasPermission, err = apiKeyBucketPermissionStore.HasPermission(dbApiKey, dbBucket, permission)

			if err != nil {
				utils.NewErrorResponse(w, "error while checking permissions", http.StatusInternalServerError)
				return
			}
			if !hasPermission {
				utils.NewErrorResponse(w, "permission denied", http.StatusForbidden)
			}

			// set headers for fileName and filePath
			urlFullPath, ok := mux.Vars(r)["filePath"]
			if !ok || urlFullPath == "" {
				utils.NewErrorResponse(w, "file path is required", http.StatusBadRequest)
				return
			}

			pathArray := strings.Split(urlFullPath, "/")
			fileName := pathArray[len(pathArray)-1]
			filePath := strings.Join(pathArray[:len(pathArray)-1], "/")

			// set headers
			r.Header.Set("fileName", fileName)
			r.Header.Set("filePath", filePath)

			next.ServeHTTP(w, r)
		})
	}
}
