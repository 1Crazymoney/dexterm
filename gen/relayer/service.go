// Code generated by goa v3.0.9, DO NOT EDIT.
//
// Relayer service
//
// Command:
// $ goa gen github.com/InjectiveLabs/injective-core/api/design -o ../../api

package relayer

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Relayer implements REST API for RelayerDaemon with SRAv3 support.
type Service interface {
	// Retrieves a list of available asset pairs and the information required to
	// trade them.
	AssetPairs(context.Context, *AssetPairsPayload) (res *AssetPairsResult, err error)
	// Retrieves a list of orders given query parameters. For querying an entire
	// orderbook snapshot, the orderbook endpoint is recommended.
	Orders(context.Context, *OrdersPayload) (res *OrdersResult, err error)
	// Retrieves a specific order by orderHash.
	OrderByHash(context.Context, *OrderByHashPayload) (res *OrderByHashResult, err error)
	// Retrieves the orderbook for a given asset pair.
	Orderbook(context.Context, *OrderbookPayload) (res *OrderbookResult, err error)
	// Submit a partial order and receive information required to complete the
	// order.
	OrderConfig(context.Context, *OrderConfigPayload) (res *OrderConfigResult, err error)
	// Retrieves a list of all fee recipient addresses for a relayer.
	FeeRecipients(context.Context) (res *FeeRecipientsResult, err error)
	// Submit a signed make order to the relayer.
	PostOrder(context.Context, *PostOrderPayload) (res *PostOrderResult, err error)
	// Submit a signed take order to the relayer.
	TakeOrder(context.Context, *TakeOrderPayload) (res *TakeOrderResult, err error)
	// Retrieves an active 0x order with meta info that is associated with the hash.
	GetActiveOrder(context.Context, *GetActiveOrderPayload) (res *GetActiveOrderResult, err error)
	// Retrieves an archive 0x order with meta info that is associated with the
	// hash.
	GetArchiveOrder(context.Context, *GetArchiveOrderPayload) (res *GetArchiveOrderResult, err error)
	// Retrieves a list of 0x orders matching the filtering rules.
	ListOrders(context.Context, *ListOrdersPayload) (res *ListOrdersResult, err error)
	// Retrieves a trade pair by name or hash.
	GetTradePair(context.Context, *GetTradePairPayload) (res *GetTradePairResult, err error)
	// Retrieves a list trade pairs.
	ListTradePairs(context.Context, *ListTradePairsPayload) (res *ListTradePairsResult, err error)
	// Retrieves a relayer account by address.
	GetAccount(context.Context, *GetAccountPayload) (res *GetAccountResult, err error)
	// Retrieves online relayer accounts only.
	GetOnlineAccounts(context.Context, *GetOnlineAccountsPayload) (res *GetOnlineAccountsResult, err error)
	// Retrieves Ethereum transactions that match specified filters.
	GetEthTransactions(context.Context, *GetEthTransactionsPayload) (res *GetEthTransactionsResult, err error)
	// Returns relayerd version.
	Version(context.Context) (res *VersionResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Relayer"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [17]string{"assetPairs", "orders", "orderByHash", "orderbook", "orderConfig", "feeRecipients", "postOrder", "takeOrder", "getActiveOrder", "getArchiveOrder", "listOrders", "getTradePair", "listTradePairs", "getAccount", "getOnlineAccounts", "getEthTransactions", "version"}

// AssetPairsPayload is the payload type of the Relayer service assetPairs
// method.
type AssetPairsPayload struct {
	// Asset data field of first asset in the pair.
	AssetDataA *string
	// Asset data field of second asset in the pair.
	AssetDataB *string
}

// AssetPairsResult is the result type of the Relayer service assetPairs method.
type AssetPairsResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Total records found in collection.
	Total int
	// The page number, starts from 1.
	Page int
	// Records limit per each page.
	PerPage int
	// Asset pairs that contain assetDataA and assetDataB (listed in any order).
	Records []*AssetPairRecord
}

// OrdersPayload is the payload type of the Relayer service orders method.
type OrdersPayload struct {
	// Filters orders where the maker asset is of certain asset proxy id.
	MakerAssetProxyID *string
	// Filters orders where the taker asset is of certain asset proxy id.
	TakerAssetProxyID *string
	// Filters orders where the contract address for the maker asset matches the
	// value specified.
	MakerAssetAddress *string
	// Filters orders where the contract address for the taker asset matches the
	// value specified.
	TakerAssetAddress *string
	// Filters orders created for this exchange address
	ExchangeAddress *string
	// Filters orders with the specified senderAddress
	SenderAddress *string
	// Filters orders with the specified makerAssetData
	MakerAssetData *string
	// Filters orders with the specified takerAssetData
	TakerAssetData *string
	// Filters orders with the specified makerFeeAssetData
	MakerFeeAssetData *string
	// Filters orders with the specified takerFeeAssetData
	TakerFeeAssetData *string
	// Filters orders where either makerAssetData or takerAssetData has the value
	// specified
	TraderAssetData *string
	// Filters orders with the specified makerAddress
	MakerAddress *string
	// Filters orders with the specified takerAddress
	TakerAddress *string
	// Filters orders where either makerAddress or takerAddress has the value
	// specified
	TraderAddress *string
	// Filters orders where feeRecipientAddress is feeRecipient address
	FeeRecipientAddress *string
}

// OrdersResult is the result type of the Relayer service orders method.
type OrdersResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Total records found in collection.
	Total int
	// The page number, starts from 1.
	Page int
	// Records limit per each page.
	PerPage int
	// If both makerAssetData and takerAssetData are specified, returned orders
	// will be sorted by price determined by (takerAssetAmount/makerAssetAmount).
	Records []*OrderRecord
}

