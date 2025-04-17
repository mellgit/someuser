package transport

import "github.com/gofiber/fiber/v2"

type Transport interface {
	Register(app *fiber.App)
	CreateUser(ctx *fiber.Ctx) error
	GetAllUsers(ctx *fiber.Ctx) error
	GetUserByID(ctx *fiber.Ctx) error
	DeleteUserByID(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
}
