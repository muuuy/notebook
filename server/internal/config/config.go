package config

type Config struct {
	ServerPort  string
	DatabaseDSN string
}

func LoadConfig() (*Config, error) {
	// Mock loading config; replace with actual logic
	return &Config{
		ServerPort:  "8080",
		DatabaseDSN: "postgres://user:password@localhost:5432/db",
	}, nil
}
