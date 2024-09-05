package jsonrpc_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mrdhat/eth-txns/jsonrpc"
)

func TestJSONRPCClient_MakeRequest(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)

		assert.Equal(t, "2.0", req["jsonrpc"])
		assert.Equal(t, "test_method", req["method"])
		assert.Equal(t, []interface{}{"param1", "param2"}, req["params"])

		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"result":  "test_result",
			"id":      req["id"],
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Create client
	client := jsonrpc.NewJSONRPCClient(server.URL, server.Client())

	// Test MakeRequest
	result, err := client.MakeRequest("test_method", []interface{}{"param1", "param2"})
	require.NoError(t, err)
	assert.Equal(t, "test_result", result)
}

func TestJSONRPCClient_MakeRequest_Error(t *testing.T) {
	// Mock server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Create client
	client := jsonrpc.NewJSONRPCClient(server.URL, server.Client())

	// Test MakeRequest with error
	_, err := client.MakeRequest("test_method", []interface{}{"param1", "param2"})
	require.Error(t, err)
}
