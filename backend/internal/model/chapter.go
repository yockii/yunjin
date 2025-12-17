package model

// 章节
type Chapter struct {
	BaseModel
	ProjectID uint64 `gorm:"not null;index" json:"project_id"`
	VolumeID  uint64 `gorm:"not null;index" json:"volume_id"`
	Seq       int    `gorm:"not null" json:"seq"` // 章节序号
	Title     string `gorm:"size:200;not null" json:"title"`
	// 正文内容（建议独立表或大字段）
	Content   string `gorm:"type:longtext" json:"content"`
	WordCount int    `gorm:"default:0" json:"word_count"`
	// 状态
	Status int `gorm:"default:0" json:"status"` // 0:草稿, 1:已发布
}

func init() {
	Models = append(Models, &Chapter{})
}
