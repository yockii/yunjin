package gm

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yockii/yunjin/internal/constant"
	"github.com/yockii/yunjin/internal/domain"
	"github.com/yockii/yunjin/internal/model"
	"gorm.io/gorm"
)

type Project struct {
	domain.ProjectCondition
	model *model.Project
}

func (t *Project) NewModel() *model.Project {
	if t.model == nil {
		t.model = &model.Project{}
	}
	return t.model
}

func (t *Project) CheckCreateFields() (string, bool) {
	if t.model == nil {
		return "model is nil", false
	}
	if t.model.Title == "" {
		return "title is empty", false
	}
	if len(t.model.Tags) == 0 {
		return "tags is empty", false
	}
	return "", true
}

func (t *Project) CheckUpdateFields(c fiber.Ctx, m *model.Project) (string, bool) {
	if t.model == nil {
		return "model is nil", false
	}
	return "", true
}
func (t *Project) CheckDeleteFields(c fiber.Ctx, m *model.Project) (string, bool) {
	if t.model == nil {
		return "model is nil", false
	}
	return "", true
}

func (t *Project) UpdateSelection() []string {
	return nil
}

func (t *Project) UpdateOmits() []string {
	return []string{"owner_id"}
}

func (t *Project) CurrentUser(userID uint64) {
	t.model.OwnerID = userID
}

func (t *Project) PaginateRequest() *domain.PaginateRequest {
	return &t.ProjectCondition.PaginateRequest
}

func (t *Project) BuildCondition(c fiber.Ctx, i gorm.Interface[*model.Project]) (ci gorm.ChainInterface[*model.Project]) {
	cond := t.ProjectCondition
	ci = i.Where("1 = 1")
	if cond.Title != "" {
		ci = ci.Where("title like ?", "%"+cond.Title+"%")
	}
	if cond.OwnerID != 0 {
		ci = ci.Where("owner_id = ?", cond.OwnerID)
	} else {
		if !fiber.Locals[bool](c, constant.LocalIsAdmin) {
			ci = ci.Where("owner_id = ?", fiber.Locals[uint64](c, constant.LocalUserID))
		}
	}
	return ci
}
