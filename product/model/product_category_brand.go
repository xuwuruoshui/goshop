package model

type ProductCategoryBrand struct {
	BaseModel
	CategoryId int32 `gorm:"type:int;index:idx_category_brand"` //联合唯一索引.2个索引是一样的，那么就是联合唯一索引了
	Category   Category
	BrandId    int32 `gorm:"type:int;index:idx_category_brand"`
	Brand      Brand
}
