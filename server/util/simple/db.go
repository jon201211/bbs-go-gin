package simple

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormModel struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

const DRIVER_MYSQL = "mysql"
const DRIVER_SQLITE = "sqlite"

func OpenDB(driver string, dsn string, config *gorm.Config, maxIdleConns, maxOpenConns int, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		}
	}

	switch driver {
	case DRIVER_SQLITE:
		if db, err = gorm.Open(sqlite.Open(dsn), config); err != nil {
			log.Errorf("opens sqlite database failed: %s", err.Error())
			return
		}
		break
	case DRIVER_MYSQL:
		if db, err = gorm.Open(mysql.Open(dsn), config); err != nil {
			log.Errorf("opens mysql database failed: %s", err.Error())
			return
		}
		break
	default:
		log.Errorf("opens database failed: %s", "unknown type")
		return
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
	} else {
		log.Error(err)
	}

	if err = db.AutoMigrate(models...); nil != err {
		log.Errorf("auto migrate tables failed: %s", err.Error())
	}
	return
}

// 获取数据库链接
func DB() *gorm.DB {
	return db
}

// 关闭连接
func CloseDB() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		log.Errorf("Disconnect from database failed: %s", err.Error())
	}
}
