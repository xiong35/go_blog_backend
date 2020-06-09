package model

// Meta model for
type Meta struct {
	ID       uint `gorm:"primary_key"`
	Key      string
	Value    int
	Addition string
}
