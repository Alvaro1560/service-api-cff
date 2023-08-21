package users

import (
	"database/sql"
	"fmt"
	"service-api-cff/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

// psql estructura de conexión a la BD de postgresql
type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newUsersPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *psql) create(m *Users) error {
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const psqlInsert = `INSERT INTO auth.users (id ,username, code_student, dni, names, lastname_father, lastname_mother, email, password, is_delete, is_block, created_at, updated_at) VALUES (:id ,:username, :code_student, :dni, :names, :lastname_father, :lastname_mother, :email, :password, :is_delete, :is_block,:created_at, :updated_at) `
	rs, err := s.DB.NamedExec(psqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *psql) update(m *Users) error {
	date := time.Now()
	m.UpdatedAt = date
	const psqlUpdate = `UPDATE auth.users SET username = :username, code_student = :code_student, dni = :dni, names = :names, lastname_father = :lastname_father, lastname_mother = :lastname_mother, email = :email, password = :password, is_delete = :is_delete, is_block = :is_block, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(psqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *psql) delete(id string) error {
	const psqlDelete = `DELETE FROM auth.users WHERE id = :id `
	m := Users{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *psql) getByID(id string) (*Users, error) {
	const psqlGetByID = `SELECT id , username, code_student, dni, names, lastname_father, lastname_mother, email, password, is_delete, is_block, created_at, updated_at FROM auth.users WHERE id = $1 `
	mdl := Users{}
	err := s.DB.Get(&mdl, psqlGetByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *psql) getAll() ([]*Users, error) {
	var ms []*Users
	const psqlGetAll = ` SELECT id , username, code_student, dni, names, lastname_father, lastname_mother, email, password, is_delete, is_block, created_at, updated_at FROM auth.users `

	err := s.DB.Select(&ms, psqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *psql) getByCodeStudent(codeStudent string) (*Users, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , username, code_student, dni, names, lastname_father, lastname_mother, email, password, is_delete, is_block, created_at, updated_at FROM auth.users  WITH (NOLOCK)  WHERE code_student = @code_student `
	mdl := Users{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("code_student", codeStudent))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}
