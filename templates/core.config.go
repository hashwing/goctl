//core/config.go
package core

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port  int `envconfig:"SERVER_PORT"`
	Token TokenCfg
}

type TokenCfg struct {
	Secret     string `envconfig:"SERVER_TOKEN_SECRET"`
	Expiration int64  `envconfig:"SERVER_EXPIRATION"`
}
