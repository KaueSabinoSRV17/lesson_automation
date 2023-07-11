package auth

import (
	"os"
	"regexp"
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

	token, err := Authenticate(cpf, password)
	if err != nil || token == "" {
		t.Error("Could not authenticate")
	}

	regexForToken := regexp.MustCompile(`^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_.+/=]+$`)
	isAValidToken := !regexForToken.MatchString(token)
	if !isAValidToken {
		t.Error("The token is not valid!")
	}
}
