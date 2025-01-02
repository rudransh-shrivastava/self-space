package utils

import (
	"testing"
)

func TestGenerateAPIKey(t *testing.T) {
	for i := 0; i < 5; i++ {
		apiKey, err := GenerateAPIKey()
		if err != nil {
			t.Errorf("error generating api key: %v", err)
		}
		if apiKey == "" {
			t.Errorf("api key empty")
		}
		expectedLength := 44
		if len(apiKey) != expectedLength {
			t.Errorf("api key not of expected length, apikey: %v has length %d", apiKey, len(apiKey))
		}
	}
}
