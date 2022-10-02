package model

type Brand struct {
	*BaseModel
	Name string `gorm:"type:varchar(32);not null"`
	Logo string `gorm:"type:varchar(32);not null;default:''"`
}