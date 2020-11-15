//pkg/server/middlewares/jwt.go
package middlewares

import (
	"{{ .Mod }}/core"
	"github.com/gin-gonic/gin"
)

func JWTAuth(auth core.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {		
		info, err := auth.ParseFromRequestToken(c.Request)
		if err != nil {
			c.JSON(401, gin.H{
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("token_info", info)
	}
}