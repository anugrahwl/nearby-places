package models

import (
	"fmt"
	"log"
	"strings"

	"github.com/umahmood/haversine"
)

const (
	RADIUS_DISTRIBUTION_CITY     float64 = 5000
	RADIUS_DISTRIBUTION_DISTRICT float64 = 3000
	RADIUS_DISTRIBUTION_VILLAGE  float64 = 1500
)

type Place struct {
	ID           uint    `json:"id"`
	CityName     string  `json:"city_name"`
	DistrictName string  `json:"district_name,omitempty"`
	VillageName  string  `json:"village_name,omitempty"`
	CategoryID   uint    `json:"category_id"`
	Name         string  `json:"name"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
}

type (
	CityPlace struct {
		Location Location
		Name     string
		Places   []Place
	}
)

var placeID uint = 1

func fmtRegionName(name string, categoryID uint) string {
	switch categoryID {
	case 5, 6, 8:
		return strings.Split(name, " ")[1]
	}
	return name
}

func generateCityPlaces(city *Region) []Place {
	places := []Place{}
	categories := Categories["city"]

	cityName := city.Name

	for _, c := range categories {
		perCategoryID := 1

		for j := 0; j < int(c.Quantity); j++ {
			name := fmt.Sprintf("%s %s %d", c.Name, fmtRegionName(cityName, c.ID), perCategoryID)
			placeLoc := GenerateRandomLocation(city.Location, RADIUS_DISTRIBUTION_CITY)
			place := Place{
				ID:         placeID,
				CityName:   cityName,
				CategoryID: c.ID,
				Name:       name,
				Longitude:  placeLoc.Longitude,
				Latitude:   placeLoc.Latitude,
			}
			places = append(places, place)
			perCategoryID++
			placeID++
		}
	}
	return places
}

func generateDistrictPlaces(district *Region, city *Region) []Place {
	places := []Place{}
	categories := Categories["district"]

	cityName := city.Name
	districtName := district.Name

	for _, c := range categories {
		perCategoryID := 1

		for j := 0; j < int(c.Quantity); j++ {
			name := fmt.Sprintf("%s %s %d", c.Name, fmtRegionName(districtName, c.ID), perCategoryID)
			placeLoc := GenerateRandomLocation(district.Location, RADIUS_DISTRIBUTION_DISTRICT)
			place := Place{
				ID:           placeID,
				CityName:     cityName,
				DistrictName: districtName,
				CategoryID:   c.ID,
				Name:         name,
				Longitude:    placeLoc.Longitude,
				Latitude:     placeLoc.Latitude,
			}
			places = append(places, place)
			perCategoryID++
			placeID++
		}
	}
	return places
}

func generateVillagePlaces(village *Region, district *Region, city *Region) []Place {
	places := []Place{}
	categories := Categories["village"]

	cityName := city.Name
	districtName := district.Name
	villageName := village.Name

	for _, c := range categories {
		perCategoryID := 1

		for j := 0; j < int(c.Quantity); j++ {
			name := fmt.Sprintf("%s %s %d", c.Name, fmtRegionName(villageName, c.ID), perCategoryID)
			placeLoc := GenerateRandomLocation(village.Location, RADIUS_DISTRIBUTION_VILLAGE)
			place := Place{
				ID:           placeID,
				CityName:     cityName,
				DistrictName: districtName,
				VillageName:  villageName,
				CategoryID:   c.ID,
				Name:         name,
				Longitude:    placeLoc.Longitude,
				Latitude:     placeLoc.Latitude,
			}
			places = append(places, place)
			perCategoryID++
			placeID++
		}
	}
	return places
}

func GeneratePlaces(cities []Region, cityDistricts, districtVillages map[string][]Region) []CityPlace {
	batchPlace := []CityPlace{}

	for i := range cities {
		city := &cities[i]
		cityPlaces := generateCityPlaces(city)

		batchPlace = append(batchPlace, CityPlace{
			Location: city.Location,
			Name:     city.Name,
			Places:   cityPlaces,
		})

		districts := cityDistricts[city.Code]
		for j := range districts {
			district := &districts[j]
			districtPlaces := generateDistrictPlaces(district, city)

			batchPlace[i].Places = append(batchPlace[i].Places, districtPlaces...)

			villages := districtVillages[district.Code]
			for k := range villages {
				village := &villages[k]
				villagePlaces := generateVillagePlaces(village, district, city)

				batchPlace[i].Places = append(batchPlace[i].Places, villagePlaces...)
			}
		}
	}

	return batchPlace
}

func LoadAll() []CityPlace {
	citiesRequest, err := FetchData(CITY_URL)
	if err != nil {
		fmt.Println(err)
	}
	cities := SeedCitiesWithRequestData(citiesRequest)

	districtRequest, err := FetchData(DISTRICT_URL)
	if err != nil {
		fmt.Println(err)
	}
	cityDistricts := SeedDistrictsWithRequestData(districtRequest)

	villagesRequest, err := FetchData(VILLAGE_URL)
	if err != nil {
		fmt.Println(err)
	}
	districtVillages := SeedVillagesWithRequestData(villagesRequest)

	batchPlace := GeneratePlaces(cities, cityDistricts, districtVillages)
	return batchPlace
}

func GetNearbyPlaces(q WebQuery, bp []CityPlace) ([]Place, error) {
	places := []Place{}
	pinned := haversine.Coord{Lat: q.Latitude, Lon: q.Longitude}

	type CityDistance struct {
		CityPlace CityPlace
		Dist      float64
	}
	closestCities := []CityDistance{}

	maxIndex := 0
	max := 0.0

	for _, cplace := range bp {

		_, distKm := haversine.Distance(pinned, haversine.Coord{
			Lat: cplace.Location.Latitude,
			Lon: cplace.Location.Longitude,
		})

		if len(closestCities) < 5 {
			closestCities = append(closestCities, CityDistance{
				CityPlace: cplace,
				Dist:      distKm,
			})

			if max != 0.0 {
				if distKm > max {
					max = distKm
					maxIndex = len(closestCities) - 1
				}
			} else {
				max = distKm
			}
		} else {

			if distKm < max {
				closestCities[maxIndex] = CityDistance{
					CityPlace: cplace,
					Dist:      distKm,
				}

				max = distKm

				for i, p := range closestCities {
					if maxIndex == i {
						continue
					}

					if p.Dist > max {
						maxIndex = i
						max = p.Dist
					}
				}
			}
		}
	}

	log.Println("5 closest cities/regeencies to given location:")

	for _, city := range closestCities {

		fmt.Println(city.CityPlace.Name)

		for _, place := range city.CityPlace.Places {
			_, distKm := haversine.Distance(pinned, haversine.Coord{
				Lat: place.Latitude,
				Lon: place.Longitude,
			})

			if distKm <= 5 {

				if q.CategoryId != 0 {
					if place.CategoryID == q.CategoryId {
						places = append(places, place)
					}
				} else {
					places = append(places, place)
				}
			}
		}
	}
	fmt.Println()

	return places, nil
}
