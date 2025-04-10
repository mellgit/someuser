package handler

import (
	"fmt"
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

	user, err := h.SomeUserService.CreateUser(payload)
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.CreateUser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (h *SomeUser) GetAllUsers(ctx *fiber.Ctx) error {

	allUsers, err := h.SomeUserService.GetAllUsers()
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.GetAllUsers",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(allUsers)
}

func (h *SomeUser) GetUserByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	user, err := h.SomeUserService.GetUserByID(uuid.MustParse(id))
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.GetUserByID",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
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
	msg := model.Message{Message: fmt.Sprintf("successfully deleted user with id %s", id)}
	return ctx.Status(fiber.StatusOK).JSON(msg)
}

func (h *SomeUser) UpdateUser(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	payload := model.UpdateUserRequest{}
	if err := ctx.BodyParser(&payload); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "ctx.BodyParser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}
	user, err := h.SomeUserService.UpdateUser(uuid.MustParse(id), payload)
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.UpdateUser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusInternalServerError).JSON(msgErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
