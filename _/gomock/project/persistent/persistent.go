// файл persistent/persistent.go
package persistent

import (
	"project/store"
)

func Lookup(s store.Store, key string) ([]byte, error) {
	// ...
	return s.Get(key)
}
