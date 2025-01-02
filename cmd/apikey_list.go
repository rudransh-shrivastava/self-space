package cmd

import (
	"fmt"

	api "github.com/rudransh-shrivastava/self-space/api/api_key"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/spf13/cobra"
)

var apikeyListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all api keys",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.NewDB()
		apiKeyStore := api.NewAPIKeyStore(db)
		apikeys, err := apiKeyStore.ListAPIKeys()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ID    Key")
		for _, key := range apikeys {
			fmt.Printf("%d     %s\n", key.ID, key.Key)
		}
	},
}