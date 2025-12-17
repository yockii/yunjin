package model

// 场景参与者，描述了场景中角色的参与情况和角色在场景中的角色
type SceneParticipants struct {
	BaseModel
	SceneID     uint64 `gorm:"not null;index" json:"scene_id"`
	CharacterID uint64 `gorm:"not null;index" json:"character_id"`
	RoleInScene string `gorm:"size:200" json:"role_in_scene"` // 角色在场景中的角色
}

func init() {
	Models = append(Models, &SceneParticipants{})
}
