package model

// 角色状态快照
type CharacterState struct {
	BaseModel
	ProjectID    uint64  `gorm:"not null;index" json:"project_id"`
	CharacterID  uint64  `gorm:"not null;index" json:"character_id"`
	StoryEventID uint64  `gorm:"not null;index" json:"story_event_id"`
	TraitID      uint64  `gorm:"not null;index" json:"trait_id"` // 状态记录的是event id剧情节点结束后的，即场景后生效
	TraitValue   float64 `gorm:"default:50" json:"trait_value"`
	Note         string  `gorm:"size:200" json:"note"` // 备注,为什么变化？
}

func init() {
	Models = append(Models, &CharacterState{})
}
