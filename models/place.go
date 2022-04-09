package models

type Place struct {
	Id           uint    `json:"id"`
	CityName     string  `json:"city_name"`
	DistrictName string  `json:"district_name"`
	VillageName  string  `json:"village_name"`
	CategoryId   uint    `json:"category_id"`
	Name         string  `json:"name"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
}

func GenerateCityPlaces(cities *[]Region) {
	// places := []Place{}

	// // city level places
	// for i := range (*cities) {
	// 	cityName := (*cities)[i].Name
	// 	cityCategories := categories["city"]

	// 	for _, c := range cityCategories {
	// 		for j := 0; j < int(c.Quantity); j++ {

	// 		}
	// 	}
	// }
}
