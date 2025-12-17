package model

// 场景，场景是小说中的一个片段，通常由一个或多个角色参与，描述了角色之间的互动和情节的发展。
type Scene struct {
	BaseModel
	ProjectID      uint64 `gorm:"not null;index" json:"project_id"`
	Seq            int    `gorm:"not null" json:"seq"`              // 顺序
	Title          string `gorm:"size:200" json:"title"`            // 标题
	PovCharacterID uint64 `gorm:"not null" json:"pov_character_id"` // 视角角色ID
	Summary        string `gorm:"type:text" json:"summary"`         // 场景摘要
	Goal           string `gorm:"type:text" json:"goal"`            // POV角色在此场景的目标
	Conflict       string `gorm:"type:text" json:"conflict"`        // 角色冲突/阻碍
	Outcome        string `gorm:"type:text" json:"outcome"`         // 场景结局
	Chapter        int    `gorm:"not null" json:"chapter"`          // 章节
	Status         int    `gorm:"not null" json:"status"`           // 状态 1-计划 2-草稿 3-完成
	EstimatedWords int    `gorm:"not null" json:"estimated_words"`  // 预计字数
}

func init() {
	Models = append(Models, &Scene{})
}
