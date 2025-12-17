package model

// 日更信息
type WritingSession struct {
	BaseModel
	UserID    uint64 `gorm:"not null;index" json:"user_id"`
	ProjectID uint64 `gorm:"not null;index" json:"project_id"`
	Date      string `gorm:"type:date;index" json:"date"` // YYYY-MM-DD
	WordCount int    `gorm:"default:0" json:"word_count"` // 当日新增字数
	Duration  int    `gorm:"default:0" json:"duration"`   // 写作时长(分钟)
}

func init() {
	Models = append(Models, &WritingSession{})
}
