//pkg/server/apis/base/base.go
package base

import (
	"github.com/gin-gonic/gin"
	"{{ .Mod }}/core"
)

//BaseController 基础controller
type BaseController struct {
	Mgr *core.APIManager
}

//Health 健康状态
func (ctrl *BaseController) Health(c *gin.Context) {
	c.JSON(200, map[string]interface{}{"status": "ok"})
}
