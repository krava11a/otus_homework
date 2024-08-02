package dialog

import (
	"context"
	"fmt"
	"homework-dialogs/internal/models"
	"homework-dialogs/internal/proto"
	"homework-dialogs/internal/services/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GrpcDialog interface {
	Send(dialog models.Dialog, xid string) error
	List(from, to, xid string) (dialogs []models.Dialog, err error)
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
	xid := md.Get("xid")[0]
	token := auth[7:]
	id, err := ds.ars.GetUUIDBy(token, xid)
	if err != nil {
		return &proto.DialogSendResponse{
			Status:  500,
			Message: err.Error(),
		}, err
	}
	err = ds.grpcDialog.Send(models.Dialog{From: id, To: req.UserId, Text: req.Text}, xid)
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
	md, _ := metadata.FromIncomingContext(ctx)
	xid := md.Get("xid")[0]
	err := ds.grpcDialog.Send(models.Dialog{From: req.UserFrom, To: req.UserTo, Text: req.Text}, xid)
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
	xid := md.Get("xid")[0]
	token := auth[7:]
	id, err := ds.ars.GetUUIDBy(token, xid)
	if err != nil {
		return &proto.DialogListResponse{
			Status:  500,
			Message: fmt.Errorf("Error in request ID:%s. Error:%s", xid, err).Error(),
		}, err
	}

	msgs, err := ds.grpcDialog.List(id, req.UserId, xid)
	if err != nil {
		return &proto.DialogListResponse{
			Status:   500,
			Message:  fmt.Errorf("Error in request ID:%s. Error:%s", xid, err).Error(),
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
	md, _ := metadata.FromIncomingContext(ctx)
	xid := md.Get("xid")[0]
	msgs, err := ds.grpcDialog.List(req.UserFrom, req.UserTo, xid)
	if err != nil {
		return &proto.DialogListResponse{
			Status:   500,
			Message:  fmt.Errorf("Error in request ID:%s. Error:%s", xid, err).Error(),
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
