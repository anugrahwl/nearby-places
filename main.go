package main

import (
	"fmt"

	"github.com/anugrahwl/nearby-places/lib"
	"github.com/anugrahwl/nearby-places/models"
)

func main() {
	citiesRequest, err := models.FetchData(models.CITY_URL)
	if err != nil {
		fmt.Println(err)
	}
	districtRequest, err := models.FetchData(models.DISTRICT_URL)
	if err != nil {
		fmt.Println(err)
	}
	villagesRequest, err := models.FetchData(models.VILLAGE_URL)
	if err != nil {
		fmt.Println(err)
	}

	cities := models.SeedCitiesWithRequestData(citiesRequest)
	cityCodeDistricts := models.SeedDistrictsWithRequestData(districtRequest)
	districtCodeVillages := models.SeedVillagesWithRequestData(villagesRequest)

	lib.PrintPrettyJson(cities)
	lib.PrintPrettyJson((*cityCodeDistricts)["32.01"])
	lib.PrintPrettyJson((*districtCodeVillages)["32.01.01"])
}
