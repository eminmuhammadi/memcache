package http

import (
	"errors"
	"fmt"
	"time"

	config "github.com/eminmuhammadi/memcache/config"
	json "github.com/goccy/go-json"
	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	INTERNAL_SERVER_ERROR = "Internal Server Error"
)

var APP_CONFIGURATION = fiber.Config{
	EnablePrintRoutes:     config.DefaultConfig.EnablePrintRoutes,
	ReduceMemoryUsage:     config.DefaultConfig.ReduceMemoryUsage,
	DisableStartupMessage: config.DefaultConfig.DisableStartupMessage,
	BodyLimit:             config.DefaultConfig.BodyLimit,
	Concurrency:           config.DefaultConfig.Concurrency,
	ReadBufferSize:        config.DefaultConfig.ReadBufferSize,
	WriteBufferSize:       config.DefaultConfig.WriteBufferSize,
	ReadTimeout:           time.Second * time.Duration(config.DefaultConfig.ReadTimeout),
	WriteTimeout:          time.Second * time.Duration(config.DefaultConfig.WriteTimeout),
	IdleTimeout:           time.Second * time.Duration(config.DefaultConfig.IdleTimeout),
	JSONEncoder:           json.Marshal,
	JSONDecoder:           json.Unmarshal,
	ErrorHandler:          ErrorHandler,
	Prefork:               config.DefaultConfig.Prefork,
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	return ctx.Status(code).SendString(fmt.Sprintf("Error: %s", err.Error()))
}

func Create() *fiber.App {
	return fiber.New(APP_CONFIGURATION)
}

func Start(app *fiber.App, db *gorm.DB, hostname string, port string) error {
	ConfigureMiddleware(app)
	CreateRoutes(db, app)

	return app.Listen(fmt.Sprintf(
		"%s:%s",
		hostname,
		port,
	))
}

func StartSecure(app *fiber.App, db *gorm.DB, hostname string, port string, certFile string, keyFile string) error {
	ConfigureMiddleware(app)
	CreateRoutes(db, app)

	return app.ListenTLS(fmt.Sprintf(
		"%s:%s",
		hostname,
		port,
	), certFile, keyFile)
}
