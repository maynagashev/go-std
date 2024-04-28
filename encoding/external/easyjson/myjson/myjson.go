// myeasyjson/myjson/myjson.go
// не забудьте запустить easyjson -all myjson.go
package myjson

//go:generate easyjson -all myjson.go
type (
	AccountBalance struct {
		AccountIdHash []byte           `json:"account_id_hash,flow"`
		Amounts       []CurrencyAmount `json:"amounts,omitempty"`
		IsBlocked     bool             `json:"is_blocked"`
	}

	CurrencyAmount struct {
		Amount   int64  `json:"amount"`
		Decimals int8   `json:"decimals"`
		Symbol   string `json:"symbol"`
	}
)
