package controller

import (
	"fmt"

	_dbPackage "github.com/eminmuhammadi/memcache/db"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CacheReq struct {
	Value string `json:"value" xml:"value" form:"value" validate:"required,min=1"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func Validator(cache *CacheReq) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(cache)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func Create(db *gorm.DB, ctx *fiber.Ctx) (string, error) {
	req := new(CacheReq)

	if err := ctx.BodyParser(req); err != nil {
		return "", err
	}

	_errors := Validator(req)

	if len(_errors) > 0 {
		return "", fmt.Errorf("validation failed: %v", _errors)
	}

	cache := _dbPackage.Cache{
		ID:        uuid.NewV4().String(),
		Value:     req.Value,
		CreatedAt: _dbPackage.TimeNow(),
		UpdatedAt: _dbPackage.TimeNow(),
	}

	if err := db.Create(&cache).Error; err != nil {
		return "", err
	}

	return cache.ID, nil
}

func Update(id string, db *gorm.DB, ctx *fiber.Ctx) (string, error) {
	req := new(CacheReq)

	if err := ctx.BodyParser(req); err != nil {
		return "", err
	}

	if id == "" {
		return "", fmt.Errorf("id is required")
	}

	_errors := Validator(req)

	if len(_errors) > 0 {
		return "", fmt.Errorf("validation failed: %v", _errors)
	}

	cache := _dbPackage.Cache{
		ID:        id,
		Value:     req.Value,
		UpdatedAt: _dbPackage.TimeNow(),
	}

	if err := db.Save(&cache).Error; err != nil {
		return "", err
	}

	return cache.ID, nil
}

func Delete(id string, db *gorm.DB, ctx *fiber.Ctx) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}

	cache := _dbPackage.Cache{
		ID: id,
	}

	if err := db.Delete(&cache).Error; err != nil {
		return err
	}

	return nil
}

func GetValue(id string, db *gorm.DB, ctx *fiber.Ctx) (string, error) {
	if id == "" {
		return "", fmt.Errorf("id is required")
	}

	cache := _dbPackage.Cache{
		ID: id,
	}

	if err := db.First(&cache).Error; err != nil {
		return "", err
	}

	return cache.Value, nil
}
