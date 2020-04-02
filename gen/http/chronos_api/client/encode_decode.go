// Code generated by goa v3.1.1, DO NOT EDIT.
//
// ChronosAPI HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/InjectiveLabs/injective-core/api/design -o ../

package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	chronosapi "github.com/InjectiveLabs/dexterm/gen/chronos_api"
	goahttp "goa.design/goa/v3/http"
)

// BuildSymbolInfoRequest instantiates a HTTP request object with method and
// path set to call the "ChronosAPI" service "symbolInfo" endpoint
func (c *Client) BuildSymbolInfoRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SymbolInfoChronosAPIPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("ChronosAPI", "symbolInfo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSymbolInfoRequest returns an encoder for requests sent to the
// ChronosAPI symbolInfo server.
func EncodeSymbolInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chronosapi.SymbolInfoPayload)
		if !ok {
			return goahttp.ErrInvalidType("ChronosAPI", "symbolInfo", "*chronosapi.SymbolInfoPayload", v)
		}
		values := req.URL.Query()
		if p.Group != nil {
			values.Add("group", *p.Group)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeSymbolInfoResponse returns a decoder for responses returned by the
// ChronosAPI symbolInfo endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeSymbolInfoResponse may return the following errors:
//	- "bad_request" (type *chronosapi.BaseChronosResponse): http.StatusBadRequest
//	- "not_found" (type *chronosapi.BaseChronosResponse): http.StatusNotFound
//	- "internal" (type *chronosapi.BaseChronosResponse): http.StatusInternalServerError
//	- error: internal error
func DecodeSymbolInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SymbolInfoResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ChronosAPI", "symbolInfo", err)
			}
			err = ValidateSymbolInfoResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("ChronosAPI", "symbolInfo", err)
			}
			res := NewSymbolInfoTradingViewSymbolInfoResponseOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body SymbolInfoBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ChronosAPI", "symbolInfo", err)
			}
			err = ValidateSymbolInfoBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("ChronosAPI", "symbolInfo", err)
			}
			return nil, NewSymbolInfoBadRequest(&body)
		case http.StatusNotFound:
			var (
				body SymbolInfoNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ChronosAPI", "symbolInfo", err)
			}
			err = ValidateSymbolInfoNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("ChronosAPI", "symbolInfo", err)
			}
			return nil, NewSymbolInfoNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body SymbolInfoInternalResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ChronosAPI", "symbolInfo", err)
			}
			err = ValidateSymbolInfoInternalResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("ChronosAPI", "symbolInfo", err)
			}
			return nil, NewSymbolInfoInternal(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("ChronosAPI", "symbolInfo", resp.StatusCode, string(body))
		}
	}
}

// BuildHistoryRequest instantiates a HTTP request object with method and path
// set to call the "ChronosAPI" service "history" endpoint
func (c *Client) BuildHistoryRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: HistoryChronosAPIPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("ChronosAPI", "history", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeHistoryRequest returns an encoder for requests sent to the ChronosAPI
// history server.
func EncodeHistoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chronosapi.HistoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("ChronosAPI", "history", "*chronosapi.HistoryPayload", v)
		}
		values := req.URL.Query()
		values.Add("symbol", p.Symbol)
		values.Add("resolution", p.Resolution)
		values.Add("from", fmt.Sprintf("%v", p.From))
		values.Add("to", fmt.Sprintf("%v", p.To))
		values.Add("countback", fmt.Sprintf("%v", p.Countback))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeHistoryResponse returns a decoder for responses returned by the
// ChronosAPI history endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeHistoryResponse may return the following errors:
//	- "bad_request" (type *chronosapi.BaseChronosResponse): http.StatusBadRequest
//	- "not_found" (type *chronosapi.BaseChronosResponse): http.StatusNotFound
//	- "internal" (type *chronosapi.BaseChronosResponse): http.StatusInternalServerError
//	- error: internal error
func DecodeHistoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body HistoryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ChronosAPI", "history", err)
			}
			err = ValidateHistoryResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("ChronosAPI", "history", err)
			}
			res := NewHistoryResponseOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body HistoryBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ChronosAPI", "history", err)
			}
			err = ValidateHistoryBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("ChronosAPI", "history", err)
			}
			return nil, NewHistoryBadRequest(&body)
		case http.StatusNotFound:
			var (
				body HistoryNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ChronosAPI", "history", err)
			}
			err = ValidateHistoryNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("ChronosAPI", "history", err)
			}
			return nil, NewHistoryNotFound(&body)
		case http.StatusInternalServerError:
			var (
				body HistoryInternalResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ChronosAPI", "history", err)
			}
			err = ValidateHistoryInternalResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("ChronosAPI", "history", err)
			}
			return nil, NewHistoryInternal(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("ChronosAPI", "history", resp.StatusCode, string(body))
		}
	}
}