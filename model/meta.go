package model

// Meta model for
type Meta struct {
	ID       uint   `gorm:"primary_key"`
	Key      string `gorm:"type:char(10)"`
	Value    int    `gorm:"default:0"`
	Addition string `gorm:"type:char(5)"`
}
