package config

import "os"

type Config struct {
	DvlaApiUrl string
}

func GenerateConfig() Config {
	return Config{
		DvlaApiUrl: os.Getenv("DVLA_API_URL"),
	}
}
