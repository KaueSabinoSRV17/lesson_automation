package auth

import (
	"github.com/KaueSabinoSRV17/lesson_automation/internal/studeoapi"
)

// Authenticate Returns a JWT token to be used in the Authorization Header for the next HTTP Calls
func Authenticate(cpf, password string) string {
	api := studeoapi.Instance
	requestBody := map[string]string{
		"username": cpf,
		"password": password,
	}
	var token string
	api.Post("/auth-api-controller/auth/token/create-ldap", requestBody, &token)
	return string(token)
}
