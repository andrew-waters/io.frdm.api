package responses

type Country struct {
	ISO       string `json:"iso"`
	ISO3      string `json:"iso3,omitempty"`
	NumCode   int    `json:"numcode,omitempty"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Continent string `json:"continent,omitempty"`
}
