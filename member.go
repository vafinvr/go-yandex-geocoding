package yageocoding

import (
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
	latitude, errlat := strconv.ParseFloat(coords[1], 64)
	if errlat != nil {
		return 0, 0
	}

	longitude, errlon := strconv.ParseFloat(coords[0], 64)
	if errlon != nil {
		return 0, 0
	}

	return
}

// Latitude of member
func (member *YaGeoMember) Latitude() float64 {
	coords := strings.Split(member.GeoObject.Point.Pos, " ")
	latitude, errlat := strconv.ParseFloat(coords[1], 64)
	if errlat != nil {
		return 0
	}
	return latitude
}

// Longitude of member
func (member *YaGeoMember) Longitude() float64 {
	coords := strings.Split(member.GeoObject.Point.Pos, " ")
	longitude, errlon := strconv.ParseFloat(coords[0], 64)
	if errlon != nil {
		return 0
	}
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

func (member *YaGeoMember) GetComponentsByKind(kind string) []*YaGeoAddressComponent {
	var newArray []*YaGeoAddressComponent
	for _, val := range member.GeoObject.MetaData.Meta.Address.Components {
		if val.Kind == kind {
			newArray = append(newArray, &val)
		}
	}
	return newArray
}