package studeoapi

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAuthenticate(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Error("Could not get Env Vars")
	}
	cpf := os.Getenv("CPF")
	password := os.Getenv("PASSWORD")
	if cpf == "" || password == "" {
		t.Error("Env vars are null")
	}

	token := Authenticate(cpf, password)
	if token == "" {
		t.Error("Token is not valid")
	}
}
