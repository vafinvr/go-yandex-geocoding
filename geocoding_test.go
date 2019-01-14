package yageocoding

import (
	"os"
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
