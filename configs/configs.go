package configs

import "os"

var (
	Domain string
	Port   string
)

const (
	defaultDomain = "localhost"
	defaultPort   = "8080"
)

func InitConfigs() {
	Domain = lookUpOrDefault("DOMAIN", defaultDomain)
	Port = lookUpOrDefault("PORT", defaultPort)
}

func lookUpOrDefault(key, defaultValue string) string {
	out, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return out
}
