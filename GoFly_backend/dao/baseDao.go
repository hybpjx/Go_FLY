package dao

import (
	"gofly/global"
	"gorm.io/gorm"
)

type BaseDao struct {
	ORM *gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{ORM: global.DB}
}
