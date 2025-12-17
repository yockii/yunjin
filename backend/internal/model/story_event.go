package model

// 剧情节点，是小说中的一个片段，通常由一个或多个角色参与，描述了角色之间的互动和情节的发展。
type StoryEvent struct {
	BaseModel
	ProjectID      uint64 `gorm:"not null;index" json:"project_id"`
	Seq            int    `gorm:"not null" json:"seq"`              // 顺序
	Title          string `gorm:"size:200" json:"title"`            // 标题
	Summary        string `gorm:"type:text" json:"summary"`         // 场景摘要
	PovCharacterID uint64 `gorm:"not null" json:"pov_character_id"` // 视角角色ID
	Goal           string `gorm:"type:text" json:"goal"`            // POV角色在此场景的目标
	Conflict       string `gorm:"type:text" json:"conflict"`        // 角色冲突/阻碍
	Outcome        string `gorm:"type:text" json:"outcome"`         // 场景结局

	IsUsed bool `gorm:"not null" json:"is_used"` // 是否使用过(写进章节)
}

func init() {
	Models = append(Models, &StoryEvent{})
}
