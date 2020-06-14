package model

// ICan model for I can s
type ICan struct {
	ID    uint   `gorm:"primary_key"`
	Icon  string `gorm:"type:varchar(40)"`
	Name  string `gorm:"type:varchar(40)"`
	Color string `gorm:"type:varchar(21)"`
}
