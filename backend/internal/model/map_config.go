package model

type MapConfig struct {
	BaseModel
	ProjectID uint64  `gorm:"not null;index" json:"project_id"`
	Name      string  `gorm:"size:200" json:"name"` // 如“九州全图”
	ImageURL  string  `gorm:"size:500" json:"image_url"`
	Width     float64 `gorm:"not null" json:"width"` // 图片原始宽度
	Height    float64 `gorm:"not null" json:"height"`

	// 核心计算参数
	ScaleRatio float64 `gorm:"default:1" json:"scale_ratio"`  // 比例尺：1像素 = 多少公里/里
	UnitName   string  `gorm:"default:'公里'" json:"unit_name"` // 单位名称
}

func init() {
	Models = append(Models, &MapConfig{})
}
