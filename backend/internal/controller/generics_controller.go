package controller

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/yockii/yunjin/internal/constant"
	gm "github.com/yockii/yunjin/internal/controller/generics_model"
	"github.com/yockii/yunjin/internal/dao"
	"github.com/yockii/yunjin/internal/database"
	"github.com/yockii/yunjin/internal/domain"
	"github.com/yockii/yunjin/internal/model"
	"gorm.io/gorm"
)

type GenericsController[M model.GenericModel, D gm.Model[M]] struct {
	newDomainInstance func() D
}

func (controller *GenericsController[M, D]) List(c fiber.Ctx) error {
	cond := controller.newDomainInstance()
	if err := c.Bind().Query(cond); err != nil {
		slog.Error("绑定失败", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Code: fiber.StatusBadRequest,
			Data: nil,
			Msg:  "绑定失败",
		})
	}
	pr := cond.PaginateRequest()
	if pr == nil {
		pr = &domain.PaginateRequest{Page: 1, Limit: 10}
	}
	if pr.Page <= 0 {
		pr.Page = 1
	}
	if pr.Limit <= 0 {
		pr.Limit = 10
	}

	var ci gorm.ChainInterface[M]
	d := gorm.G[M](database.DB)
	ci = cond.BuildCondition(c, d)

	var err error
	var count int64
	var list []M
	if count, err = ci.Count(c, dao.BaseModel.ID.Column().Name); err != nil {
		slog.Error("计数失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.NewPaginateResponse(
			fiber.StatusInternalServerError,
			nil,
			"计数失败",
			0,
			pr.Page,
			pr.Limit,
		))
	}
	if count == 0 {
		return c.Status(fiber.StatusOK).JSON(domain.NewPaginateResponse(
			fiber.StatusOK,
			nil,
			"查询成功",
			count,
			pr.Page,
			pr.Limit,
		))
	}

	if list, err = ci.Offset((pr.Page - 1) * pr.Limit).Limit(pr.Limit).Find(c); err != nil {
		slog.Error("查询失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.NewPaginateResponse(
			fiber.StatusInternalServerError,
			nil,
			"查询失败",
			count,
			pr.Page,
			pr.Limit,
		))
	}

	return c.Status(fiber.StatusOK).JSON(domain.NewPaginateResponse(
		fiber.StatusOK,
		list,
		"查询成功",
		count,
		pr.Page,
		pr.Limit,
	))
}

func (controller *GenericsController[M, D]) Create(c fiber.Ctx) error {
	instance := controller.newDomainInstance()
	m := instance.NewModel()
	if err := c.Bind().Body(m); err != nil {
		slog.Error("绑定失败", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Code: fiber.StatusBadRequest,
			Data: nil,
			Msg:  "绑定失败",
		})
	}

	if msg, pass := instance.CheckCreateFields(); !pass {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Code: fiber.StatusBadRequest,
			Data: nil,
			Msg:  msg,
		})
	}

	instance.CurrentUser(fiber.Locals[uint64](c, constant.LocalUserID))

	if err := gorm.G[M](database.DB).Create(c, &m); err != nil {
		slog.Error("创建失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Code: fiber.StatusInternalServerError,
			Data: nil,
			Msg:  "创建失败",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(domain.Response{
		Code: fiber.StatusCreated,
		Data: m,
		Msg:  "创建成功",
	})
}

func (controller *GenericsController[M, D]) Update(c fiber.Ctx) error {
	instance := controller.newDomainInstance()
	m := instance.NewModel()
	if err := c.Bind().Body(m); err != nil {
		slog.Error("绑定失败", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Code: fiber.StatusBadRequest,
			Data: nil,
			Msg:  "绑定失败",
		})
	}

	if m.GetID() == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Code: fiber.StatusBadRequest,
			Data: nil,
			Msg:  "ID不能为空",
		})
	}
	if trueInstance, err := gorm.G[M](database.DB).Where(dao.BaseModel.ID.Eq(m.GetID())).First(c); err != nil {
		slog.Error("查询失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Code: fiber.StatusInternalServerError,
			Data: nil,
			Msg:  "查询失败",
		})
	} else if msg, pass := instance.CheckUpdateFields(c, trueInstance); !pass {
		return c.Status(fiber.StatusForbidden).JSON(domain.Response{
			Code: fiber.StatusForbidden,
			Data: nil,
			Msg:  msg,
		})
	}

	var ci gorm.CreateInterface[M]
	ci = gorm.G[M](database.DB)
	if s := instance.UpdateSelection(); len(s) > 0 {
		if len(s) == 1 {
			ci = ci.Select(s[0])
		} else {
			params := make([]any, len(s)-1)
			for i := 1; i < len(s); i++ {
				params[i-1] = s[i]
			}
			ci = ci.Select(s[0], params...)
		}
	}
	if o := instance.UpdateOmits(); len(o) > 0 {
		ci = ci.Omit(o...)
	}

	if rowsAffected, err := ci.Updates(c, m); err != nil {
		slog.Error("更新失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Code: fiber.StatusInternalServerError,
			Data: nil,
			Msg:  "更新失败",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(domain.Response{
			Code: fiber.StatusOK,
			Data: rowsAffected > 0,
			Msg:  "更新成功",
		})
	}
}

func (controller *GenericsController[M, D]) Delete(c fiber.Ctx) error {
	id := fiber.Params[uint64](c, "id")
	if id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Code: fiber.StatusBadRequest,
			Data: nil,
			Msg:  "ID不能为空",
		})
	}
	if trueInstance, err := gorm.G[M](database.DB).Where(dao.BaseModel.ID.Eq(id)).First(c); err != nil {
		slog.Error("查询失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Code: fiber.StatusInternalServerError,
			Data: nil,
			Msg:  "查询失败",
		})
	} else if msg, pass := (controller.newDomainInstance()).CheckDeleteFields(c, trueInstance); !pass {
		return c.Status(fiber.StatusForbidden).JSON(domain.Response{
			Code: fiber.StatusForbidden,
			Data: nil,
			Msg:  msg,
		})
	}
	if rowsAffected, err := gorm.G[M](database.DB).Where(dao.BaseModel.ID.Eq(id)).Delete(c); err != nil {
		slog.Error("删除失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Code: fiber.StatusInternalServerError,
			Data: nil,
			Msg:  "删除失败",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(domain.Response{
			Code: fiber.StatusOK,
			Data: rowsAffected > 0,
			Msg:  "删除成功",
		})
	}
}
