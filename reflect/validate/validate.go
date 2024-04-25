package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	u := User{"admin", 20, 88}
	Validate(u)
}

// User используется для тестирования.
type User struct {
	Nick string
	Age  int `limit:"18"`
	Rate int `limit:"0,100"`
}

// Str2Int конвертирует строку в int.
func Str2Int(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Validate проверяет min и max для int c тегом limit.
func Validate(obj interface{}) bool {
	vobj := reflect.ValueOf(obj)
	objType := vobj.Type() // получаем описание типа

	// перебираем все поля структуры
	for i := 0; i < objType.NumField(); i++ {
		// берём значение текущего поля и проверяем, что это int
		if v, ok := vobj.Field(i).Interface().(int); ok {
			log.Printf("valueObj: %#v, objType: %#v\nfield:%#v, intValue: %#v\n", vobj, objType, vobj.Field(i), v)

			// подсказка: тег limit надо искать в поле objType.Field(i)
			// objType.Field(i).Tag.Lookup или objType.Field(i).Tag.Get

			// Информация о конкретном поле структуры
			field := objType.Field(i)
			tag := field.Tag
			lookup, _ := tag.Lookup("limit")
			log.Printf("field: %#v, tag: %#v lookup: %#v\n", field, tag, lookup)

			limits := strings.Split(lookup, ",")
			log.Printf("limits: %#v\n", limits)

			// Проверка минимума
			if len(limits) > 0 {
				minLimit := Str2Int(limits[0])
				if v < minLimit {
					return false
				}
			}

			// Проверка максимума
			if len(limits) > 1 {
				maxLimit := Str2Int(limits[1])
				if v > maxLimit {
					return false
				}
			}
		}
	}
	return true
}
