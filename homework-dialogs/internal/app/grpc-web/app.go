package grpcweb

import (
	"context"
	"fmt"
	"homework-dialogs/internal/proto"

	"log/slog"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var Log *slog.Logger

func Run(grpcPort, grpcWebPort, wsPort uint, log *slog.Logger) {

	Log = log
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
		header := request.Header.Get("Authorization")
		// send all the headers received from the client
		xid := request.Header.Get("X-Request-ID")
		md := metadata.Pairs("auth", header, "xid", xid)
		// md = metadata.Pairs("X-Request-ID", xid)
		// fmt.Println(xid)
		return md
	}))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := proto.RegisterDialogServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", grpcPort), opts)
	if err != nil {
		log.Error("Error strating server: %v", err)
	}

	// Запустим HTTP сервер для проксирования gRPC вызовов
	log.Info("grpc-web server started", slog.String("addr", fmt.Sprintf(":%d", grpcWebPort)))

	if err != nil {
		panic(err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", grpcWebPort), mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)

	// uuid, _ := jwt.GetUserId(token, models.App{Secret: ""})
	// go handleConnection(ws, uuid)

}
