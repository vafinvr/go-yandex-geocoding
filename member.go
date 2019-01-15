package yageocoding

import (
	"math"
	"strconv"
	"strings"
)

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
		Kind      string       `json:"kind"`
		Text      string       `json:"text"`
		Precision string       `json:"precision"`
		Address   YaGeoAddress `json:"Address"`
	} `json:"GeocoderMetaData"`
}

// Coordinates returns Latitude and Longitude of member
func (member *YaGeoMember) Coordinates() (latitude float64, longitude float64) {
	coords := strings.Split(member.GeoObject.Point.Pos, " ")
	latitude, _ = strconv.ParseFloat(coords[1], 64)
	longitude, _ = strconv.ParseFloat(coords[0], 64)
	return
}

// Latitude of member
func (member *YaGeoMember) Latitude() float64 {
	coords := strings.Split(member.GeoObject.Point.Pos, " ")
	latitude, _ := strconv.ParseFloat(coords[1], 64)
	return latitude
}

// Longitude of member
func (member *YaGeoMember) Longitude() float64 {
	coords := strings.Split(member.GeoObject.Point.Pos, " ")
	longitude, _ := strconv.ParseFloat(coords[0], 64)
	return longitude
}

// CountryCode returns country code of member
func (member *YaGeoMember) CountryCode() string {
	return member.GeoObject.MetaData.Meta.Address.CountryCode
}

// PostalCode returns postal code of member
func (member *YaGeoMember) PostalCode() string {
	return member.GeoObject.MetaData.Meta.Address.PostalCode
}

// AddressComponents returns array of address components
func (member *YaGeoMember) AddressComponents() *[]YaGeoAddressComponent {
	return &member.GeoObject.MetaData.Meta.Address.Components
}

// GetComponentsByKind returns address components sorted by kind
func (member *YaGeoMember) GetComponentsByKind(kind string) []*YaGeoAddressComponent {
	var newArray []*YaGeoAddressComponent
	for _, val := range member.GeoObject.MetaData.Meta.Address.Components {
		if val.Kind == kind {
			newArray = append(newArray, &val)
		}
	}
	return newArray
}

// Country returns country name
func (member *YaGeoMember) Country() string {
	str := ""
	for _, v := range member.GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "country" {
			str = v.Name
		}
	}
	return str
}

// Province returns array of province names
func (member *YaGeoMember) Province() []string {
	var newArray []string
	for _, v := range member.GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "province" {
			newArray = append(newArray, v.Name)
		}
	}
	return newArray
}

// Area returns area name
func (member *YaGeoMember) Area() string {
	str := ""
	for _, v := range member.GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "area" {
			str = v.Name
		}
	}
	return str
}

// Locality return name of city or another place type
func (member *YaGeoMember) Locality() string {
	str := ""
	for _, v := range member.GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "locality" {
			str = v.Name
		}
	}
	return str
}

// Street return name of street
func (member *YaGeoMember) Street() string {
	str := ""
	for _, v := range member.GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "street" {
			str = v.Name
		}
	}
	return str
}

// RangeToMember returns range in meters to another member
func (member *YaGeoMember) RangeToMember(target *YaGeoMember) float64 {
	earthRadius := float64(6371000) // Earth's radius in meters
	difLat := deg2rad(member.Latitude() - target.Latitude())
	difLng := deg2rad(member.Longitude() - target.Longitude())
	a := math.Sin(difLat/2)*math.Sin(difLat/2) +
		math.Cos(target.Latitude())*math.Cos(member.Latitude())*
			math.Sin(difLng/2)*math.Sin(difLng/2)
	c := 2 * math.Asin(math.Sqrt(a))
	distance := earthRadius * c

	return distance
}

// House returns house designation
func (member *YaGeoMember) House() string {
	str := ""
	for _, v := range member.GeoObject.MetaData.Meta.Address.Components {
		if v.Kind == "house" {
			str = v.Name
		}
	}
	return str
}
