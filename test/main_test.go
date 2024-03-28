package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAbs(t *testing.T) {

	// стоит проверить на значениях -3, 3, -2.000001, -0.000000003 и так далее

	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{name: "positive", value: 3.14, want: 3.14},
		{name: "negative", value: -3.14, want: 3.14},
		{name: "zero", value: 0, want: 0},
		{name: "negative zero", value: -0, want: 0},
		{name: "negative 3", value: -3, want: 3},
		{name: "positive 3", value: 3, want: 3},
		{name: "negative 2.000001", value: -2.000001, want: 2.000001},
		{name: "negative 0.000000003", value: -0.000000003, want: 0.000000003},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// меняем на функцию Equal из пакета assert
			assert.Equal(t, Abs(tt.value), tt.want)
		})
	}
}

func TestUser_FullName(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want string
	}{
		{name: "empty", u: User{}, want: " "},
		{name: "only first name", u: User{FirstName: "Misha"}, want: "Misha "},
		{name: "only last name", u: User{LastName: "Popov"}, want: " Popov"},
		{name: "full name", u: User{FirstName: "Misha", LastName: "Popov"}, want: "Misha Popov"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.u.FullName(), tt.want)

		})
	}
}

func TestFamily_AddNew(t *testing.T) {
	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newPerson      newPerson
		wantErr        bool
	}{
		{
			name: "add father",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Maria",
					LastName:  "Popova",
					Age:       36,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			wantErr: false,
		},
		{
			name: "catch error",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Ken",
					LastName:  "Gymsohn",
					Age:       32,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Family{
				Members: tt.existedMembers,
			}
			err := f.AddNew(tt.newPerson.r, tt.newPerson.p)
			if !tt.wantErr {
				// обязательно проверяем на ошибки
				require.NoError(t, err)
				// дополнительно проверяем, что новый человек был добавлен
				assert.Contains(t, f.Members, tt.newPerson.r)
				return
			}

			assert.Error(t, err)
		})
	}
}
