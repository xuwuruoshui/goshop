package middleware

import (
	"account/jwt_op"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" || len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "认证失败，需要登录",
			})
			c.Abort()
			return
		}

		j := jwt_op.NewJWT()
		parseToken, err := j.ParseToken(token)
		if err != nil {
			if err.Error() == jwt_op.TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": jwt_op.TokenExpired,
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "认证失败，需要登录",
			})
			c.Abort()
			return
		}
		c.Set("claims", parseToken)
		c.Next()
	}
}
