//core/api.go
package core

import (
	"net/http"

	"github.com/astaxie/beego/context"
)

type APIManager struct {
	Config *Config
	Auth   Auth
	Store  *Store
}

type TokenInfo struct {
	UserID string `json:"user_id"`
}

type Auth interface {
	JwtAuthFilter(ctx *context.Context)
	CreateToken(info TokenInfo) string
	ParseFromRequestToken(req *http.Request) (TokenInfo, error)
}
