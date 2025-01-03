package cmd

import (
	"fmt"
	"strings"

	"github.com/rudransh-shrivastava/self-space/api/apikey"
	"github.com/rudransh-shrivastava/self-space/api/apikeybucketpermission"
	"github.com/rudransh-shrivastava/self-space/api/bucket"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/spf13/cobra"
)

var apikeyDetachCmd = &cobra.Command{
	Use:   "detach <api-key> <bucket-name> <permission>",
	Short: "detach a bucket from an api key",
	Long: `detach a bucket from an api key completely or a given permission
	Usage: self-space apikey detach <api-key> <bucket-name> <permission>
	Permission can be one of (READ, WRITE, DELETE)`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := args[0]
		bucketName := args[1]
		permission := args[2]
		db, err := db.NewDB()

		// validate permission <this is hardcoded i dont know how enums work in go>
		permission = strings.ToUpper(permission)
		if permission != "READ" && permission != "WRITE" && permission != "DELETE" {
			fmt.Println("permission can only be one of (READ, WRITE, DELETE)")
			return
		}
		// validate api key
		apiKeyStore := apikey.NewAPIKeyStore(db)
		dbApiKey, err := apiKeyStore.FindAPIKeyByKey(apiKey)
		if err != nil {
			fmt.Println("api key does not exist")
			return
		}
		// validate bucket
		bucketStore := bucket.NewBucketStore(db)
		dbBucket, err := bucketStore.FindBucketByName(bucketName)
		if err != nil {
			fmt.Println("bucket does not exist")
			return
		}
		// attach bucket to api key
		apiKeyBucketPermissionStore := apikeybucketpermission.NewAPIKeyBucketPermissionStore(db)
		err = apiKeyBucketPermissionStore.DeletePermission(dbApiKey, dbBucket, permission)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("detached bucket %s from api key %s with permission %s\n", bucketName, apiKey, permission)
	},
}
