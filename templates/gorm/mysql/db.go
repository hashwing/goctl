//pkg/store/mysql/db.go
package mysql

import (
	"database/sql"

	"{{ .Mod }}/core"
	"{{ .Mod }}/pkg/store/migrate/mysql"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type store struct {
	db *gorm.DB
}

// New 初始化数据库连接
func New(cfg *core.Config) (*core.DBStore, error) {
	datasource := cfg.Server.Mysql.Username + ":" + cfg.Server.Mysql.Password + "@tcp(" + cfg.Server.Mysql.Address + ")/?charset=utf8&parseTime=true"
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	db.Exec("CREATE DATABASE " + cfg.Server.Mysql.Database + " DEFAULT CHARACTER SET utf8 COLLATE utf8_bin")
	_, err = db.Exec("use " + cfg.Server.Mysql.Database)
	if err != nil {
		return nil, err
	}
	if err := mysql.Migrate(db); err != nil {
		return nil, err
	}
	datasource = cfg.Server.Mysql.Username + ":" + cfg.Server.Mysql.Password + "@tcp(" + cfg.Server.Mysql.Address + ")/" + cfg.Server.Mysql.Database + "?charset=utf8&parseTime=true"
	gdb, err := gorm.Open("mysql", datasource)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	s := &store{gdb}
	_:=s
	dbs := &core.DBStore{
		//UserStore:        s,
	}
	return dbs, nil
}

func (s *store) GetDB() *gorm.DB {
	return s.db
}
