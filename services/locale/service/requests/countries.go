package requests

// FindCountriesInContinent request
type FindCountriesInContinent struct {
	ISO string
}

// FindAllCountries request
type FindAllCountries struct{}

//FindCountryWithISO request
type FindCountryWithISO struct {
	ISO string
}
