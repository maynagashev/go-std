package persistent

import (
	"errors"
	"project/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockStore(ctrl)

	// возвращаемая ошибка
	errEmptyKey := errors.New("Указан пустой ключ")

	m.EXPECT().Get("").Return([]byte(""), errEmptyKey)
	_, err := Lookup(m, "")
	require.ErrorIs(t, err, errEmptyKey)
}
