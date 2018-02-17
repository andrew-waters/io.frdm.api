package transport

import (
	"context"
	"net/http"

	"github.com/andrew-waters/frdm/services/locale/service"
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/go-kit/kit/endpoint"
)

func MakeFindAllContinentsEndpoint(service service.Continents) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := requests.FindAllContinents{}
		return service.FindAllContinents(req)
	}
}

func DecodeFindAllContinentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.FindAllContinents
	return request, nil
}
