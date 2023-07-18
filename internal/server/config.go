package server

import (
	"os"
	"strconv"
)

type Config struct {
	ListenAddr     string
	Difficulty     byte
	ProofTokenSize int
}

func NewConfig() *Config {
	c := &Config{
		ListenAddr:     envOrDefault("LISTEN_ADDR", "0.0.0.0:9000"),
		Difficulty:     byte(envOrDefaultInt("DIFFICULTY", 23)),
		ProofTokenSize: envOrDefaultInt("PROOF_TOKEN_SIZE", 64),
	}

	return c
}

func envOrDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func envOrDefaultInt(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}
	return defaultValue
}
