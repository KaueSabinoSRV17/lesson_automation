package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvs(path string, envNames ...string) map[string]string {
	envs := make(map[string]string)
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Could not load Env Vars")
		return nil
	}
	for _, envName := range envNames {
		envValue := os.Getenv(envName)
		envs[envName] = envValue
	}
	return envs
}
