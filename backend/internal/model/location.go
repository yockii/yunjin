package model

// 地点/地图系统
type Location struct {
	BaseModel
	ProjectID   uint64  `gorm:"not null;index" json:"project_id"`
	ParentID    uint64  `gorm:"index" json:"parent_id"` // 支持层级，如 帝国->皇宫->御书房
	Name        string  `gorm:"size:200;not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Type        string  `gorm:"size:50" json:"type"`           // 地形、建筑、房间
	MapImageURL string  `gorm:"size:500" json:"map_image_url"` // 支持上传地图图片
	MapConfigID uint64  `gorm:"index" json:"map_config_id"`    // 属于哪张地图
	CoordX      float64 `gorm:"default:0" json:"coord_x"`      // X坐标
	CoordY      float64 `gorm:"default:0" json:"coord_y"`      // Y坐标

	// 可选：该地点的地形阻力（平原=1，沼泽=2，山地=3），用于影响移动速度
	TerrainDrag float64 `gorm:"default:1" json:"terrain_drag"`
}

func init() {
	Models = append(Models, &Location{})
}
