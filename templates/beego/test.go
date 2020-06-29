//pkg/server/apis/v1/test.go
package v1

import (
	"{{ .Mod }}/pkg/server/apis/base"
)

type TestController struct {
	base.BaseController
}

func (c *LoginController) Get() {
	c.JSONResult(map[string]string{"content": "hello"}, 200)
}


