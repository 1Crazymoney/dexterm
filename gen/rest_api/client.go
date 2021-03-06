// Code generated by goa v3.1.1, DO NOT EDIT.
//
// RestAPI client
//
// Command:
// $ goa gen github.com/InjectiveLabs/injective-core/api/design -o ../

package restapi

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "RestAPI" service client.
type Client struct {
	GetActiveOrderEndpoint        goa.Endpoint
	GetArchiveOrderEndpoint       goa.Endpoint
	ListOrdersEndpoint            goa.Endpoint
	GetTradePairEndpoint          goa.Endpoint
	ListTradePairsEndpoint        goa.Endpoint
	ListDerivativeMarketsEndpoint goa.Endpoint
	GetAccountEndpoint            goa.Endpoint
	GetOnlineAccountsEndpoint     goa.Endpoint
}

// NewClient initializes a "RestAPI" service client given the endpoints.
func NewClient(getActiveOrder, getArchiveOrder, listOrders, getTradePair, listTradePairs, listDerivativeMarkets, getAccount, getOnlineAccounts goa.Endpoint) *Client {
	return &Client{
		GetActiveOrderEndpoint:        getActiveOrder,
		GetArchiveOrderEndpoint:       getArchiveOrder,
		ListOrdersEndpoint:            listOrders,
		GetTradePairEndpoint:          getTradePair,
		ListTradePairsEndpoint:        listTradePairs,
		ListDerivativeMarketsEndpoint: listDerivativeMarkets,
		GetAccountEndpoint:            getAccount,
		GetOnlineAccountsEndpoint:     getOnlineAccounts,
	}
}

// GetActiveOrder calls the "getActiveOrder" endpoint of the "RestAPI" service.
func (c *Client) GetActiveOrder(ctx context.Context, p *GetActiveOrderPayload) (res *GetActiveOrderResult, err error) {
	var ires interface{}
	ires, err = c.GetActiveOrderEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetActiveOrderResult), nil
}

// GetArchiveOrder calls the "getArchiveOrder" endpoint of the "RestAPI"
// service.
func (c *Client) GetArchiveOrder(ctx context.Context, p *GetArchiveOrderPayload) (res *GetArchiveOrderResult, err error) {
	var ires interface{}
	ires, err = c.GetArchiveOrderEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetArchiveOrderResult), nil
}

// ListOrders calls the "listOrders" endpoint of the "RestAPI" service.
func (c *Client) ListOrders(ctx context.Context, p *ListOrdersPayload) (res *ListOrdersResult, err error) {
	var ires interface{}
	ires, err = c.ListOrdersEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListOrdersResult), nil
}

// GetTradePair calls the "getTradePair" endpoint of the "RestAPI" service.
func (c *Client) GetTradePair(ctx context.Context, p *GetTradePairPayload) (res *GetTradePairResult, err error) {
	var ires interface{}
	ires, err = c.GetTradePairEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetTradePairResult), nil
}

// ListTradePairs calls the "listTradePairs" endpoint of the "RestAPI" service.
func (c *Client) ListTradePairs(ctx context.Context, p *ListTradePairsPayload) (res *ListTradePairsResult, err error) {
	var ires interface{}
	ires, err = c.ListTradePairsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListTradePairsResult), nil
}

// ListDerivativeMarkets calls the "listDerivativeMarkets" endpoint of the
// "RestAPI" service.
func (c *Client) ListDerivativeMarkets(ctx context.Context) (res *ListDerivativeMarketsResult, err error) {
	var ires interface{}
	ires, err = c.ListDerivativeMarketsEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*ListDerivativeMarketsResult), nil
}

// GetAccount calls the "getAccount" endpoint of the "RestAPI" service.
func (c *Client) GetAccount(ctx context.Context, p *GetAccountPayload) (res *GetAccountResult, err error) {
	var ires interface{}
	ires, err = c.GetAccountEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetAccountResult), nil
}

// GetOnlineAccounts calls the "getOnlineAccounts" endpoint of the "RestAPI"
// service.
func (c *Client) GetOnlineAccounts(ctx context.Context, p *GetOnlineAccountsPayload) (res *GetOnlineAccountsResult, err error) {
	var ires interface{}
	ires, err = c.GetOnlineAccountsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetOnlineAccountsResult), nil
}
