package medium

import (
	"testing"
)

func TestAccConfigLoadValidate_accessToken(t *testing.T) {
	config := Config{}
	err := config.LoadAndValidate()
	if err != nil {
		t.Fatalf("unable to access api: %s", err)
	}
}