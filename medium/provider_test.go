package medium

import(
	"os"
)

var credsEnvVars = []string{
	"MEDIUM_ACCESS_TOKEN",
}

func getTestCredsFromEnv() string {
	return multiEnvSearch(credsEnvVars)
}

func multiEnvSearch(ks []string) string {
	for _, k := range ks {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return ""
}