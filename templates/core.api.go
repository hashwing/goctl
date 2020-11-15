//core/api.go
package core

import (
	"net/http"
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
	CreateToken(info TokenInfo) string
	ParseFromRequestToken(req *http.Request) (TokenInfo, error)
	GetTokenInfo(tokenStr string) (TokenInfo, error)
}
