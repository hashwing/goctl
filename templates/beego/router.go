//pkg/server/routers/router_v1.go
package routers

import (
    "github.com/astaxie/beego"

	"{{ .Mod }}/core"
	"{{ .Mod }}/pkg/server/apis/base"
	"{{ .Mod }}/pkg/server/apis/v1"
	"{{ .Mod }}/pkg/server/middlewares"
)

func NewV1(apiMgr *core.APIManager) {
	basec := base.BaseController{
		Mgr: apiMgr,
	}
	loginc := &v1.LoginController{
		BaseController: basec,
	}
	testc := &v1.TestController{
		BaseController: basec,
	}
	health := beego.NewNamespace("/health",
		beego.NSRouter("/", &basec, "get:Health"),
	)

	public := beego.NewNamespace("/public/v1",
		beego.NSRouter("/", loginc, "post:Login"),
	)

	api := beego.NewNamespace("/api/v1",
		beego.NSRouter("/", testc, "get:Get"),
	)

	beego.AddNamespace(health, public, api)
	beego.InsertFilter("/api/*", beego.BeforeRouter,middlewares.JWTAuth(apiMgr.Auth))
}
