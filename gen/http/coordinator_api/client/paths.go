// Code generated by goa v3.1.1, DO NOT EDIT.
//
// HTTP request path constructors for the CoordinatorAPI service.
//
// Command:
// $ goa gen github.com/InjectiveLabs/injective-core/api/design -o ../

package client

// ConfigurationCoordinatorAPIPath returns the URL path to the CoordinatorAPI service configuration HTTP endpoint.
func ConfigurationCoordinatorAPIPath() string {
	return "/api/coordinator/v2/configuration"
}

// RequestTransactionCoordinatorAPIPath returns the URL path to the CoordinatorAPI service request_transaction HTTP endpoint.
func RequestTransactionCoordinatorAPIPath() string {
	return "/api/coordinator/v2/request_transaction"
}

// SoftCancelsCoordinatorAPIPath returns the URL path to the CoordinatorAPI service soft_cancels HTTP endpoint.
func SoftCancelsCoordinatorAPIPath() string {
	return "/api/coordinator/v2/soft_cancels"
}
