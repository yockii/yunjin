package model

// 剧情节点参与者，描述了剧情节点中角色的参与情况和角色在剧情节点中的角色
type StoryEventParticipants struct {
	BaseModel
	StoryEventID uint64 `gorm:"not null;index" json:"story_event_id"`
	CharacterID  uint64 `gorm:"not null;index" json:"character_id"`
	RoleInEvent  string `gorm:"size:200" json:"role_in_event"` // 角色在节点中的角色
}

func init() {
	Models = append(Models, &StoryEventParticipants{})
}
