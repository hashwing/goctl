//cmd/{{.App}}/commond/run.go
package command

import (
	"os"

	"{{ .Mod }}/core"
	"{{ .Mod }}/pkg/auth"
	"{{ .Mod }}/pkg/config"
	"{{ .Mod }}/pkg/server"
	"github.com/sirupsen/logrus"
)

func run() {
	cfg, err := config.Load(cfgFile)
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}

	apiMgr := &core.APIManager{
		Config: cfg,
		Auth:   auth.New(cfg),
	}
	server.Start(apiMgr)
}