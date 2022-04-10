package models

type Category struct {
	ID       uint
	Name     string
	Quantity uint
}

var Categories = map[string][]Category{
	// 1 2 3 4 7 9
	// 5 6 8
	"city": {
		{1, "Kantor Pemerintah", 1},
		{2, "Rumah Sakit", 3},
		{3, "Sekolah Menengah Atas", 20},
	},
	"district": {
		{4, "Kantor Pemerintah", 1},
		{5, "Sekolah Menengah Pertama", 3},
		{6, "Puskesmas", 5},
	},
	"village": {
		{7, "Kantor Pemerintah", 1},
		{8, "Sekolah Dasar", 5},
		{9, "Tempat Ibadah", 20},
	},
}
