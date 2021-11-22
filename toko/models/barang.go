package models

type Item struct {
	ID     int      `json:"id" gorm:"primary_key"`
	Nama   string   `json:"nama"`
	Pajake []Pajake `json:"pajak" gorm:"foreignKey:PajakeID"`
}
type Pajake struct {
	ID       int     `json:"id" gorm:"primary_key"`
	Nama     string  `json:"nama"`
	Rate     float32 `json:"rate"`
	PajakeID int     `json:"-"`
}
