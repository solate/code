package orm

import (
	"github.com/jinzhu/gorm"
	"github.com/solate/util/orm"
)

var DB *gorm.DB

func Init(host, prot, database, username, password string,
	maxIdleConns, maxOpenConns, connMaxLifetime int, debug bool) (err error) {
	orm := orm.NewOrm(host, prot, database, username, password)
	orm.SetOption(maxIdleConns, maxOpenConns, connMaxLifetime, debug)
	DB, err = orm.Init()
	return
}
