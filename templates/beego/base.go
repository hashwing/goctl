//pkg/server/apis/base/base.go
package base

import (
	"context"

	"github.com/astaxie/beego"

	"{{ .Mod }}/core"
)

// BaseController the base controller
type BaseController struct {
	beego.Controller
	Mgr *core.APIManager
}

func (c *BaseController) TokenInfo() core.TokenInfo {
	return c.Ctx.Input.GetData("token_info").(core.TokenInfo)
}


func (c *BaseController) JSONResult(result interface{}, statusCode int) {
	c.Ctx.Output.Status = statusCode
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *BaseController) UserContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token_info", c.Ctx.Input.GetData("token_info"))
	return ctx
}

func (c *BaseController) Health() {
    c.JSONResult(map[string]string{"status":"ok"},200)
}