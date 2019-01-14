package yageocoding

import (
	"os"
	"reflect"
	"testing"
)

func TestYaGeoMember_Coordinates(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get coordinates", func(t *testing.T) {
		gotLatitude, gotLongitude := members[0].Coordinates()
		if err != nil {
			t.Errorf("YaGeoMember.Coordinates() has error: %v", err.Error())
		}
		if gotLatitude != 55.199352 {
			t.Errorf("YaGeoMember.Coordinates() gotLatitude = %v, want %v", gotLatitude, 55.199352)
		}
		if gotLongitude != 61.315103 {
			t.Errorf("YaGeoMember.Coordinates() gotLongitude = %v, want %v", gotLongitude, 61.315103)
		}
	})
}

func TestYaGeoMember_Latitude(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get latitude", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.Latitude() has error: %v", err.Error())
		}
		if got := members[0].Latitude(); got != 55.199352 {
			t.Errorf("YaGeoMember.Latitude() = %v, want %v", got, 55.199352)
		}
	})
}

func TestYaGeoMember_Longitude(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get Longitude", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.Longitude() has error: %v", err.Error())
		}
		if got := members[0].Longitude(); got != 61.315103 {
			t.Errorf("YaGeoMember.Longitude() = %v, want %v", got, 61.315103)
		}
	})
}

func TestYaGeoMember_CountryCode(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get Country code", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].CountryCode(); got != "RU" {
			t.Errorf("YaGeoMember.CountryCode() = %v, want %v", got, "RU")
		}
	})
}

func TestYaGeoMember_PostalCode(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get postal code", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].PostalCode(); got != "454014" {
			t.Errorf("YaGeoMember.PostalCode() = %v, want %v", got, "454014")
		}
	})
}

func TestYaGeoMember_AddressComponents(t *testing.T) {
	array := []YaGeoAddressComponent{
		{Name:"Россия", Kind:"country"},
		{Name:"Уральский федеральный округ",Kind:"province"},
		{Name:"Челябинская область", Kind:"province"},
		{Name:"городской округ Челябинск", Kind:"area"},
		{Name:"Челябинск", Kind:"locality"},
		{Name:"улица Захаренко", Kind:"street"},
		{Name:"2", Kind:"house"},
	}
	want := &array
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get postal code", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].AddressComponents(); !reflect.DeepEqual(got, want) {
			t.Errorf("YaGeoMember.AddressComponents() = %v, want %v", got, want)
		}
	})
}
