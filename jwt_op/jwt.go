package jwt_op

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"goshop/conf"
	"goshop/log"
	"time"
)

var (
	TokenExpired = errors.New("token已过期")
	TokenNotValidYet = errors.New("token不再有效")
	TokenMalformed = errors.New("token非法")
	TokenInvalid = errors.New("token无效")
)


type CustomClaims struct {
	jwt.RegisteredClaims
	Id int32
	NickName string
	// 权限
	AuthorityId int32
}

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT{
	return &JWT{SigningKey: []byte(conf.AppConf.JWTConfig.SigningKey)}
}

func (j *JWT)GenerateJWT(claims CustomClaims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(j.SigningKey)
	if err != nil {
		log.Logger.Error("生成JWT错误:"+err.Error())
		return "",err
	}
	return tokenStr,nil
}

func (j *JWT) ParseToken(tokenStr string)(*CustomClaims,error){
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	// 判断各种解析错误
	if err!=nil{
		if result,ok := err.(jwt.ValidationError);ok{
			switch result.Errors {
			case  jwt.ValidationErrorMalformed:
				return nil,TokenMalformed
			case  jwt.ValidationErrorExpired:
				return nil,TokenExpired
			case jwt.ValidationErrorNotValidYet:
				return nil,TokenNotValidYet
			default:
				return nil,TokenInvalid
			}
		}
	}

	if token!=nil{
		if claims,ok := token.Claims.(*CustomClaims);ok&&token.Valid{
			return claims,nil
		}
		return nil,TokenInvalid
	}

	return nil,TokenInvalid
}

// 刷新token
func (j *JWT) RefreshToken(tokenStr string)(string,error){
	jwt.TimeFunc = func() time.Time{
		return time.Unix(0,0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return "",err
	}

	// 合法就加一个小时
	if claims,ok := token.Claims.(*CustomClaims);ok && token.Valid{
		jwt.TimeFunc = time.Now
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24*time.Hour))
		return j.GenerateJWT(*claims)
	}

	return "",TokenInvalid
}
