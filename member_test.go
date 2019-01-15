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
		{Name: "Россия", Kind: "country"},
		{Name: "Уральский федеральный округ", Kind: "province"},
		{Name: "Челябинская область", Kind: "province"},
		{Name: "городской округ Челябинск", Kind: "area"},
		{Name: "Челябинск", Kind: "locality"},
		{Name: "улица Захаренко", Kind: "street"},
		{Name: "2", Kind: "house"},
	}
	want := &array
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get address components", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].AddressComponents(); !reflect.DeepEqual(got, want) {
			t.Errorf("YaGeoMember.AddressComponents() = %v, want %v", got, want)
		}
	})
}

func TestYaGeoMember_GetComponentsByKind(t *testing.T) {
	want := []*YaGeoAddressComponent{{"house", "2"}}
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get Components By Kind", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].GetComponentsByKind("house"); !reflect.DeepEqual(got, want) {
			t.Errorf("YaGeoMember.GetComponentsByKind() = %v, want %v", got, want)
		}
	})
}

func TestYaGeoMember_Country(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get Country", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].Country(); got != "Россия" {
			t.Errorf("YaGeoMember.Country() = %v, want %v", got, "Россия")
		}
	})
}

func TestYaGeoMember_Area(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get Area", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].Area(); got != "городской округ Челябинск" {
			t.Errorf("YaGeoMember.Area() = %v, want %v", got, "городской округ Челябинск")
		}
	})
}

func TestYaGeoMember_Province(t *testing.T) {
	want := []string{"Уральский федеральный округ", "Челябинская область"}
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get postal code", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].Province(); !reflect.DeepEqual(got, want) {
			t.Errorf("YaGeoMember.Province() = %v, want %v", got, want)
		}
	})
}

func TestYaGeoMember_Locality(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get Locality", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].Locality(); got != "Челябинск" {
			t.Errorf("YaGeoMember.Locality() = %v, want %v", got, "Челябинск")
		}
	})
}

func TestYaGeoMember_Street(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get Street", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].Street(); got != "улица Захаренко" {
			t.Errorf("YaGeoMember.Street() = %v, want %v", got, "улица Захаренко")
		}
	})
}

func TestYaGeoMember_House(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	members := *response.Members()
	t.Run("Get House", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := members[0].House(); got != "2" {
			t.Errorf("YaGeoMember.House() = %v, want %v", got, "2")
		}
	})
}

func TestYaGeoMember_RangeToMember(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response1, err1 := New(key).Find("Челябинск, Захаренко, 2")
	members1 := *response1.Members()
	response2, err2 := New(key).Find("Челябинск, Захаренко, 5")
	members2 := *response2.Members()
	t.Run("Get range to member", func(t *testing.T) {
		if err1 != nil {
			t.Errorf("YaGeoMember.RangeToMember() has error: %v", err1.Error())
		}
		if err2 != nil {
			t.Errorf("YaGeoMember.RangeToMember() has error: %v", err2.Error())
		}
		if got := members1[0].RangeToMember(&members2[0]); got != 143.70860825840776 {
			t.Errorf("YaGeoMember.RangeToMember() = %v, want %v", got, 143.70860825840776)
		}
	})
}
