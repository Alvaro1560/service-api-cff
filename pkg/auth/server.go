package auth

import (
	"github.com/jmoiron/sqlx"
	"service-api-cff/internal/models"
	"service-api-cff/pkg/auth/modules"
	"service-api-cff/pkg/auth/roles"
	"service-api-cff/pkg/auth/users"
)

type ServerAuth struct {
	Users   users.PortsServerUsers
	Modules modules.PortModules
	Roles   roles.PortRoles
}

func NewServerAuth(db *sqlx.DB, user *models.User, txID string) *ServerAuth {
	repoDni := users.FactoryStorage(db, user, txID)
	repoModules := modules.FactoryStorage(db, user, txID)
	repoRoles := roles.FactoryStorage(db, user, txID)
	return &ServerAuth{
		Users:   users.NewUsersService(repoDni, user, txID),
		Modules: modules.NewModuleService(repoModules, user, txID),
		Roles:   roles.NewRoleService(repoRoles, user, txID),
	}
}
