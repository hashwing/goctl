//pkg/server/routers/router_v1.go
package routers

import (
	"{{ .Mod }}/core"
	"{{ .Mod }}/pkg/server/apis/base"
	v1 "{{ .Mod }}/pkg/server/apis/v1"
	"{{ .Mod }}/pkg/server/middlewares"

	"github.com/gin-gonic/gin"
)

func NewV1(r *gin.Engine, mgr *core.APIManager) {
	basec := base.BaseController{Mgr: mgr}
	testc := &v1.TestController{BaseController: basec}
	r.GET("/health", basec.Health)
	apiv1 := r.Group("/api")
	apiv1.Use(middlewares.JWTAuth(mgr.Auth))
	apiv1.GET("test", testc.Get)
}