package model

// 道具
type Item struct {
	BaseModel
	ProjectID   uint64 `gorm:"not null;index" json:"project_id"`
	Name        string `gorm:"size:200;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	ItemType    string `gorm:"size:50" json:"item_type"`
	IsUnique    bool   `gorm:"default:true" json:"is_unique"`
}

func init() {
	Models = append(Models, &Item{})
}
