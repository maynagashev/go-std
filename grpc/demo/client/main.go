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

// createAuthContext creates a new context with authentication token in metadata
func createAuthContext(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{"token": "secret_token"})
	return metadata.NewOutgoingContext(ctx, md)
}

func main() {
	// устанавливаем соединение с сервером
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		ctx := createAuthContext(context.Background())
		resp, err := c.AddUser(ctx, &pb.AddUserRequest{
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
	ctx := createAuthContext(context.Background())
	resp, err := c.DelUser(ctx, &pb.DelUserRequest{
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
	baseCtx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	ctx = createAuthContext(baseCtx)

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
			// Only check resp.Error when there's no error
			if resp.Error == "" {
				fmt.Println(resp.User)
			} else {
				fmt.Println(resp.Error)
			}
		}
	}

	// получаем список email пользователей
	ctx = createAuthContext(context.Background())
	emails, err := c.ListUsers(ctx, &pb.ListUsersRequest{
		Offset: 0,
		Limit:  100,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emails.Count, emails.Emails)
}
