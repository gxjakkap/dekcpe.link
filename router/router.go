package router

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/gxjakkap/dekcpe.link/utils"
)

func New() *fiber.App {
	rp, _ := os.Getwd()
	vp := filepath.Join(rp, "views")
	e := html.New(vp, ".html")

	tpc := (os.Getenv("PROXY_MODE") != "")

	f := fiber.New(fiber.Config{
		Views:                   e,
		EnableTrustedProxyCheck: tpc,
		ProxyHeader:             utils.GetForwardingHeader(),
	})

	f.Use(logger.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	return f
}
