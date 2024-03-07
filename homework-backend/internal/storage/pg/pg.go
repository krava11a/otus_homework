package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"homework-backend/internal/models"
	"homework-backend/internal/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(pgConnectioString string) (*Storage, error) {
	const op = "storage.pg.New"

	database, err := sql.Open("postgres", pgConnectioString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: database}, nil
}

func (s *Storage) CreateUser(user models.User) (string, error) {
	const op = "storage.sqlite.CreateUser"

	_, err := s.db.Exec(`INSERT INTO users (first_name, second_name,birthdate,biography,city,hP) VALUES($1,$2,$3,$4,$5,$6)`,
		user.First_name, user.Second_name, user.Birthdate, user.Biography, user.City, user.Hp)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	uuid, err := s.getUUIDbyAllStaticFields(user)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return uuid, nil
}

func (s *Storage) getUUIDbyAllStaticFields(user models.User) (uuid string, err error) {

	resRow := s.db.QueryRow(`SELECT id FROM users WHERE first_name = $1 AND second_name = $2 AND birthdate = $3 AND biography=$4 AND city = $5 AND hP=$6`,
		user.First_name, user.Second_name, user.Birthdate, user.Biography, user.City, user.Hp)
	resRow.Scan(&uuid)
	return uuid, nil
}

func (s *Storage) GetUserById(user_id string) (models.User, error) {
	const op = "storage.sqlite.GetUserById"
	var user models.User
	// query := "SELECT id,first_name, second_name,birthdate,biography,city,hP FROM users WHERE id = $1"
	// stmt, err := s.db.Prepare()
	// if err != nil {
	// 	return models.User{}, fmt.Errorf("%s: %w", op, err)
	// }

	// resRow := stmt.Query(user_id)
	resRow := s.db.QueryRow(`SELECT id,first_name, second_name,birthdate,biography,city,hP FROM users WHERE id = $1`, user_id)

	// resRow := s.db.QueryRow(`SELECT * FROM users`)
	// if err != nil {
	// 	return models.User{}, fmt.Errorf("%s: %w", op, err)
	// }

	// var eUs models.User
	err := resRow.Scan(&user.Id, &user.First_name, &user.Second_name, &user.Birthdate, &user.Biography, &user.City, &user.Hp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil

}

func (s *Storage) App() (models.App, error) {
	return models.App{
		ID:     1,
		Name:   "otus_homework",
		Secret: "",
	}, nil
}
