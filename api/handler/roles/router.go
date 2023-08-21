package Roles

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterRoles(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerRoles{dB: db, txID: txID}
	api := app.Group("/api")
	v1 := api.Group("/v1")
	getWork := v1.Group("/roles")
	getWork.Get("", h.GetRoles)
	getWork.Post("create", h.CreateRoles)
}
