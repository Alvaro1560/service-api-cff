package entity

import (
	"github.com/jmoiron/sqlx"
	"service-api-cff/internal/models"
	"service-api-cff/pkg/entity/attendance"
)

type ServerEntity struct {
	Attendance attendance.PortsServerAttendance
}

func NewServerEntity(db *sqlx.DB, user *models.User, txID string) *ServerEntity {
	repoAttendance := attendance.FactoryStorage(db, user, txID)
	return &ServerEntity{
		Attendance: attendance.NewAttendanceService(repoAttendance, user, txID),
	}
}
