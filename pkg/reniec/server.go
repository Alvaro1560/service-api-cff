package reniec

import (
	"github.com/jmoiron/sqlx"
	"service-api-cff/internal/models"
	"service-api-cff/pkg/indra/upload_metadata"
	"service-api-cff/pkg/reniec/dni"
)

type ServerReniec struct {
	Dni dni.PortsServerReniec
}

func NewServerReniec(db *sqlx.DB, user *models.User, txID string) *ServerReniec {
	repoDni := upload_metadata.FactoryStorage(db, user, txID)
	return &ServerReniec{
		Dni: dni.NewReniecService(repoDni, user, txID),
	}
}
