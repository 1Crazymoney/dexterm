// Code generated by goa v3.1.1, DO NOT EDIT.
//
// RestAPI endpoints
//
// Command:
// $ goa gen github.com/InjectiveLabs/injective-core/api/design -o ../

package restapi

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "RestAPI" service endpoints.
type Endpoints struct {
	GetActiveOrder        goa.Endpoint
	GetArchiveOrder       goa.Endpoint
	ListOrders            goa.Endpoint
	GetTradePair          goa.Endpoint
	ListTradePairs        goa.Endpoint
	ListDerivativeMarkets goa.Endpoint
	GetAccount            goa.Endpoint
	GetOnlineAccounts     goa.Endpoint
}

// NewEndpoints wraps the methods of the "RestAPI" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetActiveOrder:        NewGetActiveOrderEndpoint(s),
		GetArchiveOrder:       NewGetArchiveOrderEndpoint(s),
		ListOrders:            NewListOrdersEndpoint(s),
		GetTradePair:          NewGetTradePairEndpoint(s),
		ListTradePairs:        NewListTradePairsEndpoint(s),
		ListDerivativeMarkets: NewListDerivativeMarketsEndpoint(s),
		GetAccount:            NewGetAccountEndpoint(s),
		GetOnlineAccounts:     NewGetOnlineAccountsEndpoint(s),
	}
}

// Use applies the given middleware to all the "RestAPI" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetActiveOrder = m(e.GetActiveOrder)
	e.GetArchiveOrder = m(e.GetArchiveOrder)
	e.ListOrders = m(e.ListOrders)
	e.GetTradePair = m(e.GetTradePair)
	e.ListTradePairs = m(e.ListTradePairs)
	e.ListDerivativeMarkets = m(e.ListDerivativeMarkets)
	e.GetAccount = m(e.GetAccount)
	e.GetOnlineAccounts = m(e.GetOnlineAccounts)
}

// NewGetActiveOrderEndpoint returns an endpoint function that calls the method
// "getActiveOrder" of service "RestAPI".
func NewGetActiveOrderEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetActiveOrderPayload)
		return s.GetActiveOrder(ctx, p)
	}
}

// NewGetArchiveOrderEndpoint returns an endpoint function that calls the
// method "getArchiveOrder" of service "RestAPI".
func NewGetArchiveOrderEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetArchiveOrderPayload)
		return s.GetArchiveOrder(ctx, p)
	}
}

// NewListOrdersEndpoint returns an endpoint function that calls the method
// "listOrders" of service "RestAPI".
func NewListOrdersEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListOrdersPayload)
		return s.ListOrders(ctx, p)
	}
}

// NewGetTradePairEndpoint returns an endpoint function that calls the method
// "getTradePair" of service "RestAPI".
func NewGetTradePairEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetTradePairPayload)
		return s.GetTradePair(ctx, p)
	}
}

// NewListTradePairsEndpoint returns an endpoint function that calls the method
// "listTradePairs" of service "RestAPI".
func NewListTradePairsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListTradePairsPayload)
		return s.ListTradePairs(ctx, p)
	}
}

// NewListDerivativeMarketsEndpoint returns an endpoint function that calls the
// method "listDerivativeMarkets" of service "RestAPI".
func NewListDerivativeMarketsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListDerivativeMarkets(ctx)
	}
}

// NewGetAccountEndpoint returns an endpoint function that calls the method
// "getAccount" of service "RestAPI".
func NewGetAccountEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetAccountPayload)
		return s.GetAccount(ctx, p)
	}
}

// NewGetOnlineAccountsEndpoint returns an endpoint function that calls the
// method "getOnlineAccounts" of service "RestAPI".
func NewGetOnlineAccountsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetOnlineAccountsPayload)
		return s.GetOnlineAccounts(ctx, p)
	}
}