// OrderByHashPayload is the payload type of the Relayer service orderByHash
// method.
type OrderByHashPayload struct {
	// Specifies order hash.
	OrderHash string
}

// OrderByHashResult is the result type of the Relayer service orderByHash
// method.
type OrderByHashResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Order item.
	Order *Order
	// Additional meta data.
	MetaData interface{}
}

// OrderbookPayload is the payload type of the Relayer service orderbook method.
type OrderbookPayload struct {
	// Asset data (makerAssetData or takerAssetData) designated as the base
	// currency in the currency pair calculation of price.
	BaseAssetData string
	// Asset data (makerAssetData or takerAssetData) designated as the quote
	// currency in the currency pair calculation of price.
	QuoteAssetData string
}

// OrderbookResult is the result type of the Relayer service orderbook method.
type OrderbookResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Array of signed orders where takerAssetData is equal to baseAssetData. Bids
	// will be sorted in descending order by price.
	Bids *OrderbookRecords
	// Array of signed orders where makerAssetData is equal to baseAssetData. Asks
	// will be sorted in ascending order by price.
	Asks *OrderbookRecords
}

// OrderConfigPayload is the payload type of the Relayer service orderConfig
// method.
type OrderConfigPayload struct {
	// Specify chain ID.
	ChainID int64
	// Exchange v2 contract address.
	ExchangeAddress string
	// Address that created the order.
	MakerAddress string
	// Address that is allowed to fill the order. If set to 0, any address is
	// allowed to fill the order.
	TakerAddress string
	// Amount of makerAsset being offered by maker. Must be greater than 0.
	MakerAssetAmount string
	// Amount of takerAsset being bid on by maker. Must be greater than 0.
	TakerAssetAmount string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerAsset.
	MakerAssetData string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerAsset.
	TakerAssetData string
	// Timestamp in seconds at which order expires.
	ExpirationTimeSeconds string
}

// OrderConfigResult is the result type of the Relayer service orderConfig
// method.
type OrderConfigResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Address that is allowed to call Exchange contract methods that affect this
	// order. If set to 0, any address is allowed to call these methods.
	SenderAddress string
	// Address that will receive fees when order is filled.
	FeeRecipientAddress string
	// Amount of ZRX paid to feeRecipient by maker when order is filled. If set to
	// 0, no transfer of ZRX from maker to feeRecipient will be attempted.
	MakerFee string
	// Amount of ZRX paid to feeRecipient by taker when order is filled. If set to
	// 0, no transfer of ZRX from taker to feeRecipient will be attempted.
	TakerFee string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerFee.
	MakerFeeAssetData string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerFee.
	TakerFeeAssetData string
}

