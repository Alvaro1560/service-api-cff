package upload_metadata

import (
	"github.com/jmoiron/sqlx"
	"service-api-cff/internal/logger"
	"service-api-cff/internal/models"
)

const (
	Postgresql = "postgres"
)

type ServicesUploadMetadataRepository interface {
	updateMetadata(metadata []Metadata) (int, error)
	GetIdsAutofillValueByEntityAttributeAndValue(typeInput string, inputData string) ([]*Metadata, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUploadMetadataRepository {
	var s ServicesUploadMetadataRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newUploadMetadataPsqlRepository(db, user, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
