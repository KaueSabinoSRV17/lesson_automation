package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/KaueSabinoSRV17/lesson_automation/internal/studeoapi"
)

func Authenticate(cpf, password string) (string, error) {
	url := fmt.Sprint(studeoapi.Url + "/auth-api-controller/auth/token/create-ldap")
	requestBody := map[string]string{
		"username": cpf,
		"password": password,
	}

	json, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal("Could not create a JSON Body")
		return "", err
	}

	response, err := http.Post(url, studeoapi.JsonContentType, bytes.NewBuffer(json))
	if err != nil {
		log.Fatal("Could not Authenticate:\n\t" + err.Error())
		return "", err
	}
	defer response.Body.Close()

	token, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Could not read response body")
		return "", err
	}

	return string(token), nil
}
