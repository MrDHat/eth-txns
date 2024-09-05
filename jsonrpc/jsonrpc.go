package jsonrpc

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mrdhat/eth-txns/logger"
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
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
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

	if respBody.Error.Code != 0 {
		// For now, we chose not to bubble up errors for the sake of simplicity
		logger.Log("JSONRPC ERROR: ", respBody.Error.Message, " for ", method, params)
		return nil, nil
	}

	return respBody.Result, nil
}

func NewJSONRPCClient(url string, client *http.Client) Client {
	return &jsonRPCClient{url: url, client: client}
}