// FeeRecipientsResult is the result type of the Relayer service feeRecipients
// method.
type FeeRecipientsResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Total records found in collection.
	Total int
	// The page number, starts from 1.
	Page int
	// Records limit per each page.
	PerPage int
	// List of all fee recipient addresses for a relayer
	Records []string
}

// PostOrderPayload is the payload type of the Relayer service postOrder method.
type PostOrderPayload struct {
	// Specify chain ID.
	ChainID int64
	// Exchange v2 contract address.
	ExchangeAddress string
	// Address that created the order.
	MakerAddress string
	// Address that is allowed to fill the order. If set to 0, any address is
	// allowed to fill the order.
	TakerAddress string
	// Address that will receive fees when order is filled.
	FeeRecipientAddress string
	// Address that is allowed to call Exchange contract methods that affect this
	// order. If set to 0, any address is allowed to call these methods.
	SenderAddress string
	// Amount of makerAsset being offered by maker. Must be greater than 0.
	MakerAssetAmount string
	// Amount of takerAsset being bid on by maker. Must be greater than 0.
	TakerAssetAmount string
	// Amount of ZRX paid to feeRecipient by maker when order is filled. If set to
	// 0, no transfer of ZRX from maker to feeRecipient will be attempted.
	MakerFee string
	// Amount of ZRX paid to feeRecipient by taker when order is filled. If set to
	// 0, no transfer of ZRX from taker to feeRecipient will be attempted.
	TakerFee string
	// Timestamp in seconds at which order expires.
	ExpirationTimeSeconds string
	// Arbitrary number to facilitate uniqueness of the order's hash.
	Salt string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerAsset.
	MakerAssetData string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerAsset.
	TakerAssetData string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerFee.
	MakerFeeAssetData string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerFee.
	TakerFeeAssetData string
	// Order signature.
	Signature string
}

// PostOrderResult is the result type of the Relayer service postOrder method.
type PostOrderResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
}

// TakeOrderPayload is the payload type of the Relayer service takeOrder method.
type TakeOrderPayload struct {
	// Signed 0x take order.
	TakeOrder *Order
	// List of 0x signed make orders.
	MakeOrders []*Order
	// List of make order fill amounts
	MakeOrderFillAmounts []string
}

// TakeOrderResult is the result type of the Relayer service takeOrder method.
type TakeOrderResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
}

// GetActiveOrderPayload is the payload type of the Relayer service
// getActiveOrder method.
type GetActiveOrderPayload struct {
	// The hash of the desired 0x order.
	OrderHash string
}

// GetActiveOrderResult is the result type of the Relayer service
// getActiveOrder method.
type GetActiveOrderResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Found active 0x order.
	Order *Order
	// Additional meta data.
	MetaData interface{}
}

// GetArchiveOrderPayload is the payload type of the Relayer service
// getArchiveOrder method.
type GetArchiveOrderPayload struct {
	// The hash of the desired 0x order.
	OrderHash string
}

// GetArchiveOrderResult is the result type of the Relayer service
// getArchiveOrder method.
type GetArchiveOrderResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Order item.
	Order *Order
	// Additional meta data.
	MetaData interface{}
}

// ListOrdersPayload is the payload type of the Relayer service listOrders
// method.
type ListOrdersPayload struct {
	// Filter by status of the order
	Status *string
	// Filter by collection of the order
	Collection *string
	// Filter by trade pair hash
	TradePairHash *string
	// Enabled sort by VDF (1 = asc, -1 = desc)
	SortByVdf *int32
	// Limits the amout of results by top N
	Limit *int32
}

// ListOrdersResult is the result type of the Relayer service listOrders method.
type ListOrdersResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Filtered take orders.
	TakeOrders []*Order
	// Filtered make orders.
	MakeOrders []*Order
	// Additional meta data.
	MetaData interface{}
}

