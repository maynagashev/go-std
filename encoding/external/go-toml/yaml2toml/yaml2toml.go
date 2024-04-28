package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
}

const yamlData = `
id: 101
name: Gopher
values:
- 11
- 22
- 33
`

func main() {
	// вставьте недостающий код
	// 1) десериализуйте yamlData в переменную типа Data
	var d Data
	err := yaml.Unmarshal([]byte(yamlData), &d)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%#v\n", d)

	// 2) преобразуйте полученную переменную в TOML
	encoded, err := toml.Marshal(d)

	// 3) выведите в консоль результат
	fmt.Println(string(encoded))
}
