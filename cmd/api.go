package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ApiCmd)
}

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "some short usage goes here",
	Long:  "longer usage goes here",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("foo command called!")
	},
}