// GetTradePairPayload is the payload type of the Relayer service getTradePair
// method.
type GetTradePairPayload struct {
	// Specify name of the trade pair.
	Name *string
	// Most effective way is to specify hash of the trade pair.
	Hash *string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerAsset.
	MakerAssetData *string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerAsset.
	TakerAssetData *string
}

// GetTradePairResult is the result type of the Relayer service getTradePair
// method.
type GetTradePairResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Found trade pair.
	TradePair *TradePair
	// Additional meta data.
	MetaData interface{}
}

// ListTradePairsPayload is the payload type of the Relayer service
// listTradePairs method.
type ListTradePairsPayload struct {
	// Specify to include all trade pairs, suspended and active.
	All *bool
}

// ListTradePairsResult is the result type of the Relayer service
// listTradePairs method.
type ListTradePairsResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Filtered trade pairs.
	TradePairs []*TradePair
	// Additional meta data.
	MetaData interface{}
}

// GetAccountPayload is the payload type of the Relayer service getAccount
// method.
type GetAccountPayload struct {
	// Address of the relayer account.
	Address string
}

// GetAccountResult is the result type of the Relayer service getAccount method.
type GetAccountResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Found relayer account.
	Account *RelayerAccount
	// Additional meta data.
	MetaData interface{}
}

// GetOnlineAccountsPayload is the payload type of the Relayer service
// getOnlineAccounts method.
type GetOnlineAccountsPayload struct {
	// Specify a version to filter online accounts with this logic version.
	Version *string
	// Specify a threshold in seconds to filter online accounts within a period of
	// time.
	Threshold *int64
}

// GetOnlineAccountsResult is the result type of the Relayer service
// getOnlineAccounts method.
type GetOnlineAccountsResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Filtered online relayer accounts.
	Accounts []*RelayerAccount
	// Additional meta data.
	MetaData interface{}
}

// GetEthTransactionsPayload is the payload type of the Relayer service
// getEthTransactions method.
type GetEthTransactionsPayload struct {
	// Specify Cosmos address of the proposer for that transaction.
	Proposer *string
	// Ethereum wallet address of the transaction sender (proposer).
	From *string
	// Filter only transactions that contain specified trade hash.
	TradeHash *string
	// Specify lower boundary for the transaction's block.
	BlockFrom *int64
	// Specify upper boundary for the transaction's block.
	BlockBefore *int64
	// Filter only transactions that contain a review from specified cosmos address.
	ReviewedBy *string
}

// GetEthTransactionsResult is the result type of the Relayer service
// getEthTransactions method.
type GetEthTransactionsResult struct {
	// The maximum number of requests you're permitted to make per hour.
	RLimitLimit *int
	// The number of requests remaining in the current rate limit window.
	RLimitRemaining *int
	// The time at which the current rate limit window resets in UTC epoch seconds.
	RLimitReset *int
	// Filtered Ethereum transactions.
	Transactions []*EthTransaction
	// Additional meta data.
	MetaData interface{}
}

// VersionResult is the result type of the Relayer service version method.
type VersionResult struct {
	// Relayerd code version.
	Version string
	// Additional meta data.
	MetaData map[string]string
}

type AssetPairRecord struct {
	// First asset record of the pair.
	AssetDataA *AssetRecord
	// Second asset record of the pair.
	AssetDataB *AssetRecord
}

type AssetRecord struct {
	// ABIv2 encoded assetData representing that token.
	AssetData string
	// The minimum trade amount the relayer will accept.
	MinAmount string
	// The maximum trade amount the relayer will accept.
	MaxAmount string
	// The desired price precision a relayer would like to support within their
	// orderbook.
	Precision string
}

type OrderRecord struct {
	// Order item.
	Order *Order
	// Additional meta data.
	MetaData interface{}
}

