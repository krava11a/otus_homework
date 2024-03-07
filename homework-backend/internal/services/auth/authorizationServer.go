package auth

import (
	"context"
	"homework-backend/internal/models"
	"homework-backend/internal/proto"

	"google.golang.org/grpc"
)

type GrpcAuth interface {
	CreateUser(user models.User) (string, error)
	LoginUser(user_id, password string, app models.App) (string, error)
	GetUserById(user_id string) (models.User, error)
}

func Register(gRPCServer *grpc.Server, auth GrpcAuth) {
	proto.RegisterAuthorizationServiceServer(gRPCServer, &AuthorizationServer{auth: auth})
}

type AuthorizationServer struct {
	proto.UnimplementedAuthorizationServiceServer
	auth GrpcAuth
}

func (as *AuthorizationServer) UserRegister(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	nUser := models.User{
		First_name:  req.FirstName,
		Second_name: req.SecondName,
		Birthdate:   req.Birthdate,
		Biography:   req.Biography,
		City:        req.City,
		Password:    req.Password,
	}

	uuid, err := as.auth.CreateUser(nUser)
	if err != nil {
		return &proto.CreateUserResponse{
			Status:  400,
			Message: "Неправильные данные",
			UserId:  uuid,
		}, err
	}
	return &proto.CreateUserResponse{
		Status:  200,
		Message: "Все хорошо",
		UserId:  uuid,
	}, nil

}

func (as *AuthorizationServer) UserLogin(ctx context.Context, req *proto.UserLoginRequest) (*proto.UserLoginResponse, error) {
	app := models.App{
		ID:     1,
		Name:   "otus_homework",
		Secret: "",
	}

	token, err := as.auth.LoginUser(req.UserId, req.Password, app)
	return &proto.UserLoginResponse{
		Status:  200,
		Message: "Все хорошо",
		Token:   token,
	}, err
}

func (as *AuthorizationServer) UserGetById(ctx context.Context, req *proto.UserIdRequest) (*proto.UserGetByIdResponse, error) {
	user, err := as.auth.GetUserById(req.UserId)
	if err != nil {
		return &proto.UserGetByIdResponse{
			Status:  500,
			Message: err.Error(),
			User:    nil,
		}, err
	}
	return &proto.UserGetByIdResponse{
		Status:  200,
		Message: "Все хорошо",
		User: &proto.User{
			Id:         user.Id,
			FirstName:  user.First_name,
			SecondName: user.Second_name,
			Birthdate:  user.Birthdate,
			Biography:  user.Biography,
			City:       user.City,
		},
	}, nil
}
