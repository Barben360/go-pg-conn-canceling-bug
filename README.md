# go-pg-conn-canceling-bug

Reproduction of bug mentioned in [go-pg issue 1879](https://github.com/go-pg/pg/issues/1879)

### Reproduce
- Run Postgres server:

```
docker-compose -f docker/docker-compose.yml up
```

- Run tests

```
go mod vendor
go test -v ./tests
```

### Expected

no error

### What you get

Example with the following configuration:

```go
const poolSize = 10
const concurrency = 10
const iterations = 100
const cancelProbability = 0.1
```


```
=== RUN   TestConcurrencyAndUserCancel
    bug_test.go:45: Canceling add statement in 4.374873ms
    bug_test.go:68: An error occurred for concurrent 9 - iteration 1 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.498459ms
    bug_test.go:68: An error occurred for concurrent 4 - iteration 3 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.654716ms
    bug_test.go:45: Canceling add statement in 4.633652ms
    bug_test.go:68: An error occurred for concurrent 3 - iteration 3 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 5 - iteration 3 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.857917ms
    bug_test.go:68: An error occurred for concurrent 0 - iteration 4 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 941.923µs
    bug_test.go:68: An error occurred for concurrent 8 - iteration 6 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 758.113µs
    bug_test.go:45: Canceling add statement in 1.256123ms
    bug_test.go:45: Canceling add statement in 3.750399ms
    bug_test.go:68: An error occurred for concurrent 3 - iteration 10 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 1.384943ms
    bug_test.go:68: An error occurred for concurrent 1 - iteration 11 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 1.430335ms
    bug_test.go:45: Canceling add statement in 3.577989ms
    bug_test.go:68: An error occurred for concurrent 5 - iteration 13 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.915759ms
    bug_test.go:68: An error occurred for concurrent 7 - iteration 16 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.75484ms
pg: 2021/05/06 09:33:45 tx.go:97: tx.Rollback failed: pg: Conn is in a bad state: write tcp [::1]:36276->[::1]:5454: i/o timeout
    bug_test.go:45: Canceling add statement in 1.950135ms
    bug_test.go:68: An error occurred for concurrent 9 - iteration 19 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.10189ms
    bug_test.go:45: Canceling add statement in 3.834528ms
    bug_test.go:68: An error occurred for concurrent 9 - iteration 24 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 5 - iteration 22 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.226879ms
    bug_test.go:68: An error occurred for concurrent 4 - iteration 24 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.919037ms
    bug_test.go:45: Canceling add statement in 463.389µs
pg: 2021/05/06 09:33:46 tx.go:97: tx.Rollback failed: pg: Conn is in a bad state: write tcp [::1]:36252->[::1]:5454: i/o timeout
    bug_test.go:45: Canceling add statement in 4.237009ms
    bug_test.go:68: An error occurred for concurrent 4 - iteration 26 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.366612ms
    bug_test.go:68: An error occurred for concurrent 1 - iteration 26 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 152.467µs
    bug_test.go:45: Canceling add statement in 2.993776ms
    bug_test.go:45: Canceling add statement in 2.344816ms
    bug_test.go:68: An error occurred for concurrent 4 - iteration 29 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.475881ms
    bug_test.go:68: An error occurred for concurrent 4 - iteration 30 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 621.648µs
    bug_test.go:45: Canceling add statement in 3.59575ms
    bug_test.go:68: An error occurred for concurrent 7 - iteration 31 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.460398ms
    bug_test.go:68: An error occurred for concurrent 8 - iteration 30 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 1.448211ms
    bug_test.go:68: An error occurred for concurrent 0 - iteration 32 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.206812ms
    bug_test.go:68: An error occurred for concurrent 7 - iteration 34 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.474471ms
    bug_test.go:68: An error occurred for concurrent 4 - iteration 35 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 1.455489ms
    bug_test.go:45: Canceling add statement in 1.893107ms
    bug_test.go:45: Canceling add statement in 4.472763ms
    bug_test.go:45: Canceling add statement in 4.954324ms
    bug_test.go:68: An error occurred for concurrent 2 - iteration 36 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.704425ms
    bug_test.go:68: An error occurred for concurrent 1 - iteration 36 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 7 - iteration 39 - call 3: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 9 - iteration 38 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 200.946µs
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 4 - iteration 39 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 212.028µs
    bug_test.go:45: Canceling add statement in 2.925492ms
    bug_test.go:68: An error occurred for concurrent 7 - iteration 44 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.905103ms
    bug_test.go:68: An error occurred for concurrent 8 - iteration 41 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 1.685018ms
    bug_test.go:68: An error occurred for concurrent 9 - iteration 45 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 1.995055ms
    bug_test.go:68: An error occurred for concurrent 4 - iteration 48 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.866789ms
    bug_test.go:45: Canceling add statement in 1.072232ms
    bug_test.go:68: An error occurred for concurrent 1 - iteration 47 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 9 - iteration 51 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 313.577µs
    bug_test.go:45: Canceling add statement in 3.511333ms
    bug_test.go:68: An error occurred for concurrent 8 - iteration 56 - call 3: context canceled
    bug_test.go:74: context canceled
    bug_test.go:45: Canceling add statement in 1.959807ms
    bug_test.go:68: An error occurred for concurrent 1 - iteration 57 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.318335ms
    bug_test.go:68: An error occurred for concurrent 4 - iteration 64 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.092982ms
    bug_test.go:68: An error occurred for concurrent 7 - iteration 66 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.72719ms
    bug_test.go:45: Canceling add statement in 3.591001ms
    bug_test.go:45: Canceling add statement in 4.877745ms
    bug_test.go:68: An error occurred for concurrent 6 - iteration 64 - call 3: context canceled
    bug_test.go:74: context canceled
    bug_test.go:68: An error occurred for concurrent 8 - iteration 65 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 4 - iteration 68 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 437.933µs
    bug_test.go:45: Canceling add statement in 3.513105ms
    bug_test.go:45: Canceling add statement in 4.175592ms
    bug_test.go:68: An error occurred for concurrent 5 - iteration 64 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 2 - iteration 67 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.890837ms
    bug_test.go:68: An error occurred for concurrent 6 - iteration 70 - call 3: context canceled
    bug_test.go:74: context canceled
    bug_test.go:45: Canceling add statement in 4.741915ms
    bug_test.go:68: An error occurred for concurrent 2 - iteration 70 - call 3: context canceled
    bug_test.go:74: context canceled
    bug_test.go:45: Canceling add statement in 2.397857ms
    bug_test.go:68: An error occurred for concurrent 4 - iteration 75 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.310946ms
    bug_test.go:45: Canceling add statement in 1.487027ms
    bug_test.go:68: An error occurred for concurrent 1 - iteration 72 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 7 - iteration 76 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 587.979µs
pg: 2021/05/06 09:33:47 tx.go:97: tx.Rollback failed: pg: Conn is in a bad state: read tcp [::1]:36360->[::1]:5454: i/o timeout
    bug_test.go:45: Canceling add statement in 1.155616ms
    bug_test.go:45: Canceling add statement in 416.41µs
    bug_test.go:45: Canceling add statement in 1.121441ms
pg: 2021/05/06 09:33:47 tx.go:97: tx.Rollback failed: pg: Conn is in a bad state: read tcp [::1]:36256->[::1]:5454: i/o timeout
    bug_test.go:45: Canceling add statement in 1.187639ms
    bug_test.go:68: An error occurred for concurrent 2 - iteration 76 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.004148ms
    bug_test.go:68: An error occurred for concurrent 5 - iteration 72 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 1.244585ms
pg: 2021/05/06 09:33:47 tx.go:97: tx.Rollback failed: pg: Conn is in a bad state: read tcp [::1]:36288->[::1]:5454: i/o timeout
    bug_test.go:45: Canceling add statement in 3.067203ms
    bug_test.go:68: An error occurred for concurrent 7 - iteration 85 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.143307ms
    bug_test.go:45: Canceling add statement in 102.128µs
    bug_test.go:45: Canceling add statement in 4.427769ms
    bug_test.go:68: An error occurred for concurrent 9 - iteration 87 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.1µs
    bug_test.go:45: Canceling add statement in 1.27478ms
    bug_test.go:45: Canceling add statement in 3.970283ms
    bug_test.go:68: An error occurred for concurrent 7 - iteration 89 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.227124ms
    bug_test.go:45: Canceling add statement in 2.771362ms
    bug_test.go:68: An error occurred for concurrent 5 - iteration 83 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 0 - iteration 85 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.803931ms
    bug_test.go:45: Canceling add statement in 2.775106ms
    bug_test.go:68: An error occurred for concurrent 7 - iteration 94 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 0 - iteration 87 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.12815ms
    bug_test.go:45: Canceling add statement in 4.557329ms
    bug_test.go:45: Canceling add statement in 547.343µs
    bug_test.go:68: An error occurred for concurrent 7 - iteration 97 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 1.091205ms
    bug_test.go:68: An error occurred for concurrent 7 - iteration 98 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:68: An error occurred for concurrent 3 - iteration 91 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.782234ms
    bug_test.go:68: An error occurred for concurrent 3 - iteration 94 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 2.956099ms
    bug_test.go:68: An error occurred for concurrent 1 - iteration 95 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.112067ms
    bug_test.go:68: An error occurred for concurrent 1 - iteration 96 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 3.604201ms
pg: 2021/05/06 09:33:47 tx.go:97: tx.Rollback failed: pg: Conn is in a bad state: read tcp [::1]:36386->[::1]:5454: i/o timeout
    bug_test.go:45: Canceling add statement in 4.047676ms
    bug_test.go:68: An error occurred for concurrent 3 - iteration 99 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
    bug_test.go:45: Canceling add statement in 4.144495ms
    bug_test.go:68: An error occurred for concurrent 5 - iteration 97 - call 3: context canceled
    bug_test.go:74: context canceled
    bug_test.go:45: Canceling add statement in 3.919807ms
    bug_test.go:68: An error occurred for concurrent 5 - iteration 98 - call 3: context deadline exceeded
    bug_test.go:74: context deadline exceeded
--- FAIL: TestConcurrencyAndUserCancel (2.35s)
FAIL
FAIL    github.com/Barben360/go-pg-conn-canceling-bug/tests     2.360s
FAIL
```
