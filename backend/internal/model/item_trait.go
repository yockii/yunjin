package model

// 道具属性
type ItemTrait struct {
	BaseModel
	ProjectID    uint64  `gorm:"not null;index" json:"project_id"`
	Name         string  `gorm:"size:200" json:"name"` // 名称,如信任度、愤怒值、希望指数
	Description  string  `gorm:"size:2000" json:"description"`
	MinValue     float64 `gorm:"default:0" json:"min_value"`
	MaxValue     float64 `gorm:"default:100" json:"max_value"`
	DefaultValue float64 `gorm:"default:100" json:"default_value"`
}

func init() {
	Models = append(Models, &ItemTrait{})
}
