package model

// 角色关系定义
type CharacterRelationship struct {
	BaseModel
	ProjectID   uint64 `gorm:"not null;index" json:"project_id"`
	CharacterID uint64 `gorm:"not null;index" json:"character_id"`
	RelationID  uint64 `gorm:"not null;index" json:"relation_id"`
}

func init() {
	Models = append(Models, &CharacterRelationship{})
}
