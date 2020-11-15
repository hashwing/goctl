//pkg/server/server.go
package server

import (
	"fmt"
	"net/http"

	"{{ .Mod }}/core"
	"{{ .Mod }}/pkg/server/routers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Start(mgr *core.APIManager) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	routers.NewV1(r, mgr)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", mgr.Config.Server.Port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	logrus.Infof("Listen in :%d", mgr.Config.Server.Port)
	err:=s.ListenAndServe()
	if err!=nil{
		logrus.Errorf("Listen in :%d error %v", mgr.Config.Server.Port, err)
	}
}
