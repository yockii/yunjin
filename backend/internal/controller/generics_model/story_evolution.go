package gm

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yockii/yunjin/internal/domain"
	"github.com/yockii/yunjin/internal/model"
	"gorm.io/gorm"
)

type StoryEvolution struct {
	domain.StoryEvolutionCondition
	model *model.StoryEvolution
}

func (t *StoryEvolution) NewModel() *model.StoryEvolution {
	if t.model == nil {
		t.model = &model.StoryEvolution{}
	}
	return t.model
}

func (t *StoryEvolution) CheckCreateFields() (string, bool) {
	if t.model == nil {
		return "model is nil", false
	}
	if t.model.ProjectID == 0 {
		return "project id is empty", false
	}
	if t.model.Step == 0 {
		return "step is empty", false
	}
	return "", true
}

func (t *StoryEvolution) CheckUpdateFields(c fiber.Ctx, m *model.Project) (string, bool) {
	if t.model == nil {
		return "model is nil", false
	}
	return "", true
}
func (t *StoryEvolution) CheckDeleteFields(c fiber.Ctx, m *model.Project) (string, bool) {
	if t.model == nil {
		return "model is nil", false
	}
	return "", true
}

func (t *StoryEvolution) UpdateSelection() []string {
	return nil
}

func (t *StoryEvolution) UpdateOmits() []string {
	return []string{"project_id", "step"}
}

func (t *StoryEvolution) CurrentUser(userID uint64) {

}

func (t *StoryEvolution) PaginateRequest() *domain.PaginateRequest {
	return &t.StoryEvolutionCondition.PaginateRequest
}

func (t *StoryEvolution) BuildCondition(c fiber.Ctx, i gorm.Interface[*model.Project]) (ci gorm.ChainInterface[*model.Project]) {
	cond := t.StoryEvolutionCondition
	ci = i.Where("1 = 1")
	if cond.ProjectID != 0 {
		ci = ci.Where("project_id = ?", cond.ProjectID)
	}
	return ci
}
