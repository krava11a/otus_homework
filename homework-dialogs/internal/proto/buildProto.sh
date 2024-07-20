#!/bin/bash
protoc --go_out=. --go_opt=paths=source_relative   --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./autorization.proto 
protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative   --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./post.proto
protoc --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative   --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./dialogs.proto