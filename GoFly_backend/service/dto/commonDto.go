package dto

// CommonIDDTO ==== 通用ID对应的DTO ====
type CommonIDDTO struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// Paginate 分页对应的DTO
type Paginate struct {
	// 如果是返回字段 那就直接不返回
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

// GetPage 获取页码数
func (p *Paginate) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}

	return p.Page
}

// GetLimit 获取每页展示数量
func (p *Paginate) GetLimit() int {
	if p.Limit <= 0 {
		p.Page = 10
	}

	return p.Limit
}
