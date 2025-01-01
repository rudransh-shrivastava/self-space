package cmd

import (
	"fmt"

	api "github.com/rudransh-shrivastava/self-space/api/bucket"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/spf13/cobra"
)

var bucketName string

var bucketCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.NewDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		bucketStore := api.NewBucketStore(db)
		err = bucketStore.CreateBucket(bucketName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("created bucket with name: %s\n", bucketName)
	},
}

func init() {
	bucketCreateCmd.Flags().StringVarP(&bucketName, "name", "n", "", "name of the bucket")
	if err := bucketCreateCmd.MarkFlagRequired("name"); err != nil {
		panic(err)
	}
}
