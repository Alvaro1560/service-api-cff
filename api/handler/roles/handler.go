package Roles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service-api-cff/internal/logger"
	"service-api-cff/internal/middleware"
	"service-api-cff/internal/models"
	"service-api-cff/internal/msgs"
	"service-api-cff/pkg/auth"
)

type handlerRoles struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerRoles) CreateRoles(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseRoles{Error: true}

	rqARoles := RequestRoles{}

	err := c.BodyParser(&rqARoles)
	if err != nil {
		logger.Error.Printf("couldn't bind model BodyParser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}

	srvAuth := auth.NewServerAuth(h.dB, usr, h.txID)

	req, cod, err := srvAuth.Roles.CreateRole(rqARoles.Id, rqARoles.Name, rqARoles.Description, 1, true)
	if err != nil {
		logger.Error.Printf("Couldn't insert CreateRoles: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(cod)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerRoles) GetRoles(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseAllRoles{Error: true}
	srvAuth := auth.NewServerAuth(h.dB, h.user, h.txID)

	req, err := srvAuth.Roles.GetAllRole()
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
