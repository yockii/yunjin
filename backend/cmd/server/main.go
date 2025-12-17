package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"github.com/yockii/yunjin/internal/cache"
	"github.com/yockii/yunjin/internal/config"
	"github.com/yockii/yunjin/internal/controller"
	"github.com/yockii/yunjin/internal/database"
	"github.com/yockii/yunjin/internal/util"
)

func init() {
	viper.SetDefault("server.port", 8080)
}

func main() {
	// Initialize configuration
	config.InitLogger()
	util.InitSnowflake(viper.GetUint64("server.nodeId"))

	database.Initialize()
	defer database.Close()

	cache.InitializeRedis()
	defer cache.CloseRedis()

	database.AutoMigrate()

	app := fiber.New(fiber.Config{
		AppName:      "云烬 v1.0.0",
		ServerHeader: "云烬 Server",
	})
	controller.InitializeRouter(app)
	app.Listen(fmt.Sprintf(":%d", viper.GetInt("server.port")), fiber.ListenConfig{
		DisableStartupMessage: true,
	})
}
