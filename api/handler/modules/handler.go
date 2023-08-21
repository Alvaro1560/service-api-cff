package modules

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

type handlerModules struct {
	dB   *sqlx.DB
	user *models.User
	txID string
}

func (h *handlerModules) GetModulesUser(c *fiber.Ctx) error {
	msg := msgs.Model{}
	res := ResponseModules{Error: true}

	rqModulesUser := RequestModulesUser{}

	usr, err := middleware.GetUser(c)
	if err != nil {
		res.Error = true
		res.Code = 99
		res.Msg = "Error in token"
		return c.Status(http.StatusOK).JSON(res)
	}

	srvModules := auth.NewServerAuth(h.dB, usr, h.txID)

	err = c.BodyParser(&rqModulesUser)
	if err != nil {
		logger.Error.Printf("couldn't bind model rqModulesUser: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(1)
		res.Msg = "couldn't bind model RequestMetadata"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	var roles []string
	for _, r := range usr.Roles {
		roles = append(roles, *r)
	}

	req, err := srvModules.Modules.GetModulesByRoles(roles, rqModulesUser.Ids, rqModulesUser.Type)
	if err != nil {
		logger.Error.Printf("Couldn't insert suffragers: %v", err)
		res.Code, res.Type, res.Msg = msg.GetByCode(99)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false

	return c.Status(http.StatusOK).JSON(res)
}