// A valid signed 0x order based on the schema.
type Order struct {
	// Specify chain ID.
	ChainID int64
	// Exchange v2 contract address.
	ExchangeAddress string
	// Address that created the order.
	MakerAddress string
	// Address that is allowed to fill the order. If set to 0, any address is
	// allowed to fill the order.
	TakerAddress string
	// Address that will receive fees when order is filled.
	FeeRecipientAddress string
	// Address that is allowed to call Exchange contract methods that affect this
	// order. If set to 0, any address is allowed to call these methods.
	SenderAddress string
	// Amount of makerAsset being offered by maker. Must be greater than 0.
	MakerAssetAmount string
	// Amount of takerAsset being bid on by maker. Must be greater than 0.
	TakerAssetAmount string
	// Amount of ZRX paid to feeRecipient by maker when order is filled. If set to
	// 0, no transfer of ZRX from maker to feeRecipient will be attempted.
	MakerFee string
	// Amount of ZRX paid to feeRecipient by taker when order is filled. If set to
	// 0, no transfer of ZRX from taker to feeRecipient will be attempted.
	TakerFee string
	// Timestamp in seconds at which order expires.
	ExpirationTimeSeconds string
	// Arbitrary number to facilitate uniqueness of the order's hash.
	Salt string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerAsset.
	MakerAssetData string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerAsset.
	TakerAssetData string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerFee.
	MakerFeeAssetData string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerFee.
	TakerFeeAssetData string
	// Order signature.
	Signature string
}

// A paginated array of signed orders.
type OrderbookRecords struct {
	// Total records found in collection.
	Total int
	// The page number, starts from 1.
	Page int
	// Records limit per each page.
	PerPage int
	// Array of signed orders
	Records []*OrderRecord
}

// An object describing trade pair of two assets.
type TradePair struct {
	// A name of the pair in format AAA/BBB, where AAA - maker's asset, BBB -
	// taker's asset.
	Name string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring makerAsset.
	MakerAssetData string
	// ABIv2 encoded data that can be decoded by a specified proxy contract when
	// transferring takerAsset.
	TakerAssetData string
	// Hash of both asset data, to identify the trading pair in store.
	Hash string
	// If false, then the pair is suspended and trades cannot be made.
	Enabled bool
}

// A relayer account that enhances Cosmos Account with liveness info.
type RelayerAccount struct {
	// Cosmos address of the relayer account.
	Address string
	// Public key of the relayer account, as hex string.
	PublicKey string
	// Timestamp in UNIX seconds of the last time seen.
	LastSeen int64
	// Last logic version seen.
	LastVersion string
	// A flag of liveness status of the account. Must be considered with lastSeen
	// timestamp.
	IsOnline bool
}

// An Ethereum transaction that must be submitted by proposer and validated by
// reviewers.
type EthTransaction struct {
	// Proposer is a leader chosen from online validators pool that must commit the
	// transaction into Ethereum.
	Proposer string
	// A list of trade references to be submitted in the transaction.
	TradeHashes []string
	// The block height when the transaction has been created and must've been
	// acknowledged by the selected proposer.
	Block int64
	// TxHash is the reported hash of an Ethereum transaction.
	TxHash *string
	// Ethereum address the transaction has been sent from (proposer's wallet).
	From *string
	// ReviewedBy is a list of other validators that reviewed transaction existence
	// and correctness.
	ReviewedBy []string
}

// Error and description for bad requests.
type ErrorBadRequest struct {
	// General error code
	Code int
	// Error reason description
	Reason string
	// A list of explained validation errors.
	ValidationErrors []*ValidationError
}

// Order validation error explained
type ValidationError struct {
	// Validation error code
	Code int
	// Validation error reason description
	Reason string
	// Field name
	Field *string
}

// Error returns an error description.
func (e *ErrorBadRequest) Error() string {
	return "Error and description for bad requests."
}

// ErrorName returns "ErrorBadRequest".
func (e *ErrorBadRequest) ErrorName() string {
	return "validation_error"
}

// Error returns an error description.
func (e *ValidationError) Error() string {
	return "Order validation error explained"
}

// ErrorName returns "ValidationError".
func (e *ValidationError) ErrorName() string {
	return "ValidationError"
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeRateLimit builds a goa.ServiceError from an error.
func MakeRateLimit(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "rate_limit",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInternal builds a goa.ServiceError from an error.
func MakeInternal(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotImplemented builds a goa.ServiceError from an error.
func MakeNotImplemented(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_implemented",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
