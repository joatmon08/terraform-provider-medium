package medium

import (
	"testing"
)

func TestAccConfigLoadValidate_accessToken(t *testing.T) {
	creds := getTestCredsFromEnv()

	config := &Config{
		AccessToken: creds,
	}

	config.LoadAndValidate()

	_, err := config.Client.GetUser("")
	if err != nil {
		t.Fatalf("unable to access api: %s", err)
	}
}