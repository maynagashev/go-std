package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type (
	AccountBalance struct {
		AccountIdHash []byte           `yaml:"account_id_hash,flow"`
		Amounts       []CurrencyAmount `yaml:"amounts,omitempty"`
		IsBlocked     bool             `yaml:"is_blocked"`
	}

	CurrencyAmount struct {
		Amount   int64  `yaml:"amount"`
		Decimals int8   `yaml:"decimals"`
		Symbol   string `yaml:"symbol"`
	}
)

func main() {
	balance := AccountBalance{
		AccountIdHash: []byte{0x10, 0x20, 0x0A, 0x0B}, // 8 bit - 1 byte - 256 values - 0x00..0xFF - 2^8
		Amounts: []CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"},
			{Amount: 2510, Decimals: 2, Symbol: "USD"},
		},
		IsBlocked: true,
	}

	// преобразуем значение переменной balance в YAML-формат
	out, err := yaml.Marshal(balance)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
