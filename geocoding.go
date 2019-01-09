package yageocoding

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// YaGeoResponse contains response of API request
type YaGeoResponse struct {
	Response struct {
		ObjectCollection YaGeoObjectCollection `json:"GeoObjectCollection"`
	} `json:"response"`
}

// YaGeoObjectCollection contains metadata and members of response
type YaGeoObjectCollection struct {
	MetaData YaGeoMetaData `json:"metaDataProperty"`
	Members  []YaGeoMember `json:"featureMember"`
}

// YaGeoMetaData contains request string, count of founded elements and count of results
type YaGeoMetaData struct {
	ResponseMetaData struct {
		Request string `json:"request"`
		Found   string `json:"found"`
		Results string `json:"results"`
	} `json:"GeocoderResponseMetaData"`
}

// YaGeoInstance instance of Yandex Geocoding API
type YaGeoInstance struct {
	Key string
}

// New creates a new instance of Yandex Geocoding
func New(key string) *YaGeoInstance {
	return &YaGeoInstance{Key: key}
}

// Find returns result of search by address
func (ygi *YaGeoInstance) Find(address string) (result *YaGeoResponse, err error) {
	resp, err := http.Get(fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?format=json&geocode=%v&apikey=%v", address, ygi.Key))
	if err != nil {
		return result, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Members returns array of founded results of search
func (response *YaGeoResponse) Members() *[]YaGeoMember {
	return &response.Response.ObjectCollection.Members
}

// Address returns full address of first founded element
func (response *YaGeoResponse) Address() string {
	if len(response.Response.ObjectCollection.Members) > 0 {
		return response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Text
	}
	return ""
}

// Coordinates returns Latitude and Longitude of first member
func (response *YaGeoResponse) Coordinates() (latitude float64, longitude float64) {
	if len(response.Response.ObjectCollection.Members) == 0 {
		return 0, 0
	}
	coords := strings.Split(response.Response.ObjectCollection.Members[0].GeoObject.Point.Pos, " ")
	latitude, errlat := strconv.ParseFloat(coords[0], 64)
	if errlat != nil {
		return 0, 0
	}

	longitude, errlon := strconv.ParseFloat(coords[1], 64)
	if errlon != nil {
		return 0, 0
	}

	return
}

// Latitude of first member
func (response *YaGeoResponse) Latitude() float64 {
	if len(response.Response.ObjectCollection.Members) == 0 {
		return 0
	}
	coords := strings.Split(response.Response.ObjectCollection.Members[0].GeoObject.Point.Pos, " ")
	latitude, errlat := strconv.ParseFloat(coords[0], 64)
	if errlat != nil {
		return 0
	}
	return latitude
}

// Longitude of first member
func (response *YaGeoResponse) Longitude() float64 {
	if len(response.Response.ObjectCollection.Members) == 0 {
		return 0
	}
	coords := strings.Split(response.Response.ObjectCollection.Members[0].GeoObject.Point.Pos, " ")
	longitude, errlon := strconv.ParseFloat(coords[1], 64)
	if errlon != nil {
		return 0
	}
	return longitude
}

// CountryCode returns country code of first member
func (response *YaGeoResponse) CountryCode() string {
	if len(response.Response.ObjectCollection.Members) == 0 {
		return ""
	}
	return response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.CountryCode
}

// PostalCode returns postal code of first member
func (response *YaGeoResponse) PostalCode() string {
	if len(response.Response.ObjectCollection.Members) == 0 {
		return ""
	}
	return response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.PostalCode
}

// AddressComponents returns array of address components of first member
func (response *YaGeoResponse) AddressComponents() *[]YaGeoAddressComponent {
	if len(response.Response.ObjectCollection.Members) == 0 {
		return nil
	}
	return &response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.Components
}
