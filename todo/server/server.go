package main

import (
	pb "github.com/Nu1LSpaxe/grpc-go-for-professionals/todo/proto/v2"
)

type server struct {
	d db

	pb.UnimplementedTodoServiceServer
}
