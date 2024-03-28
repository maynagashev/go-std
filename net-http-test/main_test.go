package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Напишите тесты для хендлера. Найдите ошибки, не запуская приложение.
// Ваши тесты должны проверить:
// возвращаемый JSON;
// код ответа (200, 400 ,404, 500);
// заголовок Content-Type (должен быть application/json).
func TestUserViewHandler(t *testing.T) {

	type want struct {
		code        int
		response    string
		contentType string
	}
	usersWithOneUser := map[string]User{
		"u1": {
			ID:        "u1",
			FirstName: "Misha",
			LastName:  "Popov",
		},
	}

	tests := []struct {
		name   string
		target string
		users  map[string]User
		want   want
	}{
		{
			name:   "success",
			target: "/users?user_id=u1",
			users:  usersWithOneUser,
			want: want{
				code:        http.StatusOK,
				response:    `{"ID":"u1","FirstName":"Misha","LastName":"Popov"}`,
				contentType: "application/json",
			},
		},
		{
			name:   "user not found",
			target: "/users?user_id=u2",
			users:  usersWithOneUser,
			want: want{
				code:        http.StatusNotFound,
				response:    `{}`,
				contentType: "application/json",
			},
		},
		{
			name:   "empty user_id",
			target: "/users?user_id=",
			users:  usersWithOneUser,
			want: want{
				code:        http.StatusBadRequest,
				response:    `{"error":"user_id is empty"}`,
				contentType: "application/json",
			},
		},
		{
			name:   "internal error",
			target: "/users?user_id=u1",
			users:  nil,
			want: want{
				code:        http.StatusInternalServerError,
				response:    `{}`,
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, test.target, nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			UserViewHandler(test.users)(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.JSONEq(t, test.want.response, string(resBody))
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
