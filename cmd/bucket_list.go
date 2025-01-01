package cmd

import (
	"fmt"

	api "github.com/rudransh-shrivastava/self-space/api/bucket"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/spf13/cobra"
)

var bucketListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all buckets",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.NewDB()
		bucketStore := api.NewBucketStore(db)
		buckets, err := bucketStore.ListBuckets()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ID    NAME")
		for _, bucket := range buckets {
			fmt.Printf("%d     %s\n", bucket.ID, bucket.Name)
		}
	},
}
