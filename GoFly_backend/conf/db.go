package conf

import (
	"github.com/spf13/viper"
	"gofly/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() (*gorm.DB, error) {
	logModelLevel := logger.Info

	// 如果是开发者模式 就设置为info
	if viper.GetBool("model.development") == false {
		logModelLevel = logger.Error
	}

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		// 命名规则
		NamingStrategy: schema.NamingStrategy{
			// 表名前缀
			TablePrefix: "sys_",
			// 是否不加附属
			SingularTable: true,
		},
		Logger:                                   logger.Default.LogMode(logModelLevel),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return nil, err
	}

	defer func() {
		// 配置数据的共行个数 和最大连接数
		sqlDB, err1 := db.DB() //连接池
		if err1 != nil {
			panic(err1)
		}
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(viper.GetInt("db.MaxIdleConn"))

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(viper.GetInt("db.MaxOpenConn"))

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("db.ConnMaxLifetime")))
	}()

	db.AutoMigrate(&model.User{})

	// 创建表结构
	return db, nil
}
