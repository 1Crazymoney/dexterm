// Code generated by goa v3.1.1, DO NOT EDIT.
//
// DebugAPI HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/InjectiveLabs/injective-core/api/design -o ../

package server

import (
	"context"
	"net/http"

	debugapi "github.com/InjectiveLabs/injective-core/api/gen/debug_api"
	goahttp "goa.design/goa/v3/http"
)

// EncodeVersionResponse returns an encoder for responses returned by the
// DebugAPI version endpoint.
func EncodeVersionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*debugapi.VersionResult)
		enc := encoder(ctx, w)
		body := NewVersionResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}