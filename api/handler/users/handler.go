package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service-api-cff/internal/logger"
	"service-api-cff/internal/models"
	"service-api-cff/internal/msgs"
	"service-api-cff/pkg/auth"
)

type handlerUsers struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerUsers) CreateUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseUsers{Error: true}
	srvUsers := auth.NewServerAuth(h.dB, h.user, h.txID)
	rqUsers := RequestUsers{}

	err := c.BodyParser(&rqUsers)
	if err != nil {
		logger.Error.Printf("couldn't bind model RequestMetadata: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	req, cod, err := srvUsers.Users.CreateUsers(rqUsers.Id, rqUsers.Username, rqUsers.CodeStudent, rqUsers.Dni, rqUsers.Names, rqUsers.LastnameFather, rqUsers.LastnameMother, rqUsers.Email, rqUsers.Password, 0, 0)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = *req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerUsers) GetUsers(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseUsersAll{Error: true}
	srvUsers := auth.NewServerAuth(h.dB, h.user, h.txID)

	req, err := srvUsers.Users.GetAllUsers()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerUsers) DeleteUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseUsersAll{Error: true}
	srvUsers := auth.NewServerAuth(h.dB, h.user, h.txID)

	id := c.Params("id")

	cod, err := srvUsers.Users.DeleteUsers(id)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = nil
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
