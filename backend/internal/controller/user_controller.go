package controller

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yockii/yunjin/internal/cache"
	"github.com/yockii/yunjin/internal/dao"
	"github.com/yockii/yunjin/internal/database"
	"github.com/yockii/yunjin/internal/domain"
	"github.com/yockii/yunjin/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userController struct{}

func init() {
	controllers = append(controllers, &userController{})
}

func (controller *userController) InitializeRouter(r fiber.Router) {
	r.Post("/register", controller.register)
	r.Post("/login", controller.login)
}

func (controller *userController) register(c fiber.Ctx) error {
	user := new(model.User)
	if err := c.Bind().Body(user); err != nil {
		slog.Error("绑定失败", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Code: fiber.StatusBadRequest,
			Data: nil,
			Msg:  "绑定失败",
		})
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("密码加密失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Code: fiber.StatusInternalServerError,
			Data: nil,
			Msg:  "密码加密失败",
		})
	}
	user.Password = string(pwd)

	if err := gorm.G[model.User](database.DB).Create(c, user); err != nil {
		slog.Error("创建用户失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Code: fiber.StatusInternalServerError,
			Data: nil,
			Msg:  "创建用户失败",
		})
	}

	user.Password = ""

	return c.Status(fiber.StatusCreated).JSON(domain.Response{
		Code: fiber.StatusCreated,
		Data: user,
		Msg:  "注册成功",
	})
}

func (controller *userController) login(c fiber.Ctx) error {
	user := new(model.User)
	if err := c.Bind().Body(user); err != nil {
		slog.Error("绑定失败", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Code: fiber.StatusBadRequest,
			Data: nil,
			Msg:  "绑定失败",
		})
	}

	if u, err := gorm.G[model.User](database.DB).Where(dao.User.Username.Eq(user.Username)).First(c); err != nil {
		slog.Error("查询用户失败", "err", err)
		return c.Status(fiber.StatusNotFound).JSON(domain.Response{
			Code: fiber.StatusNotFound,
			Data: nil,
			Msg:  "用户不存在",
		})
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
		if err != nil {
			slog.Error("密码验证失败", "err", err)
			return c.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				Code: fiber.StatusUnauthorized,
				Data: nil,
				Msg:  "密码错误",
			})
		}

		u.Password = ""

		// 生成token
		return generateLoginResponse(c, &u)
	}
}

func generateLoginResponse(c fiber.Ctx, user *model.User) error {
	// token用简单的随机字符串即可，存入缓存对应用户信息
	token, err := gonanoid.New()
	if err != nil {
		slog.Error("生成token失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Code: fiber.StatusInternalServerError,
			Data: nil,
			Msg:  "生成token失败",
		})
	}
	if err = cache.Set(token, user.ID, time.Hour*24); err != nil {
		slog.Error("缓存token失败", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			Code: fiber.StatusInternalServerError,
			Data: nil,
			Msg:  "缓存token失败",
		})
	}

	user.Password = ""

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Code: fiber.StatusOK,
		Data: &domain.LoginRespData{
			Token: token,
			User:  user,
		},
		Msg: "登录成功",
	})
}
