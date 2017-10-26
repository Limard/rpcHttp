// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonrpc2

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
)

// ----------------------------------------------------------------------------
// Request and Response
// ----------------------------------------------------------------------------

// clientRequest represents a JSON-RPC request sent by a client.
type clientRequest struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`

	// A String containing the name of the method to be invoked.
	Method string `json:"method"`

	// Object to pass as request parameter to the method.
	Params interface{} `json:"params"`

	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	Id uint64 `json:"id"`
}

// clientResponse represents a JSON-RPC response returned to a client.
type clientResponse struct {
	Id      interface{}      `json:"id"`
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
}

// EncodeClientRequest encodes parameters for a JSON-RPC client request.
func EncodeClientRequest(method string, args interface{}) ([]byte, error) {
	c := &clientRequest{
		Version: "2.0",
		Method:  method,
		Params:  args,
		Id:      uint64(rand.Int63()),
	}
	return json.Marshal(c)
}

// DecodeClientResponse decodes the response body of a client request into
// the interface reply.
<<<<<<< HEAD
func DecodeClientResponse(r io.Reader, reply interface{}) (replyError *Error) {
	var c clientResponse
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		replyError = &Error{
			Code:    E_PARSE,
			Message: err.Error()}
		return
=======
func DecodeClientResponse(r io.Reader, reply interface{}) (replyError *Error, err error) {
	var c clientResponse
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return nil, err
>>>>>>> 08d48d6aa4f5f44f8ad58d1495db7936f983dbe7
	}

	// Error
	if c.Error != nil {
<<<<<<< HEAD
		replyError = &Error{}
		if err := json.Unmarshal(*c.Error, replyError); err != nil {
			replyError = &Error{
				Code:    E_PARSE,
				Message: string(*c.Error),
			}
		}
		return replyError
=======
		jsonErr := &Error{}
		if err := json.Unmarshal(*c.Error, jsonErr); err != nil {
			return nil, &Error{
				Code:    E_SERVER,
				Message: string(*c.Error),
			}
		}
		return jsonErr, nil
>>>>>>> 08d48d6aa4f5f44f8ad58d1495db7936f983dbe7
	}

	// Result
	if c.Result == nil {
<<<<<<< HEAD
		replyError = &Error{
			Code:    E_BAD_PARAMS,
			Message: ErrNullResult.Error(),
		}
		return
	}
	if err := json.Unmarshal(*c.Result, reply); err != nil {
		replyError = &Error{
			Code:    E_PARSE,
			Message: err.Error(),
		}
		return
	}
	return
}

func Call(url string, method string, request interface{}, reply interface{}) (replyError *Error) {
	jsonReqBuf, err := EncodeClientRequest(method, request)
	if err != nil {
		replyError = &Error{
			Code:    E_INVALID_REQ,
			Message: err.Error()}
=======
		return nil, ErrNullResult
	}
	return nil, json.Unmarshal(*c.Result, reply)
}

func Call(url string, method string, request interface{}, reply interface{}) (replyError *Error, err error) {
	jsonReqBuf, err := EncodeClientRequest(method, request)
	if err != nil {
>>>>>>> 08d48d6aa4f5f44f8ad58d1495db7936f983dbe7
		return
	}

	jsonReqBufR := bytes.NewReader(jsonReqBuf)
	rsp, err := http.Post(url, `application/json`, jsonReqBufR)
	if err != nil {
<<<<<<< HEAD
		replyError = &Error{
			Code:    E_SERVER,
			Message: err.Error()}
=======
>>>>>>> 08d48d6aa4f5f44f8ad58d1495db7936f983dbe7
		return
	}
	defer rsp.Body.Close()

	return DecodeClientResponse(rsp.Body, reply)
}