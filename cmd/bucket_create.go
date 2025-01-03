package cmd

import (
	"fmt"

	api "github.com/rudransh-shrivastava/self-space/api/bucket"
	"github.com/rudransh-shrivastava/self-space/config"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/spf13/cobra"
)

var bucketCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a bucket",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bucketName := args[0]
		if len(bucketName) > config.Envs.BucketNameMaxLength {
			fmt.Println("bucket name cannot be more than 100 characters")
		}

		db, err := db.NewDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		bucketStore := api.NewBucketStore(db)
		err = bucketStore.CreateBucket(bucketName)
		if err != nil {
			return
		}
		fmt.Printf("created bucket with name: %s\n", bucketName)
	},
}
