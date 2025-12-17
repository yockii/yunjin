package model

// 故事演化，描述了小说的演化过程，包括小说的步骤、内容和字数等信息。
type StoryEvolution struct {
	BaseModel
	ProjectID uint64 `gorm:"not null;index" json:"project_id"`
	Step      int    `gorm:"not null" json:"step"` // 1.一句话；2.一段式；3.一页纸；6.四页纸 （角色步骤等在其他模块）
	Content   string `gorm:"type:text" json:"content"`
	WordCount int    `gorm:"not null" json:"word_count"`
}

func init() {
	Models = append(Models, &StoryEvolution{})
}
