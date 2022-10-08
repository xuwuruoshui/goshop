package model

type Stock struct {
	BaseModel
	ProductId int32 `gorm:"type:int;index"`
	Num int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"`
}

