package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	VILLAGE_URL  string = "https://satudata.jabarprov.go.id/api-backend/bigdata/diskominfo/od_kode_wilayah_dan_nama_wilayah_desa_kelurahan"
	DISTRICT_URL string = "https://satudata.jabarprov.go.id/api-backend/bigdata/diskominfo/od_16357_kode_wilayah_dan_nama_wilayah_kecamatan"
	CITY_URL     string = "https://satudata.jabarprov.go.id/api-backend/bigdata/diskominfo/od_kode_wilayah_dan_nama_wilayah_kota_kabupaten"
)

func PrintPrettyJson(jsonData *map[string]interface{}) error {
	byteData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(byteData))
	return nil
}

func Fetch(url string) models.Request {

}

func main() {
	// vilageLimit := "?limit1" // "?limit=5957"
	// districtLimit := "?limit=100" // "?limit=627"
	cityLimit := "?limit=80"
	// res, err := http.Get(VILLAGE_URL + vilageLimit)
	res, err := http.Get(CITY_URL + cityLimit)
	if err != nil {
		fmt.Println(err)
	}

	byteData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(byteData, &jsonData)

	if err != nil {
		fmt.Println(err)
	}

	PrintPrettyJson(&jsonData)
}
