//pkg/store/mongo/db.go
package mongo

import (
	"context"
	"fmt"

	"{{ .Mod }}/core"
	"{{ .Mod }}/pkg/errors"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type store struct {
	mongo *MongoClient
}

// New 初始化数据库连接
func New(cfg *core.Config) (*core.DBStore, error) {
	datasource := "mongodb://"
	if cfg.Server.Mongo.Username != "" && cfg.Server.Mongo.Password != "" {
		datasource = fmt.Sprintf("%s%s:%s@",
			datasource,
			cfg.Server.Mongo.Username,
			cfg.Server.Mongo.Password)
	}

	if cfg.Server.Mongo.Address == "" {
		return nil, errors.NewErrInvalidParam("请配置数据库服务器地址")
	}

	if cfg.Server.Mongo.Database == "" {
		return nil, errors.NewErrInvalidParam("请配置数据库名")
	}

	datasource = fmt.Sprintf("%s%s:%d/%s?authSource=admin",
		datasource,
		cfg.Server.Mongo.Address,
		cfg.Server.Mongo.Port,
		cfg.Server.Mongo.Database)

	logrus.Debug(datasource)

	dbc, err := NewMongo(datasource, cfg.Server.Mongo.Database)
	if err != nil {
		return nil, err
	}
	s := &store{dbc}
	//注册你的store
	_ = s
	
	store := &core.DBStore{
	}
	return store, nil
}

const (
	retryTimes = 3
)

// MongoClient Mongodb 客户端
type MongoClient struct {
	Conn     *mongo.Client
	database string
}

// NewMongo Initial Mongo Client
func NewMongo(dataSourceName string, db string) (*MongoClient, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(dataSourceName))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	c := &MongoClient{
		Conn:     client,
		database: db,
	}

	return c, nil
}

// GetClient 获取Mongo连接
func (c *MongoClient) GetClient() (*mongo.Database, error) {

	if err := c.Conn.Ping(context.Background(), readpref.Primary()); err != nil {
		logrus.Error(err)
		if err := c.checkClientConnect(); err != nil {
			logrus.Error(err)
			return nil, err
		}
		return c.Conn.Database(c.database), nil
	}

	return c.Conn.Database(c.database), nil
}

// checkEngineConnect 检查Mongo连接是否断开
func (c *MongoClient) checkClientConnect() error {
	for i := 0; i < retryTimes; i++ {
		if i == retryTimes {
			return errors.NewErrInternal()
		}

		if err := c.Conn.Ping(context.Background(), readpref.Primary()); err != nil {
			logrus.Error(err)
			continue
		}
		break
	}
	return nil
}
