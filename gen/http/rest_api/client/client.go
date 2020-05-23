// Code generated by goa v3.1.1, DO NOT EDIT.
//
// RestAPI client HTTP transport
//
// Command:
// $ goa gen github.com/InjectiveLabs/injective-core/api/design -o ../

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the RestAPI service endpoint HTTP clients.
type Client struct {
	// GetActiveOrder Doer is the HTTP client used to make requests to the
	// getActiveOrder endpoint.
	GetActiveOrderDoer goahttp.Doer

	// GetArchiveOrder Doer is the HTTP client used to make requests to the
	// getArchiveOrder endpoint.
	GetArchiveOrderDoer goahttp.Doer

	// ListOrders Doer is the HTTP client used to make requests to the listOrders
	// endpoint.
	ListOrdersDoer goahttp.Doer

	// GetTradePair Doer is the HTTP client used to make requests to the
	// getTradePair endpoint.
	GetTradePairDoer goahttp.Doer

	// ListTradePairs Doer is the HTTP client used to make requests to the
	// listTradePairs endpoint.
	ListTradePairsDoer goahttp.Doer

	// GetAccount Doer is the HTTP client used to make requests to the getAccount
	// endpoint.
	GetAccountDoer goahttp.Doer

	// GetOnlineAccounts Doer is the HTTP client used to make requests to the
	// getOnlineAccounts endpoint.
	GetOnlineAccountsDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the RestAPI service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetActiveOrderDoer:    doer,
		GetArchiveOrderDoer:   doer,
		ListOrdersDoer:        doer,
		GetTradePairDoer:      doer,
		ListTradePairsDoer:    doer,
		GetAccountDoer:        doer,
		GetOnlineAccountsDoer: doer,
		CORSDoer:              doer,
		RestoreResponseBody:   restoreBody,
		scheme:                scheme,
		host:                  host,
		decoder:               dec,
		encoder:               enc,
	}
}

// GetActiveOrder returns an endpoint that makes HTTP requests to the RestAPI
// service getActiveOrder server.
func (c *Client) GetActiveOrder() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetActiveOrderRequest(c.encoder)
		decodeResponse = DecodeGetActiveOrderResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetActiveOrderRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetActiveOrderDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RestAPI", "getActiveOrder", err)
		}
		return decodeResponse(resp)
	}
}

// GetArchiveOrder returns an endpoint that makes HTTP requests to the RestAPI
// service getArchiveOrder server.
func (c *Client) GetArchiveOrder() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetArchiveOrderRequest(c.encoder)
		decodeResponse = DecodeGetArchiveOrderResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetArchiveOrderRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetArchiveOrderDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RestAPI", "getArchiveOrder", err)
		}
		return decodeResponse(resp)
	}
}

// ListOrders returns an endpoint that makes HTTP requests to the RestAPI
// service listOrders server.
func (c *Client) ListOrders() goa.Endpoint {
	var (
		encodeRequest  = EncodeListOrdersRequest(c.encoder)
		decodeResponse = DecodeListOrdersResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListOrdersRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListOrdersDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RestAPI", "listOrders", err)
		}
		return decodeResponse(resp)
	}
}

// GetTradePair returns an endpoint that makes HTTP requests to the RestAPI
// service getTradePair server.
func (c *Client) GetTradePair() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetTradePairRequest(c.encoder)
		decodeResponse = DecodeGetTradePairResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetTradePairRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetTradePairDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RestAPI", "getTradePair", err)
		}
		return decodeResponse(resp)
	}
}

// ListTradePairs returns an endpoint that makes HTTP requests to the RestAPI
// service listTradePairs server.
func (c *Client) ListTradePairs() goa.Endpoint {
	var (
		encodeRequest  = EncodeListTradePairsRequest(c.encoder)
		decodeResponse = DecodeListTradePairsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListTradePairsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListTradePairsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RestAPI", "listTradePairs", err)
		}
		return decodeResponse(resp)
	}
}

// GetAccount returns an endpoint that makes HTTP requests to the RestAPI
// service getAccount server.
func (c *Client) GetAccount() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetAccountRequest(c.encoder)
		decodeResponse = DecodeGetAccountResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetAccountRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetAccountDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RestAPI", "getAccount", err)
		}
		return decodeResponse(resp)
	}
}

// GetOnlineAccounts returns an endpoint that makes HTTP requests to the
// RestAPI service getOnlineAccounts server.
func (c *Client) GetOnlineAccounts() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetOnlineAccountsRequest(c.encoder)
		decodeResponse = DecodeGetOnlineAccountsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetOnlineAccountsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetOnlineAccountsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RestAPI", "getOnlineAccounts", err)
		}
		return decodeResponse(resp)
	}
}
