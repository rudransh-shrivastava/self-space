package middleware

import (
	"fmt"
	"net/http"

	"github.com/rudransh-shrivastava/self-space/api/apikey"
	"github.com/rudransh-shrivastava/self-space/utils"
)

func ApiKeyMiddleware(apiKeyStore *apikey.APIKeyStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				utils.NewErrorResponse(w, "API Key is required", http.StatusUnauthorized)
				return
			}
			_, err := apiKeyStore.FindAPIKeyByKey(apiKey)
			if err != nil {
				fmt.Println(err)
				utils.NewErrorResponse(w, "API Key is invalid", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
