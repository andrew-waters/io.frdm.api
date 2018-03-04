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

// MakeFindAllCountriesEndpoint creates an endpoint to serve requests for all countries
func MakeFindAllCountriesEndpoint(service service.Countries) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := requests.FindAllCountries{}
		return service.FindAllCountries(req)
	}
}

// MakeFindAllCountriesInContinentEndpoint creates an endpoint to serve requests for getting all countries in a continent
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

// MakeFindCountryByISOEndpoint creates an endpoint to serve requests for getting a single country by ISO
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

// DecodeFindAllCountriesRequest consumes the request to find all countries and returns a concrete request
func DecodeFindAllCountriesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request requests.FindAllCountries
	return request, nil
}

// DecodeFindAllCountriesInContinentRequest consumes the request to find all countries and returns a concrete
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

// DecodeFindCountryByISORequest consumes the request to find a single country by ISO and returns a concrete
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
