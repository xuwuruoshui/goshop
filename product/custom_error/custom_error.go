package custom_error

const (
	AccountNotFound = "账户不存在"
	AccountExists = "账户已存在"
	InternalError = "服务端错误"
	SaltError = "盐值为空"
	GenCaptchaError = "验证码生成错误"
	GenCaptchaBase64Error = "验证码Base64生成错误"
)
