package studeoapi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	BaseUrl    string
	Token      string
	httpClient *http.Client
}

var BaseUrl = "https://studeoapi.unicesumar.edu.br"
var JsonContentType = "application/json"
var Instance = Client{
	BaseUrl:    BaseUrl,
	httpClient: http.DefaultClient,
}

func (c *Client) Get(path string, body, resultReference interface{}) (string, error) {
	request, err := c.newRequest(http.MethodGet, path, body)
	if err != nil {
		log.Fatalln("Could not create Studeo API Http Request\n", err)
	}
	response, err := c.do(request, resultReference)
	if err != nil {
		log.Fatal("Http Error\n", err)
		return "", err
	}
	return response, nil
}

func (c *Client) Post(path string, body, resultReference interface{}) (string, error) {
	request, err := c.newRequest(http.MethodPost, path, body)
	if err != nil {
		log.Fatal("Could not create Studeo API Http Request", err)
	}
	response, err := c.do(request, resultReference)
	if err != nil {
		log.Fatal("Http Error", err)
		return "", err
	}
	return response, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	url := c.BaseUrl + path

	var buffer io.ReadWriter
	if body != nil {
		buffer = new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}
	if body != err {
		request.Header.Set("Content-Type", JsonContentType)
	}
	if c.Token != "" {
		request.Header.Set("Authorization", c.Token)
	}
	request.Header.Set("Accept", JsonContentType)

	return request, nil
}

func (c *Client) do(request *http.Request, resultReference interface{}) (string, error) {
	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}
