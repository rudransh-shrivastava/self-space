package cmd

import (
	"fmt"
	"os"

	"github.com/rudransh-shrivastava/self-space/app"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "self-space",
	Short: "self-space - is a ",
	Long: `self-space is a >>>>>>>>
	<<<<<
	>>>>>
	<<<<<`,
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
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing'%s'", err)
		os.Exit(1)
	}
}
