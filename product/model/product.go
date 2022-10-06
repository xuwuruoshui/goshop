package model

import (
	"database/sql/driver"
	"encoding/json"
)


type Product struct {
	BaseModel
	CategoryId int32 `gorm:"type:int;not null"`
	Category *Category
	BrandId int32 `gorm:"type:int;not null"`
	Brand *Brand
	// 是否正在销售
	Selling bool `gorm:"default:false;not null"`
	// 是否包邮
	IsShipFree bool `gorm:"default:false;not null"`
	// 是否热卖
	IsPop bool `gorm:"default:false;not null"`
	// 是否新品
	IsNew bool `gorm:"default:false;not null"`

	Name string `gorm:"type:varchar(64);not null"`
	// 编号
	SN string `gorm:"type:varchar(64);not null"`
	// 收藏数
	FavNum int32 `gorm:"type:int;default:0"`
	// 销售量
	SoldNum int32 `gorm:"type:int;default:0"`
	// 价格
	Price float32 `gorm:"not null"`
	// 真实价格
	RealPrice float32 `gorm:"not null"`
	// 简述
	ShortDesc string `gorm:"type:varchar(256);not null"`
	Images MyList `gorm:"type:varchar(1024);not null"`
	DescImages MyList `gorm:"type:varchar(1024);not null"`
	CoverImage string `gorm:"type:varchar(256);not null"`
}

type MyList []string

func (myList MyList)Value()(driver.Value,error){
	return json.Marshal(myList)
}

func (myList MyList)Scan(v interface{})error{
	return json.Unmarshal(v.([]byte),&myList)
}


