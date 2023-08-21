package config

import (
	"github.com/jmoiron/sqlx"
	"service-api-cff/internal/models"
	"service-api-cff/pkg/config/events"
)

type ServerConfig struct {
	Event events.PortsServerEvents
}

func NewServerConfig(db *sqlx.DB, user *models.User, txID string) *ServerConfig {
	repoEvent := events.FactoryStorage(db, user, txID)
	return &ServerConfig{
		Event: events.NewEventsService(repoEvent, user, txID),
	}
}
