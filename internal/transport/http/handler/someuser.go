package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mellgit/someuser/internal/config"
	"github.com/mellgit/someuser/internal/model"
	"github.com/mellgit/someuser/internal/service"
	log "github.com/sirupsen/logrus"
)

const someUserPath = "/api/v1/someusers"

type SomeUser struct {
	SomeUserService service.Service
	Cfg             *config.Config
	Logger          *log.Entry
}

func NewSomeUser(cfg *config.Config, someUserService service.Service, logger *log.Entry) *SomeUser {
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

// CreateUser
// @Summary      CreateUser
// @Description  Create new user
// @Tags         SomeUsers
// @Accept       json
// @Produce      json
// @Param 		 request body model.CreateUserRequest true "body"
// @Success      200 {object} model.SchemaSomeUser
// @Failure      400 {object} model.Error
// @Failure      404 {object} model.Error
// @Failure      500 {object} model.Error
// @Router       /someusers/ [post]
func (h *SomeUser) CreateUser(ctx *fiber.Ctx) error {
	payload := model.CreateUserRequest{}
	if err := ctx.BodyParser(&payload); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "ctx.BodyParser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(msgErr)

	}

	user, err := h.SomeUserService.CreateUser(payload)
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.CreateUser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(msgErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

// GetAllUsers
// @Summary      GetAllUsers
// @Description  Get all users
// @Tags         SomeUsers
// @Accept       json
// @Produce      json
// @Success      200 {array} model.SchemaSomeUser
// @Failure      400 {object} model.Error
// @Failure      404 {object} model.Error
// @Failure      500 {object} model.Error
// @Router       /someusers/ [get]
func (h *SomeUser) GetAllUsers(ctx *fiber.Ctx) error {

	allUsers, err := h.SomeUserService.GetAllUsers()
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.GetAllUsers",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(msgErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(allUsers)
}

// GetUserByID
// @Summary      GetUserByID
// @Description  Get one user by id
// @Tags         SomeUsers
// @Accept       json
// @Produce      json
// @Param        id path string true "UUID or string"
// @Success      200 {object} model.SchemaSomeUser
// @Failure      400 {object} model.Error
// @Failure      404 {object} model.Error
// @Failure      503 {object} model.Error
// @Router       /someusers/{id} [get]
func (h *SomeUser) GetUserByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	user, err := h.SomeUserService.GetUserByID(id)
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.GetUserByID",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(msgErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

// DeleteUserByID
// @Summary      DeleteUserByID
// @Description  Delete one user by id
// @Tags         SomeUsers
// @Accept       json
// @Produce      json
// @Param        id path string true "UUID"
// @Success      200 {object} model.Message
// @Failure      400 {object} model.Error
// @Failure      404 {object} model.Error
// @Failure      503 {object} model.Error
// @Router       /someusers/{id} [delete]
func (h *SomeUser) DeleteUserByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	if err := h.SomeUserService.DeleteUserByID(id); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.DeleteUserByID",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(msgErr)
	}
	msg := model.Message{Message: fmt.Sprintf("successfully deleted user with id %s", id)}
	return ctx.Status(fiber.StatusOK).JSON(msg)
}

// UpdateUser
// @Summary      UpdateUser
// @Description  Update full record by id
// @Tags         SomeUsers
// @Accept       json
// @Produce      json
// @Param        id path string true "UUID"
// @Param 		 request body model.UpdateUserRequest true "body"
// @Success      200 {object} model.SchemaSomeUser
// @Failure      400 {object} model.Error
// @Failure      404 {object} model.Error
// @Failure      503 {object} model.Error
// @Router       /someusers/{id} [patch]
func (h *SomeUser) UpdateUser(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	payload := model.UpdateUserRequest{}
	if err := ctx.BodyParser(&payload); err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "ctx.BodyParser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(msgErr)
	}
	user, err := h.SomeUserService.UpdateUser(id, payload)
	if err != nil {
		h.Logger.WithFields(log.Fields{
			"action": "SomeUserService.UpdateUser",
		}).Errorf("%v", err)
		msgErr := model.Error{Error: err.Error()}
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(msgErr)
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
