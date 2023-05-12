package repository

import (
	"rest-api/pkg/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateRole(role models.Role) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	sqlRaw :=
		`INSERT INTO Roles (name) 
		VALUES ($1)
		RETURNING id`

	id := 0

	err = tx.QueryRow(sqlRaw,
		role.Name).Scan(&id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, nil

}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {

	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	sqlRaw :=
		`INSERT INTO Users (name, username, email, password, roleid)
		VALUES ($1, $2, $3, $4, (SELECT  id FROM Roles WHERE id = $5))
		RETURNING id`

	id := 0

	err = tx.QueryRow(sqlRaw,
		user.Name, user.Username, user.Email, user.Password, user.RoleID).Scan(&id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User

	sqlQuery := `SELECT * FROM Users WHERE username=$1 AND password=$2`

	err := r.db.Get(&user, sqlQuery, username, password)

	if err != nil {
		return models.User{}, err
	}

	return user, err

}
