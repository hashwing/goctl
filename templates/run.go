//cmd/{{.App}}/command/run.go
package command

import (
	"os"

	"{{ .Mod }}/core"
	"{{ .Mod }}/pkg/auth"
	"{{ .Mod }}/pkg/config"
	"{{ .Mod }}/pkg/server"
	"github.com/sirupsen/logrus"
	{{- if .EnableMongo }}
	"{{ .Mod }}/pkg/store/mongo"
	{{- end }}
)

func run() {
	cfg, err := config.Load(cfgFile)
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}

	{{- if .EnableMongo }}
	db,err:=mongo.New(cfg)
	if err!=nil{
		logrus.Error(err)
		os.Exit(-1)
	}
	store:=&core.Store{db}
	{{else}}
	store:=&core.Store{}
	{{- end }}
	
	apiMgr := &core.APIManager{
		Config: cfg,
		Auth:   auth.New(cfg),
		Store: store,
	}
	server.Start(apiMgr)
}