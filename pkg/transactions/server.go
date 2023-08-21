package report

import (
	"github.com/jmoiron/sqlx"
	"service-api-cff/internal/models"
	"service-api-cff/pkg/transactions/report"
)

type ServerReport struct {
	Event report.PortsServerReport
}

func NewServerReport(db *sqlx.DB, user *models.User, txID string) *ServerReport {
	repoReport := report.FactoryStorage(db, user, txID)
	return &ServerReport{
		Event: report.NewReportService(repoReport, user, txID),
	}
}
