package repos

import (
	"MorselShogiew/Users-service-rest/logger"
	"MorselShogiew/Users-service-rest/models"
	"MorselShogiew/Users-service-rest/provider"

	"github.com/jmoiron/sqlx"
)

type DBRepo interface {
	AddUser(user models.User) error
	DeleteUser(id int) error
	GetUsers() (*[]models.User, error)
}

type DB struct {
	db *sqlx.DB
	l  logger.Logger
}

func NewDBRepo(p provider.Provider, l logger.Logger) DBRepo {
	return &DB{p.GetDBConn(), l}
}

func (r DB) AddUser(user models.User) error {
	var query = `INSERT INTO users(name,mail) values($1,$2);`
	if _, err := r.db.Exec(query, user.Name, user.Mail); err != nil {
		return err
	}
	return nil
}

func (r DB) DeleteUser(id int) error {
	var query = `DELETE FROM users WHERE user_id=($1);`
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (r DB) GetUsers() (*[]models.User, error) {
	rows, err := r.db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var res []models.User

	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.Id, &u.Name, &u.Mail)
		if err != nil {
			r.l.Error("error on select data", err)
			continue
		}
		res = append(res, u)
	}

	// if err := json.Unmarshal(data, &res); err != nil {
	// 	return nil, err
	// }
	return &res, nil
}
