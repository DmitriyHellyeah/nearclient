package jsonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	jsonrpcVersion = "2.0"
)

type rpcClient struct {
	host string

	httpClient    *http.Client
	customHeaders map[string]string
}

type RPCClientOpts struct {
	HTTPClient    *http.Client
	CustomHeaders map[string]string
}

type RPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      int         `json:"id"`
}

type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      int         `json:"id"`
}

func (RPCResponse *RPCResponse) Decode(to interface{}) error {
	res, err := json.Marshal(RPCResponse.Result)
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, to)
	if err != nil {
		return err
	}
	return nil
}

type RPCError struct {
	Name  string `json:"name"`
	Cause struct {
		Name string `json:"name"`
		Info struct {
			ErrorMessage string `json:"error_message,omitempty"`
			BlockReference interface{} `json:"block_reference,omitempty"`
		} `json:"info"`
	} `json:"cause,omitempty"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s : %s \n Data: %v ", strconv.Itoa(e.Code), e.Message, e.Data)
}

type HTTPError struct {
	Code int
	err  error
}

func (e *HTTPError) Error() string {
	return e.err.Error()
}

type RPCClient interface {
	Call(method string, params interface{}) (*RPCResponse, error)
	CallRaw(request *RPCRequest) (*RPCResponse, error)
	CallWithDecode(method string, params, out interface{}) error
}

func NewRPCClient(host string) RPCClient {
	return NewClientWithOpts(host, nil)
}

func NewClientWithOpts(host string, opts *RPCClientOpts) RPCClient {
	rpcClient := &rpcClient{
		host:          host,
		httpClient:    &http.Client{},
		customHeaders: make(map[string]string),
	}

	if opts == nil {
		return rpcClient
	}

	if opts.HTTPClient != nil {
		rpcClient.httpClient = opts.HTTPClient
	}

	if opts.CustomHeaders != nil {
		for k, v := range opts.CustomHeaders {
			rpcClient.customHeaders[k] = v
		}
	}

	return rpcClient
}

func (client *rpcClient) Call(method string, params interface{}) (*RPCResponse, error) {
	request := &RPCRequest{
		Method:  method,
		Params:  params,
		JSONRPC: jsonrpcVersion,
	}
	return client.doCall(request)
}

func (client *rpcClient) CallRaw(request *RPCRequest) (*RPCResponse, error) {
	return client.doCall(request)
}

func (client *rpcClient) CallWithDecode(method string, params, out interface{}) error {
	rpcResponse, err := client.Call(method, params)

	if err != nil {
		return err
	}

	if rpcResponse.Error != nil {
		return rpcResponse.Error
	}
	return rpcResponse.Decode(out)
}

func (client *rpcClient) newRequest(req interface{}) (*http.Request, error) {
	body, err := json.Marshal(req)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", client.host, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	for k, v := range client.customHeaders {
		request.Header.Set(k, v)
	}

	return request, nil
}

func (client *rpcClient) doCall(RPCRequest *RPCRequest) (*RPCResponse, error) {
	httpRequest, err := client.newRequest(RPCRequest)
	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v: %v", RPCRequest.Method, client.host, err.Error())
	}

	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("rpc call %v() on %v: %v", RPCRequest.Method, httpRequest.URL.String(), err.Error())
	}
	defer func() { _ = httpResponse.Body.Close() }()

	var rpcResponse *RPCResponse
	decoder := json.NewDecoder(httpResponse.Body)
	decoder.DisallowUnknownFields()
	decoder.UseNumber()

	err = decoder.Decode(&rpcResponse)

	if err != nil {
		if httpResponse.StatusCode >= 400 {
			return nil, &HTTPError{
				Code: httpResponse.StatusCode,
				err:  fmt.Errorf("rpc call %v() on %v status code: %v. could not decode body to rpc response: %v", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode, err.Error()),
			}
		}
		return nil, fmt.Errorf("rpc call %v() on %v status code: %v. could not decode body to rpc response: %v", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode, err.Error())
	}

	if rpcResponse == nil {
		// if we have some http error, return it
		if httpResponse.StatusCode >= 400 {
			return nil, &HTTPError{
				Code: httpResponse.StatusCode,
				err:  fmt.Errorf("rpc call %v() on %v status code: %v. rpc response missing", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode),
			}
		}
		return nil, fmt.Errorf("rpc call %v() on %v status code: %v. rpc response missing", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode)
	}

	if rpcResponse.Error != nil {
		return nil, &HTTPError{
			Code: httpResponse.StatusCode,
			err:  fmt.Errorf("rpc call %v() on %v status code: %v. rpc response is not valid. %v", RPCRequest.Method, httpRequest.URL.String(), httpResponse.StatusCode, rpcResponse.Error),
		}
	}
	return rpcResponse, nil
}
