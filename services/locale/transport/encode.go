package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/jsonapi"
)

// EncodeResponse is used by all endpoints to write the response interface{} to the w https.ResponseWriter
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok && err != nil {
		return err
		// 	// Not a Go kit transport error, but a business-logic error.
		// 	// Provide those as HTTP errors.
		// 	// encodeError(ctx, e.error(), w)
		// 	return json.NewEncoder(w).Encode(e)
	}
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	bytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Write(bytes)
	return nil
}
