package models

import (
	"fmt"
	"strings"
)

type (
	Location struct {
		Latitude  float64
		Longitude float64
	}
)

type (
	Region struct {
		Code     string
		Name     string
		Level    string
		Location Location
	}
)

func formatCity(cityName string) string {
	cityName = strings.Replace(cityName, "KAB.", "KABUPATEN", 1)
	cityName = strings.Title(strings.ToLower(cityName))
	return cityName
}

func formatDistrict(districtName string) string {
	return "Kecamatan " + strings.Title(strings.ToLower(districtName))
}

func formatVillage(vilageName, code string) string {
	subs := strings.Split(code, ".")
	if subs[3][0] == '1' {
		vilageName = "Kelurahan " + strings.Title(strings.ToLower(vilageName))
	} else {
		vilageName = "Desa " + strings.Title(strings.ToLower(vilageName))
	}
	return vilageName
}

func IsRegionValid(r *RegionRequest) bool {
	if r.Latitude == 0 || r.Longitude == 0 {
		return false
	}
	return true
}

func GetUpperRegionCOde(code string) string {
	i := strings.LastIndex(code, ".")
	return code[:i]
}

func SeedCitiesWithRequestData(r *Request) *[]Region {
	cities := []Region{}
	for i := range r.Data {
		if !IsRegionValid(&r.Data[i]) {
			continue
		}

		code := fmt.Sprintf("%.2f", r.Data[i].CityCode.(float64))
		name := formatCity(r.Data[i].CityName)

		cities = append(cities, Region{
			Code:  code,
			Name:  name,
			Level: "Kota",
			Location: Location{
				Longitude: r.Data[i].Longitude,
				Latitude:  r.Data[i].Latitude,
			},
		},
		)
	}
	return &cities
}

func SeedDistrictsWithRequestData(r *Request) *map[string][]*Region {
	cityCodeDistricts := map[string][]*Region{}

	for i := range r.Data {
		if !IsRegionValid(&r.Data[i]) {
			continue
		}

		code := r.Data[i].DistrictCode
		name := formatDistrict(r.Data[i].DistrictName)
		cityCode := GetUpperRegionCOde(code)

		distict := &Region{
			Code:  code,
			Name:  name,
			Level: "Kecamatan",
			Location: Location{
				Longitude: r.Data[i].Longitude,
				Latitude:  r.Data[i].Latitude,
			},
		}

		if _, ok := cityCodeDistricts[cityCode]; ok {
			cityCodeDistricts[cityCode] = append(cityCodeDistricts[cityCode], distict)
		} else {
			cityCodeDistricts[cityCode] = []*Region{distict}
		}
	}

	return &cityCodeDistricts
}

func SeedVillagesWithRequestData(r *Request) *map[string][]*Region {
	districtCodeVillages := map[string][]*Region{}

	for i := range r.Data {
		if !IsRegionValid(&r.Data[i]) {
			continue
		}

		code := r.Data[i].VilageCode
		name := formatVillage(r.Data[i].VilageName, code)
		district := GetUpperRegionCOde(code)

		village := &Region{
			Code:  code,
			Name:  name,
			Level: "Kelurahan",
			Location: Location{
				Longitude: r.Data[i].Longitude,
				Latitude:  r.Data[i].Latitude,
			},
		}

		if _, ok := districtCodeVillages[district]; ok {
			districtCodeVillages[district] = append(districtCodeVillages[district], village)
		} else {
			districtCodeVillages[district] = []*Region{village}
		}
	}

	return &districtCodeVillages
}
