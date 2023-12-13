package injector

import (
	"config-management-go/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitGormDB() (*gorm.DB, func(), error) {
	cfg := config.C.Gorm
	db, err := gorm.Open(cfg.Driver, cfg.Dsn)
	cleanFunc := func() {
		db.Close()
	}
	if err != nil {
		return db, cleanFunc, err
	}

	return db, cleanFunc, nil
}
