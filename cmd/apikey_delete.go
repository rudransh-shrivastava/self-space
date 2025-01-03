package cmd

import (
	"fmt"

	"github.com/rudransh-shrivastava/self-space/api/apikey"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/spf13/cobra"
)

var apikeyDeleteCmd = &cobra.Command{
	Use:   "delete <api-key>",
	Short: "delete an api key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := args[0]
		db, err := db.NewDB()

		apiKeyStore := apikey.NewAPIKeyStore(db)
		err = apiKeyStore.DeleteAPIKeyByKey(apiKey)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("deleted api key %s", apiKey)
	},
}
