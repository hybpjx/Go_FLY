package dao

import (
	"gofly/service/dto"
	"gorm.io/gorm"
)

// Paginate 通用分页函数实现与定义
func Paginate(p dto.Paginate) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((p.GetPage() - 1) * p.Limit).Limit(p.GetLimit())
	}
}
