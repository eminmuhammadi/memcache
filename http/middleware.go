package http

import (
	"github.com/ansrivas/fiberprometheus/v2"
	config "github.com/eminmuhammadi/memcache/config"
	_dbPackage "github.com/eminmuhammadi/memcache/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	uuid "github.com/satori/go.uuid"
)

func ConfigureMiddleware(app *fiber.App) {
	EnableRecovers(app)
	EnableCompression(app)
	EnableETAG(app)
	EnableCORS(app)
	EnableRequestID(app)
	EnableLogger(app)

	// Monitoring
	EnableMetrics(app)
	EnableHealthCheck(app)
	EnableMonitoring(app)
}

func EnableRecovers(app *fiber.App) {
	app.Use(recover.New())
}

func EnableCompression(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
}

func EnableETAG(app *fiber.App) {
	app.Use(etag.New(etag.Config{
		Weak: true,
	}))
}

func EnableCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		MaxAge:       86400,
	}))
}

func EnableRequestID(app *fiber.App) {
	app.Use(requestid.New(requestid.Config{
		Header: "X-REQUEST-ID",
		Generator: func() string {
			return uuid.NewV4().String()
		},
	}))
}

func EnableLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "${time} [${ip}] ${status} ${protocol} ${method} ${path} ${latency}\n",
		TimeZone:   config.DefaultConfig.Timezone,
		TimeFormat: _dbPackage.TimeNowString(),
	}))
}

func EnableMonitoring(app *fiber.App) {
	app.Get("/_/monitoring", monitor.New(
		monitor.Config{
			APIOnly: true,
		},
	))
}

func EnableHealthCheck(app *fiber.App) {
	app.Get("/_/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
}

func EnableMetrics(app *fiber.App) {
	prometheus := fiberprometheus.New("github.com/eminmuhammadi/memcache")
	prometheus.RegisterAt(app, "/_/metrics")
	app.Use(prometheus.Middleware)
}
