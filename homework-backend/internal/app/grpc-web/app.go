package grpcweb

import (
	"context"
	"fmt"
	"homework-backend/internal/lib/jwt"
	"homework-backend/internal/models"
	"homework-backend/internal/proto"
	"homework-backend/internal/storage/rabbit"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // Размер буфера чтения
	WriteBufferSize: 1024, // Размер буфера записи
	// Позволяет определить, должен ли сервер сжимать сообщения
	EnableCompression: true,
}

var QueuePath string
var Log *slog.Logger

func Run(grpcPort, grpcWebPort, wsPort uint, queuePath string, log *slog.Logger) {
	QueuePath = queuePath
	Log = log
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
		xid := request.Header.Get("X-Request-Id")
		md := metadata.Pairs("xid", xid)
		return md
	}))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := proto.RegisterPostServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", grpcPort), opts)
	if err != nil {
		log.Error("Error strating server: %v", err)
	}
	err = proto.RegisterAuthorizationServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", grpcPort), opts)
	if err != nil {
		log.Error("Error strating server: %v", err)
	}

	// Запустим HTTP сервер для проксирования gRPC вызовов
	log.Info("grpc-web server started", slog.String("addr", fmt.Sprintf(":%d", grpcWebPort)))

	// http.HandleFunc("/ws", handleConnections)
	go runWS(fmt.Sprintf(":%d", wsPort))
	http.ListenAndServe(fmt.Sprintf(":%d", grpcWebPort), mux)
}

func runWS(wsPort string) {
	http.HandleFunc("/ws", handler)

	err := http.ListenAndServe(wsPort, nil)
	if err != nil {
		log.Print("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// обновление соединения до WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
	}
	// defer ws.Close()
	auth := r.Header.Get("Authorization")
	token := auth[7:]
	uuid, _ := jwt.GetUserId(token, models.App{Secret: ""})
	go handleConnection(ws, uuid)
}

func handleConnection(ws *websocket.Conn, uuid string) {
	defer ws.Close()
	// цикл обработки сообщений
	for {

		rqueue, err := rabbit.New(QueuePath)

		if err != nil {
			Log.Error("Error strating server: %v", err)
		}
		msgs := rqueue.ReadFrom(uuid)
		// messageType, message, err := ws.ReadMessage()
		// if err != nil {
		// 	log.Println(err)
		// 	break
		// }
		// log.Printf("Received: %s", message)

		// эхо ансвер
		for _, msg := range msgs {
			if err := ws.WriteMessage(1, []byte(msg)); err != nil {
				log.Println(err)
				break
			}
		}

		time.Sleep(1 * time.Second)

	}
}
