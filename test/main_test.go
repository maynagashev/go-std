package main

import (
	"errors"
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
			got := Abs(tt.value)
			if got != tt.want {
				t.Errorf("Abs(%v) = %v; want %v", tt.value, got, tt.want)
			}
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
			if got := tt.u.FullName(); got != tt.want {
				t.Errorf("User.FullName() = %v; want %v", got, tt.want)
			}
		})
	}
}

func TestFamily_AddNew(t *testing.T) {
	f := Family{}
	err := f.AddNew(Father, Person{
		FirstName: "Misha",
		LastName:  "Popov",
		Age:       56,
	})
	if err != nil {
		t.Errorf("Family.AddNew() error = %v; want nil", err)
	}
	if _, ok := f.Members[Father]; !ok {
		t.Errorf("Family.AddNew() error = %v; want nil", errors.New("father not added"))
	}
	err = f.AddNew(Father, Person{
		FirstName: "Drug",
		LastName:  "Mishi",
		Age:       57,
	})
	if err == nil {
		t.Errorf("Family.AddNew() error = %v; want not nil", err)
	}
}
