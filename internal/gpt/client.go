package gpt

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var Model = "gpt-3.5-turbo"
var BaseUrl = "https://api.openai.com/v1"

type ChatMember struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAiApi struct {
	ApiKey      string
	ChatMembers []ChatMember
}

func DefaultClient(apiKey string) *OpenAiApi {
	return &OpenAiApi{
		ApiKey: apiKey,
		ChatMembers: []ChatMember{
			{Role: "system", Content: "You are a helpful assistant"},
			{Role: "user", Content: "Hello!"},
		},
	}
}

func (o *OpenAiApi) Post(path string, input any) (response []byte, err error) {
	requestJson, err := json.Marshal(input)
	if err != nil {
		return
	}

	httpResponse, err := o.Do(http.MethodPost, path, bytes.NewReader(requestJson))
	if err != nil {
		return
	}
	defer httpResponse.Body.Close()

	response, err = io.ReadAll(httpResponse.Body)
	return
}

func (o *OpenAiApi) Do(method, path string, body io.Reader) (response *http.Response, err error) {
	url := BaseUrl + path
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	request.Header.Add("Authorization", "Bearer "+o.ApiKey)
	request.Header.Add("Content-Type", "application/json")
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	return
}
