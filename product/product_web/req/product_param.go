package req

type ProductReq struct {
	Id int32 `json:"id"`
	Name string `json:"name" binding:"required,min=2,max=32"`
	SN string `json:"sn" binding:"required,min=2,lt=16"`
	Stocks int32 `json:"stocks" binding:"required,min=1"`
	CategoryId int32 `json:"categoryId" binding:"required"`
	Price float32 `json:"price" binding:"required,min=0"`
	RealPrice float32 `json:"realPrice" binding:"required,min=0"`
	ShortDesc string `json:"shortDesc" binding:"required,min=3"`
	Desc string `json:"desc" binding:"required,min=3"`
	Images []string `json:"images" binding:"required,min=1"`
	DescImages []string `json:"descImages" binding:"required,min=1"`
	CoverImage string `json:"coverImage" binding:"required,url"`
	BrandId int32 `json:"brandId" binding:"required"`
	IsNew bool `json:"isNew" binding:"required"`
	IsPop bool `json:"isPop" binding:"required"`
	Selling bool `json:"selling" binding:"required"`
	FavNum int32 `json:"favNum" binding:"required"`
	SoldNum int32 `json:"soldNum" binding:"required"`
}
