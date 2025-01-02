package cmd

import "github.com/spf13/cobra"

var apikeyCmd = &cobra.Command{
	Use:   "apikey",
	Short: "apikey command is used to manage api keys",
}

func init() {
	apikeyCmd.AddCommand(apikeyCreateCmd)
	apikeyCmd.AddCommand(apikeyListCmd)
	apikeyCmd.AddCommand(apikeyAttachCmd)
}
