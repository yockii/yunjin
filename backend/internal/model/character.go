package model

// 小说人物角色
type Character struct {
	BaseModel
	ProjectID   uint64 `gorm:"not null;index" json:"project_id"`
	Name        string `gorm:"not null" json:"name"`
	RoleType    string `gorm:"not null" json:"role_type"`
	OneSentence string `gorm:"type:text" json:"one_sentence"` // 一句话简介
	Motivation  string `gorm:"type:text" json:"motivation"`   // 动机
	Goal        string `gorm:"type:text" json:"goal"`         // 目标
	Conflict    string `gorm:"type:text" json:"conflict"`     // 冲突
	Epiphany    string `gorm:"type:text" json:"epiphany"`     // 顿悟
	DetailedBio string `gorm:"type:text" json:"detailed_bio"` // 人物宝典（背景、性格、外貌等）
	Notes       string `gorm:"type:text" json:"notes"`        // 备注
}

func init() {
	Models = append(Models, &Character{})
}
