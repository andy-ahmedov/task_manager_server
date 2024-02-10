package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/andy-ahmedov/task_manager_grpc_api/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv    *grpc.Server
	taskServer api.TaskServiceServer
}

func New(taskServ api.TaskServiceServer) *Server {
	return &Server{
		grpcSrv:    grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor)),
		taskServer: taskServ,
	}
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n", info.FullMethod, time.Since(start), err)
	return resp, err
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	api.RegisterTaskServiceServer(s.grpcSrv, s.taskServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return err
}
