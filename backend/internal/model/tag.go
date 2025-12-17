package model

type Tag struct {
	BaseModel
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

func init() {
	Models = append(Models, &Tag{})
}
