// Code generated by goa v3.1.1, DO NOT EDIT.
//
// RelayerAPI client HTTP transport
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

// Client lists the RelayerAPI service endpoint HTTP clients.
type Client struct {
	// AssetPairs Doer is the HTTP client used to make requests to the assetPairs
	// endpoint.
	AssetPairsDoer goahttp.Doer

	// Orders Doer is the HTTP client used to make requests to the orders endpoint.
	OrdersDoer goahttp.Doer

	// OrderByHash Doer is the HTTP client used to make requests to the orderByHash
	// endpoint.
	OrderByHashDoer goahttp.Doer

	// Orderbook Doer is the HTTP client used to make requests to the orderbook
	// endpoint.
	OrderbookDoer goahttp.Doer

	// OrderConfig Doer is the HTTP client used to make requests to the orderConfig
	// endpoint.
	OrderConfigDoer goahttp.Doer

	// FeeRecipients Doer is the HTTP client used to make requests to the
	// feeRecipients endpoint.
	FeeRecipientsDoer goahttp.Doer

	// PostOrder Doer is the HTTP client used to make requests to the postOrder
	// endpoint.
	PostOrderDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the RelayerAPI service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AssetPairsDoer:      doer,
		OrdersDoer:          doer,
		OrderByHashDoer:     doer,
		OrderbookDoer:       doer,
		OrderConfigDoer:     doer,
		FeeRecipientsDoer:   doer,
		PostOrderDoer:       doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// AssetPairs returns an endpoint that makes HTTP requests to the RelayerAPI
// service assetPairs server.
func (c *Client) AssetPairs() goa.Endpoint {
	var (
		encodeRequest  = EncodeAssetPairsRequest(c.encoder)
		decodeResponse = DecodeAssetPairsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAssetPairsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AssetPairsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RelayerAPI", "assetPairs", err)
		}
		return decodeResponse(resp)
	}
}

// Orders returns an endpoint that makes HTTP requests to the RelayerAPI
// service orders server.
func (c *Client) Orders() goa.Endpoint {
	var (
		encodeRequest  = EncodeOrdersRequest(c.encoder)
		decodeResponse = DecodeOrdersResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildOrdersRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.OrdersDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RelayerAPI", "orders", err)
		}
		return decodeResponse(resp)
	}
}

// OrderByHash returns an endpoint that makes HTTP requests to the RelayerAPI
// service orderByHash server.
func (c *Client) OrderByHash() goa.Endpoint {
	var (
		decodeResponse = DecodeOrderByHashResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildOrderByHashRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.OrderByHashDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RelayerAPI", "orderByHash", err)
		}
		return decodeResponse(resp)
	}
}

// Orderbook returns an endpoint that makes HTTP requests to the RelayerAPI
// service orderbook server.
func (c *Client) Orderbook() goa.Endpoint {
	var (
		encodeRequest  = EncodeOrderbookRequest(c.encoder)
		decodeResponse = DecodeOrderbookResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildOrderbookRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.OrderbookDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RelayerAPI", "orderbook", err)
		}
		return decodeResponse(resp)
	}
}

// OrderConfig returns an endpoint that makes HTTP requests to the RelayerAPI
// service orderConfig server.
func (c *Client) OrderConfig() goa.Endpoint {
	var (
		encodeRequest  = EncodeOrderConfigRequest(c.encoder)
		decodeResponse = DecodeOrderConfigResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildOrderConfigRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.OrderConfigDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RelayerAPI", "orderConfig", err)
		}
		return decodeResponse(resp)
	}
}

// FeeRecipients returns an endpoint that makes HTTP requests to the RelayerAPI
// service feeRecipients server.
func (c *Client) FeeRecipients() goa.Endpoint {
	var (
		decodeResponse = DecodeFeeRecipientsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildFeeRecipientsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.FeeRecipientsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RelayerAPI", "feeRecipients", err)
		}
		return decodeResponse(resp)
	}
}

// PostOrder returns an endpoint that makes HTTP requests to the RelayerAPI
// service postOrder server.
func (c *Client) PostOrder() goa.Endpoint {
	var (
		encodeRequest  = EncodePostOrderRequest(c.encoder)
		decodeResponse = DecodePostOrderResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPostOrderRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PostOrderDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("RelayerAPI", "postOrder", err)
		}
		return decodeResponse(resp)
	}
}
