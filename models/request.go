package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	VILLAGE_URL  string = "https://satudata.jabarprov.go.id/api-backend/bigdata/diskominfo/od_kode_wilayah_dan_nama_wilayah_desa_kelurahan?limit=5957" // 5957
	DISTRICT_URL string = "https://satudata.jabarprov.go.id/api-backend/bigdata/diskominfo/od_16357_kode_wilayah_dan_nama_wilayah_kecamatan?limit=627" // 627
	CITY_URL     string = "https://satudata.jabarprov.go.id/api-backend/bigdata/diskominfo/od_kode_wilayah_dan_nama_wilayah_kota_kabupaten?limit=27"   // 27
)

type (
	RegionRequest struct {
		CityCode     interface{} `json:"kemendagri_kota_kode"`
		DistrictCode string      `json:"kemendagri_kecamatan_kode"`
		VilageCode   string      `json:"kemendagri_kelurahan_kode"`

		CityName     string `json:"kemendagri_kota_nama"`
		DistrictName string `json:"kemendagri_kecamatan_nama"`
		VilageName   string `json:"kemendagri_kelurahan_nama"`

		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	Request struct {
		Data []RegionRequest `json:"data"`
	}
)

func FetchData(url string) (*Request, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	byteData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := Request{}

	err = json.Unmarshal(byteData, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
