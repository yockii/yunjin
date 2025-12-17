package model

// 道具属性快照
type ItemState struct {
	BaseModel
	ProjectID    uint64 `gorm:"not null;index" json:"project_id"`
	ItemID       uint64 `gorm:"not null;index" json:"item_id"`
	StoryEventID uint64 `gorm:"not null;index" json:"story_event_id"`

	// 归属权变更 (核心补充)
	OwnerID  uint64 `gorm:"index" json:"owner_id"`    // 持有者CharacterID，0表示无主或环境持有
	Location string `gorm:"size:200" json:"location"` // 如果无主，可能在某个地点

	// 属性变更 (如果某个属性在这个场景没变，可以不存，或者存全量，视策略而定)
	TraitID    uint64  `gorm:"index" json:"trait_id"`
	TraitValue float64 `gorm:"default:0" json:"trait_value"`

	Note string `gorm:"size:200" json:"note"` // 变更备注：如“战斗中折断”
}

func init() {
	Models = append(Models, &ItemState{})
}
