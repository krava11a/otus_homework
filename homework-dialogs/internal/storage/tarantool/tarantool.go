package tarantool

import (
	"context"
	"fmt"
	"homework-dialogs/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/tarantool/go-tarantool/v2"
	"github.com/tarantool/go-tarantool/v2/datetime"
	_ "github.com/tarantool/go-tarantool/v2/datetime"
	_ "github.com/tarantool/go-tarantool/v2/decimal"
	_ "github.com/tarantool/go-tarantool/v2/uuid"
)

type Storage struct {
	addr string
}

type tarantoolDialogs struct {
	Ds []tDialog
}

type tDialog struct {
	Id   uuid.UUID         `json:"id"`
	From uuid.UUID         `json:"from"`
	To   uuid.UUID         `json:"to"`
	Text string            `json:"text"`
	Date datetime.Datetime `json:"timestamp"`
}

func New(taraConnectionAddr string) (*Storage, error) {
	// Connect to the database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	dialer := tarantool.NetDialer{
		Address: taraConnectionAddr,
		// User:     "sampleuser",
		// Password: "123456",
	}
	opts := tarantool.Opts{
		Timeout: time.Second,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {
		return nil, fmt.Errorf("Tarantool Connection refused:", err)
	}

	defer conn.CloseGraceful()

	return &Storage{addr: taraConnectionAddr}, nil
}

func (s *Storage) Send(dialog models.Dialog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(dialog)
	dialer := tarantool.NetDialer{
		Address: s.addr,
		// User:     "sampleuser",
		// Password: "123456",
	}
	opts := tarantool.Opts{
		Timeout: time.Second,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {

		return fmt.Errorf("Tarantool Connection refused:", err)
	}

	defer conn.CloseGraceful()

	data, err := conn.Do(
		tarantool.NewCallRequest("send").Args([]interface{}{dialog.From, dialog.To, dialog.Text}),
	).Get()
	if err != nil {
		if err.Error() != "msgpack: unknown ext id=2" {
			return fmt.Errorf("Tarantool Got an error:", err)
		}

	}
	fmt.Println("Tarantool Stored procedure result:", data)
	return nil
}

func (s *Storage) List(from, to string) ([]models.Dialog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	dialer := tarantool.NetDialer{
		Address: s.addr,
		// User:     "sampleuser",
		// Password: "123456",
	}
	opts := tarantool.Opts{
		Timeout: time.Second,
	}

	conn, err := tarantool.Connect(ctx, dialer, opts)
	if err != nil {

		return nil, fmt.Errorf("Tarantool Connection refused:", err)
	}

	defer conn.CloseGraceful()
	res := tarantoolDialogs{}
	err = conn.Do(
		tarantool.NewCallRequest("list").Args([]interface{}{from, to}),
	).GetTyped(&res)
	if err != nil {
		return nil, fmt.Errorf("Tarantool Got an error:", err)
	}
	r := []models.Dialog{}
	for _, v := range res.Ds {
		r = append(r, models.Dialog{
			From: v.From.String(),
			To:   v.From.String(),
			Text: v.Text,
		})
	}

	return r, nil
}
