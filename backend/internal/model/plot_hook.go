package model

// 伏笔/挖坑
type PlotHook struct {
	BaseModel
	ProjectID   uint64 `gorm:"not null;index" json:"project_id"`
	Name        string `gorm:"size:200;not null" json:"name"` // 伏笔代号，如“玉佩的缺口”
	Description string `gorm:"type:text" json:"description"`  // 详细描述
	Importance  int    `gorm:"default:1" json:"importance"`   // 重要程度 1-5
	// 挖坑位置
	SetupSceneID uint64 `gorm:"index" json:"setup_scene_id"` // 在哪个场景埋下的
	// 填坑位置 (Nullable)
	ResolveSceneID *uint64 `gorm:"index" json:"resolve_scene_id"` // 在哪个场景解决的
	// 状态管理
	// 0: 未填 (Open)
	// 1: 已填 (Resolved)
	// 2: 废弃 (Abandoned) - 作者决定不用了
	// 3: 跨书伏笔 (Cross-Book) - 故意留给下一部
	Status int `gorm:"default:0" json:"status"`

	// 宇宙观/跨书扩展
	TargetSeries string `gorm:"size:200" json:"target_series"` // 预设填坑的系列名，如 "三体2"
	Notes        string `gorm:"type:text" json:"notes"`        // 作者备注

}

func init() {
	Models = append(Models, &PlotHook{})
}
