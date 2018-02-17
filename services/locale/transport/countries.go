package transport

import (
	"context"
	"errors"
	"net/http"

	"github.com/andrew-waters/frdm/services/locale/service"
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

func MakeFindAllCountriesEndpoint(service service.Countries) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := requests.FindAllCountries{}
		return service.FindAllCountries(req)
	}
}

func DecodeFindAllCountriesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.FindAllCountries
	return request, nil
}

func MakeFindAllCountriesInContinentEndpoint(service service.Countries) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := requests.FindCountriesInContinent{
			ISO: "GB",
		}

		return service.FindCountriesInContinent(req)
	}
}

func DecodeFindAllCountriesInContinentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.FindCountriesInContinent

	vars := mux.Vars(r)

	iso, ok := vars["iso"]
	if !ok {
		return nil, errors.New("no iso")
	}

	request.ISO = iso

	return request, nil
}
