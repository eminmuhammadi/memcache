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

// The http path for monitoring.
const MONITORING_PATH = "/_/monitoring"

// The http path for healthcheck.
const HEALTHCHECK_PATH = "/_/healthcheck"

// The http path for metrics.
const METRICS_PATH = "/_/metrics"

// Prometheus service name.
const PROMETHEUS_SERVICE_NAME = "github.com/eminmuhammadi/memcache"

// Default configuration for the middleware.
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

// It enables the recovery middleware.
func EnableRecovers(app *fiber.App) {
	app.Use(recover.New())
}

// It enables the compression middleware.
func EnableCompression(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
}

// It enables the etag middleware.
func EnableETAG(app *fiber.App) {
	app.Use(etag.New(etag.Config{
		Weak: true,
	}))
}

// It enables the cors middleware.
func EnableCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		MaxAge:       86400,
	}))
}

// It enables the request id middleware.
func EnableRequestID(app *fiber.App) {
	app.Use(requestid.New(requestid.Config{
		Header: "X-REQUEST-ID",
		Generator: func() string {
			return uuid.NewV4().String()
		},
	}))
}

// It enables the logger middleware.
func EnableLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "${time} [${ip}] ${status} ${protocol} ${method} ${path} ${latency}\n",
		TimeZone:   config.DefaultConfig.Timezone,
		TimeFormat: _dbPackage.TimeNowString(),
	}))
}

// It enables the monitoring middleware.
func EnableMonitoring(app *fiber.App) {
	app.Get(MONITORING_PATH, monitor.New(
		monitor.Config{
			APIOnly: true,
		},
	))
}

// It enables the health check middleware.
func EnableHealthCheck(app *fiber.App) {
	app.Get(HEALTHCHECK_PATH, func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
}

// It enables the metrics middleware.
func EnableMetrics(app *fiber.App) {
	prometheus := fiberprometheus.New(PROMETHEUS_SERVICE_NAME)

	prometheus.RegisterAt(app, METRICS_PATH)

	app.Use(prometheus.Middleware)
}
