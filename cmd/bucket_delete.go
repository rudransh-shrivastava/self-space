package cmd

import (
	"fmt"

	api "github.com/rudransh-shrivastava/self-space/api/bucket"
	"github.com/rudransh-shrivastava/self-space/config"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/rudransh-shrivastava/self-space/utils"
	"github.com/spf13/cobra"
)

var bucketDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a bucket",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bucketName := args[0]
		if len(bucketName) > config.Envs.BucketNameMaxLength {
			fmt.Println("bucket does not exist")
		}

		db, err := db.NewDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		bucketStore := api.NewBucketStore(db)
		err = bucketStore.DeleteBucket(bucketName)
		if err != nil {
			fmt.Println(err)
			return
		}
		bucketPath := config.Envs.BucketPath + bucketName
		err = utils.DeleteDirectory(bucketPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("deleted bucket with name: %s\n", bucketName)
	},
}
