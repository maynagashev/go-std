package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sort"
	"sync"

	// импортируем пакет со сгенерированными protobuf-файлами
	pb "demo/proto"
)

const SecretToken = "secret_token"

// UsersServer поддерживает все необходимые методы сервера.
type UsersServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedUsersServer

	// используем sync.Map для хранения пользователей
	users sync.Map
}

// AddUser реализует интерфейс добавления пользователя.
func (s *UsersServer) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	var response pb.AddUserResponse

	if _, ok := s.users.Load(in.User.Email); ok {
		response.Error = fmt.Sprintf("Пользователь с email %s уже существует", in.User.Email)
	} else {
		s.users.Store(in.User.Email, in.User)
	}
	return &response, nil
}

// ListUsers реализует интерфейс получения списка пользователей.
func (s *UsersServer) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var list []string

	s.users.Range(func(key, _ interface{}) bool {
		list = append(list, key.(string))
		return true
	})
	// сортируем слайс из email
	sort.Strings(list)

	offset := int(in.Offset)
	end := int(in.Offset + in.Limit)
	if end > len(list) {
		end = len(list)
	}
	if offset >= end {
		offset = 0
		end = 0
	}
	response := pb.ListUsersResponse{
		Count:  int32(len(list)),
		Emails: list[offset:end],
	}
	return &response, nil
}

// GetUser реализует интерфейс получения информации о пользователе.
func (s *UsersServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var response pb.GetUserResponse
	var token string

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md.Get("token")
		if len(values) > 0 {
			// ключ содержит слайс строк, получаем первую строку
			token = values[0]
		}
	}
	fmt.Println("received token", token)

	if user, ok := s.users.Load(in.Email); ok {
		response.User = user.(*pb.User)
	} else {
		return nil, status.Errorf(codes.NotFound, `Пользователь с email %s не найден`, in.Email)
	}
	return &response, nil
}

// DelUser реализует интерфейс удаления информации о пользователе.
func (s *UsersServer) DelUser(ctx context.Context, in *pb.DelUserRequest) (*pb.DelUserResponse, error) {
	var response pb.DelUserResponse

	if _, ok := s.users.LoadAndDelete(in.Email); !ok {
		response.Error = fmt.Sprintf("Пользователь с email %s не найден", in.Email)
	}
	return &response, nil
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var token string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("token")
		if len(values) > 0 {
			token = values[0]
		}
	}
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}
	if token != SecretToken {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	return handler(ctx, req)
}

func main() {
	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	// регистрируем сервис
	pb.RegisterUsersServer(s, &UsersServer{})

	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
