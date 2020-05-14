//core/config.go
package core

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port  int `envconfig:"SERVER_PORT"`
	Token TokenCfg
	{{- if .EnableMongo }}
	Mongo MongoConfig
	{{- end }} 
}

type TokenCfg struct {
	Secret     string `envconfig:"SERVER_TOKEN_SECRET"`
	Expiration int64  `envconfig:"SERVER_EXPIRATION"`
}

{{- if .EnableMongo }}
// MongoConfig MongoDB数据库连接参数
type MongoConfig struct {
	Address  string `envconfig:"ALGALON_SERVER_MONGO_ADDRESS" default:"localhost"`
	Port     uint   `envconfig:"ALGALON_SERVER_MONGO_PORT" default:"27017"`
	Username string `envconfig:"ALGALON_SERVER_MONGO_USERNAME" default:"root"`
	Password string `envconfig:"ALGALON_SERVER_MONGO_PASSWD" default:"sunrunvas"`
	Database string `envconfig:"ALGALON_SERVER_MONGO_DATABASE" default:"{{ .App }}"`
}

{{- end }}