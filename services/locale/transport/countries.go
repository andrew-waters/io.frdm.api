package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/andrew-waters/frdm/services/locale/service"
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/andrew-waters/frdm/services/locale/service/responses"
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
		r, ok := request.(requests.FindCountriesInContinent)
		if ok {
			return service.FindCountriesInContinent(r)
		}
		err = errors.New("invalid request")
		return nil, err
	}
}

func MakeFindCountryByISOEndpoint(service service.Countries) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(requests.FindCountryWithISO)
		if ok {
			return service.FindCountryWithISO(r)
		}
		err = errors.New("invalid request")
		return nil, err
	}
}

func DecodeFindCountryByISORequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.FindCountryWithISO
	vars := mux.Vars(r)
	iso, ok := vars["iso"]
	if !ok {
		return nil, errors.New("no iso")
	}
	request.ISO = iso
	return request, nil
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

func EncodeFindAllCountriesInContinentRequest(_ context.Context, w http.ResponseWriter, response []responses.Country) error {

	// if e, ok := response.(error); ok && e != nil {
	// 	panic("error")
	// 	// 	// Not a Go kit transport error, but a business-logic error.
	// 	// 	// Provide those as HTTP errors.
	// 	// 	// encodeError(ctx, e.error(), w)
	// 	// 	return json.NewEncoder(w).Encode(e)
	// }
	// w.Header().Set("Content-Type", jsonapi.MediaType)
	// w.WriteHeader(http.StatusOK)
	// return jsonapi.MarshalPayload(w, response)
	return json.NewEncoder(w).Encode(response)
}
