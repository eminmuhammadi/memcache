package http

import (
	controller "github.com/eminmuhammadi/memcache/controller"
	fiber "github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// CreateRoutes is a function that creates the routes.
func CreateRoutes(db *gorm.DB, app *fiber.App) {
	// =======================================
	//	GET /:id
	//  Get Cache Value
	// =======================================
	app.Get("/:id<guid>", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		// check id is uuid or not
		_, err := uuid.FromString(id)
		if err != nil {
			return ctx.SendStatus(fiber.StatusNoContent)
		}

		value, err := controller.GetValue(id, db, ctx)

		if err != nil {
			return err
		}

		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		if value == "" {
			return ctx.SendStatus(fiber.StatusNotFound)
		}

		return ctx.Status(fiber.StatusOK).SendString(value)
	})

	// =======================================
	//	POST /
	//  Create Cache Value
	// =======================================
	app.Post("/", func(ctx *fiber.Ctx) error {
		value, err := controller.Create(db, ctx)

		if err != nil {
			return err
		}

		if value == "" {
			return ctx.SendStatus(fiber.StatusNotFound)
		}

		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		return ctx.Status(fiber.StatusCreated).SendString(value)
	})

	// =======================================
	//	DELETE /:id
	//  Delete Cache Value
	// =======================================
	app.Delete("/:id<guid>", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		// check id is uuid or not
		_, err := uuid.FromString(id)
		if err != nil {
			return ctx.SendStatus(fiber.StatusNoContent)
		}

		if err := controller.Delete(id, db, ctx); err != nil {
			return err
		}

		return ctx.SendStatus(fiber.StatusAccepted)
	})

	// =======================================
	//	PUT /:id
	//  Update Cache Value
	// =======================================
	app.Put("/:id<guid>", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		// check id is uuid or not
		_, err := uuid.FromString(id)
		if err != nil {
			return ctx.SendStatus(fiber.StatusNoContent)
		}

		value, err := controller.Update(id, db, ctx)

		if err != nil {
			return err
		}

		if value == "" {
			return ctx.SendStatus(fiber.StatusNotFound)
		}

		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		return ctx.Status(fiber.StatusNoContent).SendString(value)
	})
}
