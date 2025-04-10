package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/service"
	log "github.com/sirupsen/logrus"
)

const someUserPath = "/api/v1/someusers"

type SomeUser struct {
	SomeUserService *service.SomeUser
	Cfg             *config.Config
	Logger          *log.Entry
}

func NewSomeUser(cfg *config.Config, someUserService *service.SomeUser, logger *log.Entry) *SomeUser {
	return &SomeUser{
		Cfg:             cfg,
		SomeUserService: someUserService,
		Logger:          logger,
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)

	}

	if err := h.SomeUserService.CreateUser(payload); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.CreateUser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}

	msg := model.Message{Message: "createUser"}
	return ctx.Status(fiber.StatusOK).JSON(msg)
}

func (h *SomeUser) GetAllUsers(ctx *fiber.Ctx) error {

	if err := h.SomeUserService.GetAllUsers(); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.GetAllUsers",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}
	msg := model.Message{Message: "getAllUsers"}
	return ctx.Status(fiber.StatusOK).JSON(msg)
}

func (h *SomeUser) GetUserByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if err := h.SomeUserService.GetUserByID(uuid.MustParse(id)); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.GetUserByID",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}
	msg := model.Message{Message: "getUserByID"}
	return ctx.Status(fiber.StatusOK).JSON(msg)
}

func (h *SomeUser) DeleteUserByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if err := h.SomeUserService.DeleteUserByID(uuid.MustParse(id)); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.DeleteUserByID",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}
	msg := model.Message{Message: "deleteUserByID"}
	return ctx.Status(fiber.StatusOK).JSON(msg)
}

func (h *SomeUser) UpdateUser(ctx *fiber.Ctx) error {

	payload := model.UpdateUserRequest{}
	if err := ctx.BodyParser(&payload); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "ctx.BodyParser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}
	if err := h.SomeUserService.UpdateUser(payload); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.UpdateUser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}
	msg := model.Message{Message: "updateUser"}
	return ctx.Status(fiber.StatusOK).JSON(msg)
}
