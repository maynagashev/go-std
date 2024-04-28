package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

type (
	AccountBalance struct {
		AccountIdHash []byte           `toml:"account_id_hash"`
		Amounts       []CurrencyAmount `toml:"amounts,omitempty"`
		IsBlocked     bool             `toml:"is_blocked" comment:"Deprecated" commented:"true"`
	}

	CurrencyAmount struct {
		Amount   int64  `toml:"amount"`
		Decimals int8   `toml:"decimals"`
		Symbol   string `toml:"symbol"`
	}
)

func main() {
	balance := AccountBalance{
		AccountIdHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"},
			{Amount: 2510, Decimals: 2, Symbol: "USD"},
		},
		IsBlocked: true,
	}

	// преобразуем значение переменной balance в TOML-формат
	out, err := toml.Marshal(balance)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
