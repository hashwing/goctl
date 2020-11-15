//pkg/server/middlewares/jwt.go
package middlewares

import (
	"{{ .Mod }}/core"
	"github.com/astaxie/beego/context"
)

func JWTAuth(auth core.Auth) func(ctx *context.Context) {
	return func (ctx *context.Context) {		
		info, err := auth.ParseFromRequestToken(ctx.Request)
		if err != nil {
			ctx.Output.Status = 401
			ctx.Output.JSON(map[string]string{"msg": "must auth"}, false, false)
			return
		}
		ctx.Input.SetData("token_info", info)
	}
}