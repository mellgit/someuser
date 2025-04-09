package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
	log "github.com/sirupsen/logrus"
)

const someUserPath = "/api/v1/someusers"

type SomeUser struct {
	// import service layer
	Cfg    *config.Config
	Logger *log.Entry
}

func NewSomeUser(cfg *config.Config, logger *log.Entry) *SomeUser {
	return &SomeUser{
		Cfg:    cfg,
		Logger: logger,
	}
}

func (h *SomeUser) Register(app *fiber.App) {

	someUser := app.Group(someUserPath)
	someUser.Post("/", h.CreateUser)
	someUser.Get("/", h.GetAllUsers)
	someUser.Get("/:id", h.GetUserByID)
	someUser.Delete("/:id", h.DeleteUserByID)
	someUser.Patch("/:id", h.UpdateUser)
}

func (h *SomeUser) CreateUser(ctx *fiber.Ctx) error {
	payload := model.CreateUserRequest{}
	if err := ctx.BodyParser(&payload); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "ctx.BodyParser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusNotFound).JSON(msgErr)

	}
	h.Logger.Infof("%v", payload)

	// service here

	msg := model.Message{Message: "Create User"}
	return ctx.Status(fiber.StatusOK).JSON(msg)

	h.Logger.Infof("%v", payload)
	return nil
}

func (h *SomeUser) GetAllUsers(ctx *fiber.Ctx) error {
	return nil
}

func (h *SomeUser) GetUserByID(ctx *fiber.Ctx) error {
	return nil
}

func (h *SomeUser) DeleteUserByID(ctx *fiber.Ctx) error {
	return nil
}

func (h *SomeUser) UpdateUser(ctx *fiber.Ctx) error {
	return nil
}
