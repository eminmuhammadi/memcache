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

// This is a message that is sent when the server is down.
const INTERNAL_SERVER_ERROR string = "Internal Server Error"

// App configuration.
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

// ErrorHandler is the error handler for the app.
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	return ctx.Status(code).SendString(fmt.Sprintf("Error: %s", err.Error()))
}

// Create is a function that creates the app.
func Create() *fiber.App {
	return fiber.New(APP_CONFIGURATION)
}

// Start is a function that starts the app.
func Start(app *fiber.App, db *gorm.DB, hostname string, port string) error {
	// Configure the middleware.
	ConfigureMiddleware(app)

	// Create the routes.
	CreateRoutes(db, app)

	return app.Listen(fmt.Sprintf(
		"%s:%s",
		hostname,
		port,
	))
}

// StartSecure is a function that starts the app in secure mode.
func StartSecure(app *fiber.App, db *gorm.DB, hostname string, port string, certFile string, keyFile string) error {
	// Configure the middleware.
	ConfigureMiddleware(app)

	// Create the routes.
	CreateRoutes(db, app)

	return app.ListenTLS(fmt.Sprintf(
		"%s:%s",
		hostname,
		port,
	), certFile, keyFile)
}
