package model

// 项目，每个小说都是一个项目
type Project struct {
	BaseModel
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	OwnerID     uint64 `gorm:"not null;index" json:"owner_id"`
	CurrentStep int    `gorm:"default:1" json:"current_step"`

	Tags []Tag `gorm:"many2many:project_tags;" json:"tags"`
}

func init() {
	Models = append(Models, &Project{})
}
