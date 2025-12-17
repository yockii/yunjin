package controller

import (
	"github.com/gofiber/fiber/v3"
	gm "github.com/yockii/yunjin/internal/controller/generics_model"
	"github.com/yockii/yunjin/internal/model"
)

type projectController struct {
	GenericsController[*model.Project, *gm.Project]
}

func init() {
	controllers = append(controllers, &projectController{
		GenericsController: GenericsController[*model.Project, *gm.Project]{
			newDomainInstance: func() *gm.Project {
				return &gm.Project{}
			},
		},
	})
}

func (controller *projectController) InitializeRouter(router fiber.Router) {
	r := router.Group("/project")
	r.Get("/", controller.List)
	r.Post("/", controller.Create)
	r.Put("/:id", controller.Update)
	r.Delete("/:id", controller.Delete)
}
