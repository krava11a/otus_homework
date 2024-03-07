package auth

import (
	"errors"
	"fmt"

	"homework-backend/internal/lib/jwt"
	"homework-backend/internal/lib/logger/sl"
	"homework-backend/internal/models"
	"homework-backend/internal/storage"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Auth struct {
	log         *slog.Logger
	usrCreater  UserCreater
	usrProvider UserProvider
	appProvider AppProvider
	tokenTTL    time.Duration
}

func New(
	log *slog.Logger,
	userCreater UserCreater,
	userProvider UserProvider,
	appProvider AppProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		log:         log,
		usrCreater:  userCreater,
		usrProvider: userProvider,
		appProvider: appProvider,
		tokenTTL:    tokenTTL,
	}
}

func (a *Auth) CreateUser(user models.User) (string, error) {

	const op = "createUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("first name", user.First_name),
	)

	log.Info("registering user")

	if user.First_name == "" {
		return "", status.Error(codes.InvalidArgument, "First Name is required")
	}

	if user.Birthdate == "" {
		return "", status.Error(codes.InvalidArgument, "Birthdate is required")
	}

	if user.Biography == "" {
		return "", status.Error(codes.InvalidArgument, "Biography is required")
	}
	if user.City == "" {
		return "", status.Error(codes.InvalidArgument, "City is required")
	}
	if user.Password == "" {
		return "", status.Error(codes.InvalidArgument, "Password is required")
	}

	// err := OpenConnectionDb()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }

	// defer CloseConnectionDb()

	hp, err := HashPassword(user.Password)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	user.Hp = hp

	// _, err = Db.ExecContext(context.Background(),
	// 	`INSERT INTO users (first_name, second_name,birthdate,biography,city,hP) VALUES($1,$2,$3,$4,$5,$6);`,
	// 	user.First_name, user.Second_name, user.Birthdate, user.Biography, user.City, hp)
	// CheckError(err)
	// user.Password = hp
	// uuid, err := getUUIDbyAllStaticFields(user)
	// CheckError(err)

	uuid, err := a.usrCreater.CreateUser(user)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}
	return uuid, nil
}

// func getUUIDbyAllStaticFields(user models.User) (uuid string, err error) {

// 	resRow := Db.QueryRowContext(context.Background(), `SELECT id FROM users WHERE first_name = $1 AND second_name = $2 AND birthdate = $3 AND biography=$4 AND city = $5 AND hP=$6`,
// 		user.First_name, user.Second_name, user.Birthdate, user.Biography, user.City, user.Password)
// 	resRow.Scan(&uuid)
// 	return uuid, nil
// }

// func GetUserById(user_id string) (user models.User, err error) {
// 	err = OpenConnectionDb()
// 	if err != nil {
// 		fmt.Println(err)
// 		return models.User{}, err
// 	}

// 	defer CloseConnectionDb()

// 	resRow := Db.QueryRowContext(context.Background(), `SELECT id,first_name, second_name,birthdate,biography,city,hP FROM users WHERE id = $1`)
// 	resRow.Scan(&user)
// 	return user, nil

// }

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 11)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *Auth) LoginUser(user_id, password string, app models.App) (string, error) {

	const op = "Auth.Login"
	log := a.log.With(
		slog.String("op", op),
		slog.String("user_id", user_id),
	)
	log.Info("attempting to login user")

	if user_id == "" {
		return "", status.Error(codes.InvalidArgument, "user_id is required")
	}

	if password == "" {
		return "", status.Error(codes.InvalidArgument, "password is required")
	}

	user, err := a.usrProvider.GetUserById(user_id)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			a.log.Warn("user not found", sl.Err(err))

			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		a.log.Error("failed to get user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	// eUser, err := GetUserById(user_id)
	// CheckError(err)

	if !CheckPasswordHash(password, user.Hp) {
		a.log.Info("invalid credentials")
		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")

	token, err := jwt.NewToken(user_id, app, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (a *Auth) GetUserById(user_id string) (user models.User, err error) {
	const op = "Auth.Login"
	log := a.log.With(
		slog.String("op", op),
		slog.String("user_id", user_id),
	)
	log.Info("attempting to GET user by ID")

	if user_id == "" {
		return models.User{}, status.Error(codes.InvalidArgument, "user_id is required")
	}

	user, err = a.usrProvider.GetUserById(user_id)
	if err != nil {
		a.log.Error("failed to get user by id", sl.Err(err))

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil

}
