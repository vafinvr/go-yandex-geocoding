package yageocoding

// YaGeoAddress contains country code, postal code, formatted address and array of address components with type and name
type YaGeoAddress struct {
	CountryCode string                  `json:"country_code"`
	PostalCode  string                  `json:"postal_code"`
	Formatted   string                  `json:"formatted"`
	Components  []YaGeoAddressComponent `json:"Components"`
}

// YaGeoAddressComponent contains type and name of address component
type YaGeoAddressComponent struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}
