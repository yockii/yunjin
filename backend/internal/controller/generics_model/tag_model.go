package gm

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yockii/yunjin/internal/constant"
	"github.com/yockii/yunjin/internal/domain"
	"github.com/yockii/yunjin/internal/model"
	"gorm.io/gorm"
)

type Tag struct {
	domain.TagCondition
	model *model.Tag
}

func (t *Tag) NewModel() *model.Tag {
	if t.model != nil {
		return t.model
	}
	t.model = &model.Tag{}
	return t.model
}

func (t *Tag) CheckCreateFields() (string, bool) {
	if t.model == nil {
		return "model is nil", false
	}
	if t.model.Name == "" {
		return "name is empty", false
	}
	return "", true
}

func (t *Tag) CheckUpdateFields(c fiber.Ctx, m *model.Tag) (string, bool) {
	if fiber.Locals[bool](c, constant.LocalIsAdmin) {
		return "", true
	}
	return "not admin", false
}
func (t *Tag) CheckDeleteFields(c fiber.Ctx, m *model.Tag) (string, bool) {
	if fiber.Locals[bool](c, constant.LocalIsAdmin) {
		return "", true
	}
	return "not admin", false
}

func (t *Tag) UpdateSelection() []string {
	return nil
}

func (t *Tag) UpdateOmits() []string {
	return nil
}

func (t *Tag) CurrentUser(uint64) {
}

func (t *Tag) PaginateRequest() *domain.PaginateRequest {
	return &t.TagCondition.PaginateRequest
}

func (t *Tag) BuildCondition(c fiber.Ctx, i gorm.Interface[*model.Tag]) (ci gorm.ChainInterface[*model.Tag]) {
	cond := t.TagCondition
	ci = i.Where("1 = 1")
	if cond.Name != "" {
		ci = ci.Where("name like ?", "%"+cond.Name+"%")
	}
	return ci
}
