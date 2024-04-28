package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	balance := AccountBalance{
		AccountIdHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []*CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"},
			{Amount: 2510, Decimals: 2, Symbol: "USD"},
		},
		IsBlocked: true,
	}

	// сериализуем значение переменной balance
	protoBz, err := proto.Marshal(&balance)
	if err != nil {
		panic(err)
	}

	var balanceCopy AccountBalance
	// декодируем данные в новую переменную
	if err := proto.Unmarshal(protoBz, &balanceCopy); err != nil {
		panic(err)
	}

	// визуально сравниваем значения переменных
	fmt.Printf("balanceInit: %s\n", balance.String())
	fmt.Printf("protoBz: %v\n", protoBz)
	fmt.Printf("balanceCopy: %s\n", balanceCopy.String())
}
