package model

type Category struct {
	*BaseModel
	Name string `gorm:"type:varchar(32);not null"`
	ParentCategoryId int32
	ParentCategory *Category
	SubCategory []*Category `gorm:"foreignKey:ParentCategoryId;references:Id"`

}