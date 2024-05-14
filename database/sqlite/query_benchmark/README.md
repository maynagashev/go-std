```bash
go test -bench=. query_test.go
```

Без индекса 40к записей
```bash
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkQuery-12            306           5333791 ns/op
PASS
ok      command-line-arguments  2.418s
```

С индексом 40к записей
```bash
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkQuery-12           1000           1204595 ns/op
PASS
ok      command-line-arguments  1.552s
```
