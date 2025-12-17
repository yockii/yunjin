package controller

import (
	"github.com/gofiber/fiber/v3"
	gm "github.com/yockii/yunjin/internal/controller/generics_model"
	"github.com/yockii/yunjin/internal/model"
)

type tagController struct {
	GenericsController[*model.Tag, *gm.Tag]
}

func init() {
	tc := &tagController{}
	tc.newDomainInstance = func() *gm.Tag {
		return new(gm.Tag)
	}
	controllers = append(controllers, tc)
}

func (controller *tagController) InitializeRouter(router fiber.Router) {
	r := router.Group("/tag")
	r.Post("/create", controller.Create)
	r.Get("/list", controller.List)
}
