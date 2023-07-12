package studeoapi

// Authenticate Returns a JWT token to be used in the Authorization Header for the next HTTP Calls
func Authenticate(cpf, password string) string {
	api := Instance
	requestBody := map[string]string{
		"username": cpf,
		"password": password,
	}
	var response AuthResponse
	api.Post("/auth-api-controller/auth/token/create-ldap", requestBody, &response)
	return response.Token
}
