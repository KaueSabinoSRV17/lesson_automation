package env

import (
	"testing"
)

func TestGetEnvs(t *testing.T) {
	envNames := []string{"CPF", "PASSWORD", "OPENAI_API_KEY"}
	envs := GetEnvs("../../../.env", envNames...)
	for _, envName := range envNames {
		if envs[envName] == "" {
			t.Error("Should not be empty")
		}
	}
}
