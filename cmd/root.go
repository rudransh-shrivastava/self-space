package cmd

import (
	"fmt"
	"os"

	"github.com/rudransh-shrivastava/self-space/app"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "self-space",
	Long: `self-space is a self hosted cloud storage service`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.NewDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		apiServer := app.NewApiServer(db)
		apiServer.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing'%s'", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(bucketCmd)
	rootCmd.AddCommand(apikeyCmd)
}
