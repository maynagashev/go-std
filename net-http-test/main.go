package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User — основной объект для теста.
type User struct {
	ID        string
	FirstName string
	LastName  string
}

// UserViewHandler — хендлер, который нужно протестировать.
// Он должен возвращать информацию о пользователе в формате JSON.
func UserViewHandler(users map[string]User) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		// fix
		rw.Header().Set("Content-Type", "application/json")
		if userId == "" {
			//http.Error(rw, "user_id is empty", http.StatusBadRequest)

			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("{\"error\":\"user_id is empty\"}"))

			return
		}

		user, ok := users[userId]
		if !ok {
			http.Error(rw, "user not found", http.StatusNotFound)
			return
		}

		jsonUser, err := json.Marshal(user)
		if err != nil {
			http.Error(rw, "can't provide a json. internal error",
				http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonUser)
	}
}

func main() {
	users := make(map[string]User)
	u1 := User{
		ID:        "u1",
		FirstName: "Misha",
		LastName:  "Popov",
	}
	u2 := User{
		ID:        "u2",
		FirstName: "Sasha",
		LastName:  "Popov",
	}
	users["u1"] = u1
	users["u2"] = u2

	http.HandleFunc("/users", UserViewHandler(users))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
