package cmd

import (
	"fmt"

	api "github.com/rudransh-shrivastava/self-space/api/api_key"
	"github.com/rudransh-shrivastava/self-space/db"
	"github.com/rudransh-shrivastava/self-space/utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var apikeyCreateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate a new api key",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := db.NewDB()
		if err != nil {
			fmt.Println(err)
			return
		}
		apiKeyStore := api.NewAPIKeyStore(db)
		apikey, err := utils.GenerateAPIKey()
		if err != nil {
			fmt.Println(err)
			return
		}
		hashedApiKey, err := bcrypt.GenerateFromPassword([]byte(apikey), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = apiKeyStore.CreateAPIKey(string(hashedApiKey))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("generated api key: %s\n", apikey)
	},
}
