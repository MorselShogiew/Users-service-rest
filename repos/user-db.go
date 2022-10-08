package repos

import (
	"MorselShogiew/Users-service-rest/logger"
	"MorselShogiew/Users-service-rest/models"
	"MorselShogiew/Users-service-rest/provider"
	"encoding/json"

	"github.com/jmoiron/sqlx"
)

type DBRepo interface {
	AddUser(user models.User) error
	DeleteUser(id int) error
	GetUsers() (*[]models.User, error)
}

type DB struct {
	db *sqlx.DB
	logger.Logger
}

func NewDBRepo(p provider.Provider, l logger.Logger) DBRepo {
	return &DB{p.GetDBConn(), l}
}

func (r DB) AddUser(user models.User) error {
	var query = `INSERT INTO user(id,name,mail) values($1,$2,$3);`
	if _, err := r.db.Exec(query, user.Id, user.Name, user.Mail); err != nil {
		return err
	}
	return nil
}

func (r DB) DeleteUser(id int) error {
	var query = `DELETE FROM user WHERE id=($1);`
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (r DB) GetUsers() (*[]models.User, error) {
	var data []byte
	var query = `select * from user;`
	if err := r.db.QueryRowx(query).Scan(&data); err != nil {
		return nil, err
	}

	var res []models.User
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
