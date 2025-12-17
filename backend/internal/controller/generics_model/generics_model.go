package gm

import (
	"github.com/gofiber/fiber/v3"
	"github.com/yockii/yunjin/internal/domain"
	"github.com/yockii/yunjin/internal/model"
	"gorm.io/gorm"
)

type Model[M model.GenericModel] interface {
	NewModel() M
	CheckCreateFields() (string, bool)
	CheckUpdateFields(fiber.Ctx, M) (string, bool)
	CheckDeleteFields(fiber.Ctx, M) (string, bool)
	UpdateSelection() []string
	UpdateOmits() []string
	CurrentUser(uint64)
	PaginateRequest() *domain.PaginateRequest
	BuildCondition(fiber.Ctx, gorm.Interface[M]) gorm.ChainInterface[M]
}
