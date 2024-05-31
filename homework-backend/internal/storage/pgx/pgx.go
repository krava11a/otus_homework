package pgx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"homework-backend/internal/models"
	"homework-backend/internal/storage"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	dbpool *pgxpool.Pool
}

func New(pgConnectioString string) (*Storage, error) {
	const op = "storage.pgx.New"

	dbpool, err := pgxpool.New(context.Background(), pgConnectioString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return &Storage{dbpool: dbpool}, nil
}

func (s *Storage) CreateUser(user models.User) (string, error) {
	const op = "storage.pgx.CreateUser"

	_, err := s.dbpool.Exec(context.Background(), `INSERT INTO users (first_name, second_name,birthdate,biography,city,hP) VALUES($1,$2,$3,$4,$5,$6)`,
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

	resRow := s.dbpool.QueryRow(context.Background(), `SELECT id FROM users WHERE first_name = $1 AND second_name = $2 AND birthdate = $3 AND biography=$4 AND city = $5 AND hP=$6`,
		user.First_name, user.Second_name, user.Birthdate, user.Biography, user.City, user.Hp)
	resRow.Scan(&uuid)
	return uuid, nil
}

func (s *Storage) GetUserById(user_id string) (models.User, error) {
	const op = "storage.pgx.GetUserById"
	var user models.User
	// query := "SELECT id,first_name, second_name,birthdate,biography,city,hP FROM users WHERE id = $1"
	// stmt, err := s.db.Prepare()
	// if err != nil {
	// 	return models.User{}, fmt.Errorf("%s: %w", op, err)
	// }

	// resRow := stmt.Query(user_id)
	resRow := s.dbpool.QueryRow(context.Background(), `SELECT id,first_name, second_name,TO_CHAR(birthdate, 'yyyy-mm-dd'),biography,city,hP FROM users WHERE id = $1`, user_id)

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

func (s *Storage) UsersGetByPrefixFirstNameAndSecondName(first_name, second_name string) ([]models.User, error) {

	// fmt.Println(s.db)

	const op = "storage.pgx.UsersGetByPrefixFirstNameAndSecondName"
	// urlExample := "postgres://postgres:example@localhost:5432/otus_homework"

	// conn, err := pgx.Connect(context.Background(), urlExample)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	// resRows, err := conn.Query(context.Background(), `SELECT id,first_name, second_name,TO_CHAR(birthdate, 'yyyy-mm-dd'),biography,city,hP FROM users WHERE first_name LIKE $1||'%' AND second_name LIKE $2||'%' ORDER BY id`, first_name, second_name)
	resRows, err := s.dbpool.Query(context.Background(), `SELECT id,first_name, second_name,TO_CHAR(birthdate, 'yyyy-mm-dd'),biography,city,hP FROM users WHERE first_name LIKE $1||'%' AND second_name LIKE $2||'%' ORDER BY id`, first_name, second_name)
	// if errors.Is(err, sql.ErrNoRows) {
	// 	return []models.User{}, err
	// }
	if err != nil {
		return []models.User{}, err
	}
	defer resRows.Close()
	users := make([]models.User, 0, 100)
	for resRows.Next() {
		var user models.User
		err := resRows.Scan(&user.Id, &user.First_name, &user.Second_name, &user.Birthdate, &user.Biography, &user.City, &user.Hp)
		// err := resRows.Scan(&user)
		if err != nil {
			return []models.User{}, err
		}
		users = append(users, user)

	}

	return users, nil
}

func (s *Storage) App() (models.App, error) {
	return models.App{
		ID:     1,
		Name:   "otus_homework",
		Secret: "",
	}, nil
}

func (s *Storage) FriendSet(user_id, friend_id string) error {
	const op = "storage.pgx.FriendSet"

	_, err := s.dbpool.Exec(context.Background(), `INSERT INTO friends (id,id_friend) VALUES($1,$2)`,
		user_id, friend_id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) FriendDelete(user_id, friend_id string) error {
	const op = "storage.pgx.FriendDelete"

	_, err := s.dbpool.Exec(context.Background(), `DELETE FROM friends WHERE id= $1  AND id_friend = $2`,
		user_id, friend_id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
