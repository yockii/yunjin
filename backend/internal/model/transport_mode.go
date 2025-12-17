package model

type TransportMode struct {
	BaseModel
	ProjectID uint64  `gorm:"not null;index" json:"project_id"`
	Name      string  `gorm:"size:100" json:"name"`  // 如：步行、御剑飞行、马车
	Speed     float64 `gorm:"not null" json:"speed"` // 速度：每小时多少公里/里
}

func init() {
	Models = append(Models, &TransportMode{})
}
