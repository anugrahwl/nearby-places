package models

type Category struct {
	Id       uint
	Name     string
	Quantity uint
}

var Categories = map[string][]Category{
	"city": {
		{1, "Rumah Sakit", 3},
		{2, "Sekolah Menengah Atas", 20},
		{3, "Kantor Pemerintah", 1},
	},
	"district": {
		{4, "Rumah Sakit", 5},
		{5, "Sekolah Menengah Pertama", 3},
		{6, "Kantor Pemerintah", 1},
	},
	"village": {
		{7, "Sekolah Dasar", 5},
		{8, "Tempat Ibadah", 20},
		{9, "Kantor Pemerintah", 1},
	},
}
