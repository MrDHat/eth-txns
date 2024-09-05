package jsonrpc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client interface {
	MakeRequest(method string, params []interface{}) (interface{}, error)
}

type request struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type response struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      int         `json:"id"`
}

type jsonRPCClient struct {
	url    string
	client *http.Client
}

func (j *jsonRPCClient) MakeRequest(method string, params []interface{}) (interface{}, error) {
	req := request{
		JSONRPC: "2.0",
		ID:      1,
		Params:  params,
		Method:  method,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := j.client.Post(j.url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respBody response
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, err
	}

	return respBody.Result, nil
}

func NewJSONRPCClient(url string, client *http.Client) Client {
	return &jsonRPCClient{url: url, client: client}
}
