//pkg/auth/jwt.go
package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"{{ .Mod }}/core"

	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type TokenClaims struct {
	core.TokenInfo
	jwt.StandardClaims
}

type jwtAuth struct {
	cfg *core.Config
}

func New(cfg *core.Config) core.Auth {
	return &jwtAuth{cfg}
}

func (a *jwtAuth) CreateToken(info core.TokenInfo) string {
	expireToken := time.Now().Add(time.Second * time.Duration(a.cfg.Server.Token.Expiration)).Unix()
	claims := TokenClaims{
		info,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "demo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(a.cfg.Server.Token.Secret))
	return signedToken
}

func (a *jwtAuth) GetTokenInfo(tokenStr string) (core.TokenInfo, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(a.cfg.Server.Token.Secret), nil
	})
	if err != nil {
		return core.TokenInfo{}, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims.TokenInfo, nil
	}

	return core.TokenInfo{}, errors.New("must auth")

}

func (a *jwtAuth) ParseFromRequestToken(req *http.Request) (core.TokenInfo, error) {
	tokenStr := strings.TrimPrefix(req.Header.Get("Authorization"), "Bearer ")
	if tokenStr == "" {
		tokenStr = req.FormValue("token")
	}
	info, err := a.GetTokenInfo(tokenStr)
	return info, err
}

func (a *jwtAuth) JwtAuthFilter(ctx *context.Context) {
	if ctx.Request.FormValue("token") != "" {
		info, err := a.GetTokenInfo(ctx.Request.FormValue("token"))
		if err == nil {
			ctx.Input.SetData("token_info", info)
			return
		}
	}
	token, err := request.ParseFromRequestWithClaims(ctx.Request,
		request.AuthorizationHeaderExtractor,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(a.cfg.Server.Token.Secret), nil
		})
	if err != nil || token.Claims.(*TokenClaims).UserID == "" {
		ctx.Output.Status = 401
		ctx.Output.JSON(err, false, false)
		return
	}

	ctx.Input.SetData("token_info", token.Claims.(*TokenClaims).TokenInfo)
}

