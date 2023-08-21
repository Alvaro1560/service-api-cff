package modules

import (
	"database/sql"
	"fmt"
	"service-api-cff/internal/helper"
	"service-api-cff/internal/logger"
	"service-api-cff/internal/models"

	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func NewModuleSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) Create(m *Module) error {
	const sqlInsert = `INSERT INTO auth.modules (id ,name, description, class, id_user, created_at, updated_at) VALUES (:id ,:name, :description, :class, :id_user, GetDate(), GetDate()) `
	m.IdUser = s.user.ID
	_, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't insert Module: %v", err)
		return err
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) Update(m *Module) error {
	const sqlUpdate = `UPDATE auth.modules SET name = :name, description = :description, class = :class, id_user =:id_user, updated_at = GetDate() WHERE id = :id `
	m.IdUser = s.user.ID
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't update Module: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *sqlserver) Delete(id string) error {
	m := Module{ID: id, IdUser: s.user.ID}
	const sqlDelete = `UPDATE auth.modules SET is_delete = 1, id_user =:id_user, updated_at = GetDate(), deleted_at = GetDate() WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		logger.Error.Printf(s.TxID, " - couldn't delete Module: %v", err)
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) GetByID(id string) (*Module, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , name, description, class, created_at, updated_at FROM auth.modules  WITH (NOLOCK)  WHERE id = @id AND is_delete = 0`
	mdl := Module{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetByID Module: %v", err)
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) GetAll() ([]*Module, error) {
	var ms []*Module
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , name, description, class, created_at, updated_at FROM auth.modules  WITH (NOLOCK)  WHERE  is_delete = 0`
	query := sqlGetAll
	err := s.DB.Select(&ms, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetAll auth.modules: %v", err)
		return ms, err
	}
	return ms, nil
}

// GetModulesByRoles consulta todos los registros de la BD
func (s *sqlserver) GetModulesByRoles(roleIDs []string, ids []string, typeArg int) ([]*Module, error) {
	var ms []*Module
	const sqlGetModulesByRoles = `SELECT DISTINCT convert(nvarchar(50), m.id) id , m.name, m.description, m.class, c.url_front path, m.created_at, m.updated_at FROM auth.modules m WITH (NOLOCK) 
				JOIN [auth].[modules_components] c WITH (NOLOCK)  ON (m.id = c.module_id)
				JOIN [auth].[modules_components_elements] e WITH (NOLOCK)  ON (c.id = e.component_id)
				JOIN [auth].[roles_elements] re  WITH (NOLOCK)  ON (e.id = re.element_id)
				WHERE lower(re.role_id) in (%s) AND m.type = @typeArg 
				AND m.is_delete = 0 AND c.is_delete = 0 AND e.is_delete = 0 AND re.is_delete = 0
`
	query := fmt.Sprintf(sqlGetModulesByRoles, helper.SliceToString(roleIDs))
	err := s.DB.Select(&ms, query, sql.Named("typeArg", typeArg))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error.Printf(s.TxID, " - couldn't execute GetModulesByRoles auth.modules: %v", err)
		return ms, err
	}
	return ms, nil
}
