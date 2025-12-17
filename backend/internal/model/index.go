package model

import (
	"github.com/yockii/yunjin/internal/util"
	"gorm.io/cli/gorm/genconfig"
	"gorm.io/gorm"
)

type GenericModel interface {
	GetID() uint64
}

var Models []any

type BaseModel struct {
	ID        uint64 `gorm:"primary_key;auto_increment:false" json:"id" query:"id" params:"id" form:"id"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = util.NextSnowflakeID()
	return
}

func (m *BaseModel) GetID() uint64 {
	return m.ID
}

// //// gorm cli 配置
var _ = genconfig.Config{
	OutPath: "../dao",
	// IncludeInterfaces: []any{},
	IncludeStructs: Models,
}
