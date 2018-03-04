package transport

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

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

func MakeFindContinentWithISO(service service.Continents) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {

		r, ok := request.(requests.FindContinentByISO)
		if ok {
			return service.FindByISO(r)
		}
		err = errors.New("invalid request")
		return nil, err
	}
}

func DecodeFindContinentWithISORequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.FindContinentByISO

	vars := mux.Vars(r)
	iso, ok := vars["iso"]
	if ok {
		request.ISO = iso
		return request, nil
	}
	return nil, errors.New("no iso")
}

func DecodeFindAllContinentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.FindAllContinents
	return request, nil
}
