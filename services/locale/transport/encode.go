package transport

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	// if e, ok := response.(error); ok && e != nil {
	// 	// Not a Go kit transport error, but a business-logic error.
	// 	// Provide those as HTTP errors.
	// 	// encodeError(ctx, e.error(), w)
	// 	return json.NewEncoder(w).Encode(e)
	// }

	return json.NewEncoder(w).Encode(response)
}
