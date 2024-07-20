package pgx

import (
	"context"
	"fmt"
	"homework-dialogs/internal/models"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	dbpool *pgxpool.Pool
}

func New(pgConnectioString string) (*Storage, error) {
	const op = "storage.pgx.New"

	dbpool, err := pgxpool.New(context.Background(), pgConnectioString)
	if err != nil {
		return &Storage{}, fmt.Errorf("Unable to create connection pool: %v\n", err)
		// os.Exit(1)
	}
	s := Storage{dbpool: dbpool}
	// err = s.InitDb()
	// if err != nil {
	// 	return &s, fmt.Errorf("Unable to Init DB Dialogs: %v\n", err)
	// }

	err = s.initTable()
	if err != nil {
		return &s, fmt.Errorf("Unable to Init Table Dialogs: %v\n", err)
	}

	return &s, nil
}

func (s *Storage) InitDb() error {
	const op = "storage.pgx.InitDB"

	dbExits := ""
	resRow := s.dbpool.QueryRow(context.Background(), `SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower('dialogs');`)
	resRow.Scan(&dbExits)

	if strings.Compare(dbExits, "dialogs") != 0 {
		_, err := s.dbpool.Exec(context.Background(), `CREATE DATABASE dialogs`)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

	}

	return nil
}

func (s *Storage) initTable() error {
	const op = "storage.pgx.InitTable"

	tExits := ""
	resRow := s.dbpool.QueryRow(context.Background(), `SELECT table_name FROM information_schema.tables WHERE table_name='dialogs';`)
	resRow.Scan(&tExits)

	if strings.Compare(tExits, "dialogs") != 0 {
		_, err := s.dbpool.Exec(context.Background(), `CREATE TABLE dialogs (id SERIAL PRIMARY KEY,"from" UUID NOT NULL, "to" UUID NOT NULL, "text" VARCHAR(255) NOT NULL, "timestamp"  DATE NOT NULL);`)

		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil

}

func (s *Storage) Send(dialog models.Dialog) error {
	const op = "storage.pgx.DialogSendMsg"

	_, err := s.dbpool.Exec(context.Background(), `INSERT INTO dialogs ("from", "to", text, timestamp) VALUES($1,$2,$3,CURRENT_TIMESTAMP)`, dialog.From, dialog.To, dialog.Text)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) List(from, to string) ([]models.Dialog, error) {
	const op = "storage.pgx.DialogListMsgs"

	rows, err := s.dbpool.Query(context.Background(), `SELECT "from", "to", text FROM dialogs WHERE "from" = $1 AND "to" = $2 ORDER BY timestamp `, from, to)
	if err != nil {
		return []models.Dialog{}, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	m := models.Dialog{}
	msgs := make([]models.Dialog, 0, 100)

	for rows.Next() {

		err := rows.Scan(&m.From, &m.To, &m.Text)
		if err != nil {
			return []models.Dialog{}, err
		}
		msgs = append(msgs, m)

	}

	return msgs, nil
}
