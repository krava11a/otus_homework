package dialog

import (
	"context"
	"homework-dialogs/internal/models"
	"homework-dialogs/internal/proto"
	"homework-dialogs/internal/services/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GrpcDialog interface {
	Send(dialog models.Dialog) error
	List(from, to string) (dialogs []models.Dialog, err error)
}

type DialogServer struct {
	proto.UnimplementedDialogServiceServer
	grpcDialog GrpcDialog
	ars        *auth.AuthRemoteService
}

func Register(gRPCServer *grpc.Server, grpcDialog GrpcDialog, ars *auth.AuthRemoteService) {
	proto.RegisterDialogServiceServer(gRPCServer, &DialogServer{grpcDialog: grpcDialog, ars: ars})
}

func (ds *DialogServer) Send(ctx context.Context, req *proto.DialogSendRequest) (*proto.DialogSendResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	auth := md.Get("authorization")[0]
	token := auth[7:]
	id, err := ds.ars.GetUUIDBy(token)
	if err != nil {
		return &proto.DialogSendResponse{
			Status:  500,
			Message: err.Error(),
		}, err
	}
	err = ds.grpcDialog.Send(models.Dialog{From: id, To: req.UserId, Text: req.Text})
	if err != nil {
		return &proto.DialogSendResponse{
			Status:  500,
			Message: err.Error(),
		}, err
	}

	return &proto.DialogSendResponse{
		Status:  200,
		Message: "Сообщение отправлено",
	}, err
}

func (ds *DialogServer) SendWithoutToken(ctx context.Context, req *proto.DialogSendWTRequest) (*proto.DialogSendResponse, error) {

	err := ds.grpcDialog.Send(models.Dialog{From: req.UserFrom, To: req.UserTo, Text: req.Text})
	if err != nil {
		return &proto.DialogSendResponse{
			Status:  500,
			Message: err.Error(),
		}, err
	}

	return &proto.DialogSendResponse{
		Status:  200,
		Message: "Сообщение отправлено",
	}, err
}

func (ds *DialogServer) List(ctx context.Context, req *proto.DialogListRequest) (*proto.DialogListResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	auth := md.Get("authorization")[0]
	token := auth[7:]
	id, err := ds.ars.GetUUIDBy(token)
	if err != nil {
		return &proto.DialogListResponse{
			Status:  500,
			Message: err.Error(),
		}, err
	}

	msgs, err := ds.grpcDialog.List(id, req.UserId)
	if err != nil {
		return &proto.DialogListResponse{
			Status:   500,
			Message:  err.Error(),
			Messages: []*proto.DialogMessage{},
		}, err
	}

	var resMsgs = []*proto.DialogMessage{}
	for _, m := range msgs {
		resMsgs = append(resMsgs, &proto.DialogMessage{
			From: m.From,
			To:   m.To,
			Text: m.Text,
		})
	}

	return &proto.DialogListResponse{
		Status:   200,
		Message:  "Все хорошо",
		Messages: resMsgs,
	}, err
}

func (ds *DialogServer) ListWithoutToken(ctx context.Context, req *proto.DialogListWTRequest) (*proto.DialogListResponse, error) {

	msgs, err := ds.grpcDialog.List(req.UserFrom, req.UserTo)
	if err != nil {
		return &proto.DialogListResponse{
			Status:   500,
			Message:  err.Error(),
			Messages: []*proto.DialogMessage{},
		}, err
	}

	var resMsgs = []*proto.DialogMessage{}
	for _, m := range msgs {
		resMsgs = append(resMsgs, &proto.DialogMessage{
			From: m.From,
			To:   m.To,
			Text: m.Text,
		})
	}

	return &proto.DialogListResponse{
		Status:   200,
		Message:  "Все хорошо",
		Messages: resMsgs,
	}, err
}
