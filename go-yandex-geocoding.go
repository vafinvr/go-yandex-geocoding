package yageo

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
	Members   []YaGeoMember `json:"featureMember"`
}

// YaGeoMetaData contains request string, count of founded elements and count of results
type YaGeoMetaData struct {
	ResponseMetaData struct {
		Request string `json:"request"`
		Found   string `json:"found"`
		Results string `json:"results"`
	} `json:"GeocoderResponseMetaData"`
}

// YaGeoMember is a structure of founded element. Contains metadata, description, name and coordinates of founded element.
type YaGeoMember struct {
	GeoObject struct {
		MetaData    YaGeoMemberMetaData `json:"metaDataProperty"`
		Description string              `json:"description"`
		Name        string              `json:"name"`
		Point       struct {
			Pos string `json:"pos"`
		} `json:"Point"`
	} `json:"GeoObject"`
}

// YaGeoMemberMetaData contains type of founded element, address and precision
type YaGeoMemberMetaData struct {
	Meta struct {
		Kind      string `json:"kind"`
		Text      string `json:"text"`
		Precision string `json:"precision"`
	} `json:"GeocoderMetaData"`
}

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
func (response YaGeoResponse) Members() []YaGeoMember {
	return response.Response.ObjectCollection.Members
}

// Address returns full address of first founded element
func (response YaGeoResponse) Address() string {
	if len(response.Response.ObjectCollection.Members) > 0 {
		return response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Text
	}
	return ""
}

// Coordinates returns Latitude and Longitude of member
func (member YaGeoMember) Coordinates() (latitude float64, longitude float64) {
	coords := strings.Split(member.GeoObject.Point.Pos, " ")
	latitude, errlat := strconv.ParseFloat(coords[0], 64)
	if errlat != nil {
		return 0,0
	}

	longitude, errlon := strconv.ParseFloat(coords[1], 64)
	if errlon != nil {
		return 0,0
	}

	return
}

// Latitude of member
func (member YaGeoMember) Latitude() float64 {
	coords := strings.Split(member.GeoObject.Point.Pos, " ")
	latitude, errlat := strconv.ParseFloat(coords[0], 64)
	if errlat != nil {
		return 0
	}
	return latitude
}

// Longitude of member
func (member YaGeoMember) Longitude() float64 {
	coords := strings.Split(member.GeoObject.Point.Pos, " ")
	longitude, errlon := strconv.ParseFloat(coords[1], 64)
	if errlon != nil {
		return 0
	}
	return longitude
}