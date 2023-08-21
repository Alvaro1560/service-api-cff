package indra

import (
	"github.com/jmoiron/sqlx"
	"service-api-cff/internal/models"
	"service-api-cff/pkg/indra/upload_metadata"
)

type ServerIndra struct {
	SrvUploadMetadata upload_metadata.PortsServerUploadMetadata
}

func NewServerIndra(db *sqlx.DB, user *models.User, txID string) *ServerIndra {
	repoUploadMetadata := upload_metadata.FactoryStorage(db, user, txID)
	return &ServerIndra{
		SrvUploadMetadata: upload_metadata.NewUploadMetadataService(repoUploadMetadata, user, txID),
	}
}
