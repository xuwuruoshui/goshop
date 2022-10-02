package model

type Advertise struct {
	*BaseModel
	Index int32 `gorm:"type:int;not null;default:1"`
	Image string `gorm:"type:varchar(256);not null"`
	Url string `gorm:"type:varchar(256);not null"`
	Sort int32 `gorm:"type:int;not null;default:1"`
}