package model

type User struct {
	BaseModel
	Username string `gorm:"size:50;not null;unique" json:"username"`
	Password string `gorm:"size:100;not null" json:"password"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
}

func init() {
	Models = append(Models, &User{})
}
