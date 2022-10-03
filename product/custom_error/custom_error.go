package custom_error

const (
	DelBrandFailed = "品牌删除失败"
	BrandExist     = "品牌已存在"
	BrandNotExist = "品牌不存在"

	DeleteAdvertiseFailed = "广告删除失败"
	AdvertiseNotExist = "广告不存在"

	CategoryNotExist = "分类不存在"
	CategoryMarshalFailed = "分类序列化失败"
	InternalError = "服务端错误"

)
