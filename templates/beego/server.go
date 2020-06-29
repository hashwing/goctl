//pkg/server/server.go
package server

import (
	"{{ .Mod }}/core"
    "{{ .Mod }}/pkg/server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func Start(apimgr *core.APIManager) {
	routers.NewV1(apimgr)
	beego.BConfig.AppName = "{{ .App }}"
	beego.BConfig.RunMode = beego.PROD
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.Log.FileLineNum = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFExpire = 3600
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},
	}))

	beego.BConfig.Listen.HTTPPort = apimgr.Config.Server.Port
	beego.Run()
}