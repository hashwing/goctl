//pkg/server/apis/v1/login.go
package v1

import (
	"{{ .Mod }}/core"
	"{{ .Mod }}/pkg/server/apis/base"
    typev1 "{{ .Mod }}/pkg/server/apis/v1/type"
)

type LoginController struct {
	base.BaseController
}

func (c *LoginController) Login() {
	info := core.TokenInfo{
		UserID: "",
	}
	o := typev1.LoginOutput{
		Token: c.Mgr.Auth.CreateToken(info),
	}
	c.JSONResult(o, 200)
}


