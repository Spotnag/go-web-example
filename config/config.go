package config

import (
	"log"
	"os"
)

type Config struct {
	CloudflareAPIToken  string
	CloudflareAccountID string
	MailgunAPIKey       string
	MailgunDomain       string
}

var AppConfig *Config

func InitConfig() {
	AppConfig = &Config{
		CloudflareAPIToken:  getEnv("CLOUDFLARE_API_TOKEN"),
		CloudflareAccountID: getEnv("CLOUDFLARE_ACCOUNT_ID"),
		MailgunAPIKey:       getEnv("MAILGUN_API_KEY"),
		MailgunDomain:       getEnv("MAILGUN_DOMAIN"),
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("missing required environment variable %s", key)
	}
	return value
}
