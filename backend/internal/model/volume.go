package model

type Volume struct {
	BaseModel
	ProjectID uint64 `gorm:"not null;index" json:"project_id"`
	Title     string `gorm:"size:200;not null" json:"title"`
	Order     int    `gorm:"not null" json:"order"` // 卷序
	Summary   string `gorm:"type:text" json:"summary"`
}

func init() {
	Models = append(Models, &Volume{})
}
