//pkg/server/apis/v1/test.go
package v1

import (
	"{{ .Mod }}/pkg/server/apis/base"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	base.BaseController
}

func (ctrl *TestController) Get(c *gin.Context) {
	c.JSON(200, map[string]string{"content": "hello"})
}