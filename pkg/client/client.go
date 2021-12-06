package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

func New(baseUrl string) Client {
	return &client{BaseUrl: baseUrl}
}

func (h client) Get(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", h.BaseUrl, endpoint), bytes.NewBuffer([]byte{}))
}

func (h client) GetWith(endpoint string, params interface{}) (*http.Request, error) {
	queryString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?%s", h.BaseUrl, endpoint, queryString.Encode()), bytes.NewBuffer([]byte{}))
}

func (h client) Post(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, h.BaseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

func (h client) PostWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

func (h client) Put(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

func (h client) PutWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

func (h client) Patch(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodPatch, h.BaseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

func (h client) PatchWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

func (h client) Delete(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodDelete, h.BaseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

func (h client) DeleteWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, h.BaseUrl+endpoint, bytes.NewBuffer(json))
}

func (h client) Do(request *http.Request) (Response, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &ResponseStruct{
		Status:        response.Status,
		StatusCode:    response.StatusCode,
		Header:        response.Header,
		ContentLength: response.ContentLength,
		Body:          body,
	}, nil
}

func (r ResponseStruct) Get() ResponseStruct {
	return r
}

func (r ResponseStruct) To(value interface{}) {
	fmt.Println(r.Body)
	err := json.Unmarshal(r.Body, &value)
	if err != nil {
		value = nil
	}
}
