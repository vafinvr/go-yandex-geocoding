package yageocoding

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
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

	if len(result.Response.ObjectCollection.Members) == 0 {
		return result, errors.New("Not found")
	}

	return result, nil
}

// RangeBtw returns range in meters between two addresses (generates two requests to API)
func (ygi *YaGeoInstance) RangeBtw(address1, address2 string) (float64, error) {
	addr1, err1 := ygi.Find(address1)
	if err1 != nil {
		return 0, err1
	}
	addr2, err2 := ygi.Find(address2)
	if err2 != nil {
		return 0, err2
	}

	earthRadius := float64(6371000) // Earth's radius in meters
	difLat := deg2rad(addr1.Latitude() - addr2.Latitude())
	difLng := deg2rad(addr1.Longitude() - addr2.Longitude())
	a := math.Sin(difLat/2)*math.Sin(difLat/2) +
		math.Cos(addr2.Latitude())*math.Cos(addr1.Latitude())*
			math.Sin(difLng/2)*math.Sin(difLng/2)
	c := 2 * math.Asin(math.Sqrt(a))
	distance := earthRadius * c

	return distance, nil
}

// RangeToResponse returns range in meters to another address by completed response
func (response *YaGeoResponse) RangeToResponse(resp *YaGeoResponse) float64 {
	earthRadius := float64(6371000) // Earth's radius in meters
	difLat := deg2rad(response.Latitude() - resp.Latitude())
	difLng := deg2rad(response.Longitude() - resp.Longitude())
	a := math.Sin(difLat/2)*math.Sin(difLat/2) +
		math.Cos(resp.Latitude())*math.Cos(response.Latitude())*
			math.Sin(difLng/2)*math.Sin(difLng/2)
	c := 2 * math.Asin(math.Sqrt(a))
	distance := earthRadius * c

	return distance
}

// Members returns array of founded results of search
func (response *YaGeoResponse) Members() *[]YaGeoMember {
	return &response.Response.ObjectCollection.Members
}

// Address returns full address of first founded element
func (response *YaGeoResponse) Address() string {
	str := ""
	if len(response.Response.ObjectCollection.Members) > 0 {
		str = response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Text
	}
	return str
}

// Coordinates returns Latitude and Longitude of first member
func (response *YaGeoResponse) Coordinates() (latitude float64, longitude float64) {
	latitude = 0
	longitude = 0
	if len(response.Response.ObjectCollection.Members) > 0 {
		coords := strings.Split(response.Response.ObjectCollection.Members[0].GeoObject.Point.Pos, " ")
		latitude, _ = strconv.ParseFloat(coords[1], 64)
		longitude, _ = strconv.ParseFloat(coords[0], 64)
	}
	return
}

// Latitude of first member
func (response *YaGeoResponse) Latitude() float64 {
	latitude := 0.0
	if len(response.Response.ObjectCollection.Members) > 0 {
		coords := strings.Split(response.Response.ObjectCollection.Members[0].GeoObject.Point.Pos, " ")
		latitude, _ = strconv.ParseFloat(coords[1], 64)
	}
	return latitude
}

// Longitude of first member
func (response *YaGeoResponse) Longitude() float64 {
	longitude := 0.0
	if len(response.Response.ObjectCollection.Members) > 0 {
		coords := strings.Split(response.Response.ObjectCollection.Members[0].GeoObject.Point.Pos, " ")
		longitude, _ = strconv.ParseFloat(coords[0], 64)
	}
	return longitude
}

// CountryCode returns country code of first member
func (response *YaGeoResponse) CountryCode() string {
	str := ""
	if len(response.Response.ObjectCollection.Members) > 0 {
		str = response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.CountryCode
	}
	return str
}

// PostalCode returns postal code of first member
func (response *YaGeoResponse) PostalCode() string {
	str := ""
	if len(response.Response.ObjectCollection.Members) > 0 {
		str = response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.PostalCode
	}
	return str
}

// AddressComponents returns array of address components of first member
func (response *YaGeoResponse) AddressComponents() *[]YaGeoAddressComponent {
	var arr *[]YaGeoAddressComponent = nil
	if len(response.Response.ObjectCollection.Members) > 0 {
		arr = &response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.Components
	}
	return arr
}

// Country returns country name for first address
func (response *YaGeoResponse) Country() string {
	str := ""
	for _, v := range response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "country" {
			str = v.Name
		}
	}
	return str
}

// Province returns array of province names for first address
func (response *YaGeoResponse) Province() []string {
	var newArray []string
	for _, v := range response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "province" {
			newArray = append(newArray, v.Name)
		}
	}
	return newArray
}

// Area returns area name for first address
func (response *YaGeoResponse) Area() string {
	str := ""
	for _, v := range response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "area" {
			str = v.Name
		}
	}
	return str
}

// Locality return name of city or another place type (for first address)
func (response *YaGeoResponse) Locality() string {
	str := ""
	for _, v := range response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "locality" {
			str = v.Name
		}
	}
	return str
}

// Street return name of street for first address
func (response *YaGeoResponse) Street() string {
	str := ""
	for _, v := range response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "street" {
			str = v.Name
		}
	}
	return str
}

// House returns house designation for first address
func (response *YaGeoResponse) House() string {
	str := ""
	for _, v := range response.Response.ObjectCollection.Members[0].GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "house" {
			str = v.Name
		}
	}
	return str
}