# errgroup
优化官方 errgroup 包遇到错误仍继续执行的问题


```
$go test -bench=. -benchtime=10s -benchmem     
goos: darwin
goarch: amd64
pkg: github.com/why444216978/errgroup
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkMyErrgroup-8                 10        1003383932 ns/op             400 B/op          5 allocs/op
BenchmarkGolangErrgroup-8              2        6004337452 ns/op             352 B/op          7 allocs/op
PASS
ok      github.com/why444216978/errgroup        29.065s
```
