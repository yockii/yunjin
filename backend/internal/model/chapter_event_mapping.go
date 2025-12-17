package model

type ChapterEventMapping struct {
	BaseModel
	ChapterID    uint64 `gorm:"not null;index" json:"chapter_id"`
	StoryEventID uint64 `gorm:"not null;index" json:"story_event_id"`
	// 可选：记录该事件在章节正文中的大概位置（如果需要精细化）
	OrderInChapter int `gorm:"default:0" json:"order_in_chapter"`
}

func init() {
	Models = append(Models, &ChapterEventMapping{})
}
