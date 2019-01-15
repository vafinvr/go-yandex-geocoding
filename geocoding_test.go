package yageocoding

import (
	"os"
	"reflect"
	"testing"
)

func TestYaGeoInstance_RangeBtw(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	type fields struct {
		Key string
	}
	type args struct {
		address1 string
		address2 string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "Range test",
			fields:  fields{key},
			args:    args{"Челябинск, Захаренко, 2", "Челябинск, Захаренко, 5"},
			want:    143.70860825840776,
			wantErr: false,
		},
		{
			name:    "Range test error 1",
			fields:  fields{key},
			args:    args{"SomePlaceLOL_404_notFound", "Челябинск, Захаренко, 5"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Range test error 2",
			fields:  fields{key},
			args:    args{"Челябинск, Захаренко, 2", "SomePlaceLOL_404_notFound_GiveMeErrorPlease"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ygi := &YaGeoInstance{
				Key: tt.fields.Key,
			}
			got, err := ygi.RangeBtw(tt.args.address1, tt.args.address2)
			if (err != nil) != tt.wantErr {
				t.Errorf("YaGeoInstance.RangeBtw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("YaGeoInstance.RangeBtw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYaGeoResponse_Longitude(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	t.Run("Get Longitude", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoResponse.Longitude() has error: %v", err.Error())
		}
		if got := response.Longitude(); got != 61.315103 {
			t.Errorf("YaGeoResponse.Longitude() = %v, want %v", got, 61.315103)
		}
	})
}

func TestYaGeoResponse_Latitude(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	t.Run("Get Latitude", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoResponse.Latitude() has error: %v", err.Error())
		}
		if got := response.Latitude(); got != 55.199352 {
			t.Errorf("YaGeoResponse.Latitude() = %v, want %v", got, 55.199352)
		}
	})
}

func TestYaGeoResponse_Coordinates(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	gotLatitude, gotLongitude := response.Coordinates()
	t.Run("Get Longitude", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoResponse.Coordinates() has error: %v", err.Error())
		}
		if gotLatitude != 55.199352 {
			t.Errorf("YaGeoResponse.Coordinates() gotLatitude = %v, want %v", gotLatitude, 55.199352)
		}
		if gotLongitude != 61.315103 {
			t.Errorf("YaGeoResponse.Coordinates() gotLongitude = %v, want %v", gotLongitude, 61.315103)
		}
	})
}

func TestYaGeoResponse_CountryCode(t *testing.T) {
	response, err := New(os.Getenv("YAGEO_KEY")).Find("Челябинск, Захаренко, 2")
	t.Run("Get country code", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoResponse.CountryCode() has error: %v", err.Error())
		}
		if got := response.CountryCode(); got != "RU" {
			t.Errorf("YaGeoResponse.CountryCode() = %v, want %v", got, "RU")
		}
	})
}

func TestYaGeoResponse_PostalCode(t *testing.T) {
	response, err := New(os.Getenv("YAGEO_KEY")).Find("Челябинск, Захаренко, 2")
	t.Run("Get postal code", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoResponse.CountryCode() has error: %v", err.Error())
		}
		if got := response.PostalCode(); got != "454014" {
			t.Errorf("YaGeoResponse.PostalCode() = %v, want %v", got, "454014")
		}
	})
}

func TestYaGeoResponse_AddressComponents(t *testing.T) {
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
	t.Run("Get postal code", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := response.AddressComponents(); !reflect.DeepEqual(got, want) {
			t.Errorf("YaGeoMember.AddressComponents() = %v, want %v", got, want)
		}
	})
}

func TestYaGeoResponse_RangeToResponse(t *testing.T) {
	response1, err := New(os.Getenv("YAGEO_KEY")).Find("Челябинск, Захаренко, 2")
	response2, err := New(os.Getenv("YAGEO_KEY")).Find("Челябинск, Захаренко, 5")
	t.Run("Get postal code", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoResponse.CountryCode() has error: %v", err.Error())
		}
		if got := response1.RangeToResponse(response2); got != 143.70860825840776 {
			t.Errorf("YaGeoResponse.RangeToResponse() = %v, want %v", got, 143.70860825840776)
		}
	})
}

func TestYaGeoResponse_Address(t *testing.T) {
	response, err := New(os.Getenv("YAGEO_KEY")).Find("Челябинск, Захаренко, 2")
	t.Run("Get postal code", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoResponse.CountryCode() has error: %v", err.Error())
		}
		if got := response.Address(); got != "Россия, Челябинск, улица Захаренко, 2" {
			t.Errorf("YaGeoResponse.Address() = %v, want %v", got, "Россия, Челябинск, улица Захаренко, 2")
		}
	})
}

func TestYaGeoInstance_Find(t *testing.T) {
	type fields struct {
		Key string
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Test Find() function",
			fields{os.Getenv("YAGEO_KEY")},
			args{"Челябинск, Захаренко, 2"},
			false,
		},
		{
			"Test Find() function error",
			fields{os.Getenv("YAGEO_KEY")},
			args{"SomePlaceLOL_404_notFound"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ygi := &YaGeoInstance{
				Key: tt.fields.Key,
			}
			gotResult, err := ygi.Find(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("YaGeoInstance.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult.Response.ObjectCollection.MetaData.ResponseMetaData.Request != tt.args.address {
				t.Errorf("YaGeoInstance.Find() = %v, want %v", gotResult.Response.ObjectCollection.MetaData.ResponseMetaData.Request, tt.args.address)
			}
		})
	}
}

func TestYaGeoResponse_Country(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	t.Run("Get Country", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := response.Country(); got != "Россия" {
			t.Errorf("YaGeoMember.Country() = %v, want %v", got, "Россия")
		}
	})
}

func TestYaGeoResponse_Area(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	t.Run("Get Area", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := response.Area(); got != "городской округ Челябинск" {
			t.Errorf("YaGeoMember.Area() = %v, want %v", got, "городской округ Челябинск")
		}
	})
}

func TestYaGeoResponse_Province(t *testing.T) {
	want := []string{"Уральский федеральный округ", "Челябинская область"}
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	t.Run("Get Province", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := response.Province(); !reflect.DeepEqual(got, want) {
			t.Errorf("YaGeoMember.Province() = %v, want %v", got, want)
		}
	})
}

func TestYaGeoResponse_Locality(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	t.Run("Get Locality", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := response.Locality(); got != "Челябинск" {
			t.Errorf("YaGeoMember.Locality() = %v, want %v", got, "Челябинск")
		}
	})
}

func TestYaGeoResponse_Street(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	t.Run("Get Street", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := response.Street(); got != "улица Захаренко" {
			t.Errorf("YaGeoMember.Street() = %v, want %v", got, "улица Захаренко")
		}
	})
}

func TestYaGeoResponse_House(t *testing.T) {
	key := os.Getenv("YAGEO_KEY")
	response, err := New(key).Find("Челябинск, Захаренко, 2")
	t.Run("Get House", func(t *testing.T) {
		if err != nil {
			t.Errorf("YaGeoMember.CountryCode() has error: %v", err.Error())
		}
		if got := response.House(); got != "2" {
			t.Errorf("YaGeoMember.House() = %v, want %v", got, "2")
		}
	})
}