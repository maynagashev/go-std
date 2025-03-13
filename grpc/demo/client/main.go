package main

import (
	"context"
	pb "demo/proto"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// authInterceptor добавляет токен аутентификации ко всем исходящим gRPC запросам
func authInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// Добавляем токен в исходящий контекст
	md := metadata.New(map[string]string{"token": "secret_token"})
	newCtx := metadata.NewOutgoingContext(ctx, md)

	// Продолжаем вызов с использованием нового контекста
	return invoker(newCtx, method, req, reply, cc, opts...)
}

func main() {
	// устанавливаем соединение с сервером с использованием перехватчика
	conn, err := grpc.Dial(
		":3200",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(authInterceptor),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// получаем переменную интерфейсного типа UsersClient,
	// через которую будем отправлять сообщения
	c := pb.NewUsersClient(conn)

	// функция, в которой будем отправлять сообщения
	TestUsers(c)
}

func TestUsers(c pb.UsersClient) {
	// набор тестовых данных
	users := []*pb.User{
		{Name: "Сергей", Email: "serge@example.com", Sex: pb.User_MALE},
		{Name: "Света", Email: "sveta@example.com", Sex: pb.User_FEMALE},
		{Name: "Денис", Email: "den@example.com", Sex: pb.User_MALE},
		// при добавлении этой записи должна вернуться ошибка:
		// пользователь с email sveta@example.com уже существует
		{Name: "Sveta", Email: "sveta@example.com", Sex: pb.User_FEMALE},
	}
	for _, user := range users {
		// добавляем пользователей
		resp, err := c.AddUser(context.Background(), &pb.AddUserRequest{
			User: user,
		})
		if err != nil {
			log.Fatal(err)
		}
		if resp.Error != "" {
			fmt.Println(resp.Error)
		}
	}
	// удаляем одного из пользователей
	resp, err := c.DelUser(context.Background(), &pb.DelUserRequest{
		Email: "serge@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	if resp.Error != "" {
		fmt.Println(resp.Error)
	}
	// если запрос будет выполняться дольше 200 миллисекунд, то вернётся ошибка
	// с кодом codes.DeadlineExceeded и сообщением context deadline exceeded
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// получаем информацию о пользователях
	// во втором случае должна вернуться ошибка:
	// пользователь с email serge@example.com не найден
	for _, userEmail := range []string{"sveta@example.com", "serge@example.com"} {
		resp, err := c.GetUser(ctx, &pb.GetUserRequest{
			Email: userEmail,
		})
		if err != nil {
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					// выведет, что пользователь не найден
					fmt.Println(`NOT FOUND`, e.Message())
				} else {
					// в остальных случаях выводим код ошибки в виде строки и сообщение
					fmt.Println(e.Code(), e.Message())
				}
			} else {
				fmt.Printf("Не получилось распарсить ошибку %v", err)
			}
		} else {
			// Проверяем resp.Error только если нет ошибки
			if resp.Error == "" {
				fmt.Println(resp.User)
			} else {
				fmt.Println(resp.Error)
			}
		}
	}

	// получаем список email пользователей
	emails, err := c.ListUsers(context.Background(), &pb.ListUsersRequest{
		Offset: 0,
		Limit:  100,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emails.Count, emails.Emails)
}
