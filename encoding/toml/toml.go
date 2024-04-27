package main

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
	// 2) преобразуйте полученную переменную в TOML
	// 3) выведите в консоль результат
	// ...
}
