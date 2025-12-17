package model

// 关系状态快照，主要对关系的值变化进行记录
type RelationshipState struct {
	BaseModel
	ProjectID      uint64  `gorm:"not null;index" json:"project_id"`
	RelationshipID uint64  `gorm:"not null;index" json:"relationship_id"`
	StoryEventID   uint64  `gorm:"not null;index" json:"story_event_id"`
	TraitID        uint64  `gorm:"not null;index" json:"trait_id"`
	TraitValue     float64 `gorm:"not null;index" json:"trait_value"`
	Note           string  `gorm:"not null;index" json:"note"`
}

func init() {
	Models = append(Models, &RelationshipState{})
}
