package models

const (
	CityLevel     uint = 1
	DistrictLevel uint = 2
	VillageLevel  uint = 3
)

type (
	RegionRequest struct {
		CityCode     string `json:"kemendagri_kota_kode"`
		DistrictCode string `json:"kemendagri_kecamatan_kode"`
		VilageCode   string `json:"kemendagri_kelurahan_kode"`

		CityName     string `json:"kemendagri_kota_nama"`
		DistrictName string `json:"kemendagri_kecamatan_nama"`
		VilageName   string `json:"kemendagri_kelurahan_nama"`

		Latitude  float64 `json:"latidute"`
		Longitude float64 `json:"longitude"`
	}

	Request struct {
		Data []RegionRequest `json:"data"`
	}
)

type (
	Location struct {
		Latitude  float64
		Longitude float64
	}
)

type (
	City struct {
		Code     string
		Name     string
		Level    uint
		Location Location
	}

	District struct {
		Code     string
		Name     string
		Level    uint
		Location Location
	}

	Village struct {
		Code     string
		Name     string
		Level    uint
		Location Location
	}
)

type (
	Cities    []City
	Districts []District
	Villages  []Village
)

func SeedCitiesWithRequestData(c *[]City, r *Request) {
	for i := range r.Data {
		*c = append(*c, City{
			Code:  r.Data[i].CityCode,
			Name:  r.Data[i].CityName,
			Level: CityLevel,
			Location: Location{
				Latitude:  r.Data[i].Latitude,
				Longitude: r.Data[i].Longitude},
		},
		)
	}
}
func SeedDistrictsWithRequestData(d *[]District, r *Request) {
	for i := range r.Data {
		*d = append(*d, District{
			Code:  r.Data[i].DistrictCode,
			Name:  r.Data[i].DistrictName,
			Level: DistrictLevel,
			Location: Location{
				Latitude:  r.Data[i].Latitude,
				Longitude: r.Data[i].Longitude},
		},
		)
	}
}

func SeedVilagesWithRequestData(v *[]Village, r *Request) {
	for i := range r.Data {
		*v = append(*v, Village{
			Code:  r.Data[i].VilageCode,
			Name:  r.Data[i].VilageName,
			Level: VillageLevel,
			Location: Location{
				Latitude:  r.Data[i].Latitude,
				Longitude: r.Data[i].Longitude},
		},
		)
	}
}
