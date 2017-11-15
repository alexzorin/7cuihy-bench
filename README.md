# Benchmark test for [this thread](https://www.reddit.com/r/webdev/comments/7cuihy/advice_speed_of_raw_calculations_node_golang/)

The benchmark sorts a 1e6 (1 million) array into (less than 500, greater than 500).

Generation of the data is not benchmarked.

I optimised the original Node code slightly to avoid some dumb speed pitfalls.

Node v6.11.5 vs Go 1.9.

## Raw Results

### Go
For Go, there are two tests, one using float64 (because there's no way to avoid using Float64 in Node), and one that uses uint64 (because it is an optimisation available to Go and it is much faster than using floats).

We can see that it can sort 1mil integers in 1ms and 1mil floats in 4.9ms.
```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/alexzorin/7cuihy-bench
Benchmark1milSort-4                          300           4963129 ns/op
Benchmark1milSortOnlyIntegers-4             1000           1169818 ns/op
PASS
ok      github.com/alexzorin/7cuihy-bench       3.444s
```

### Node
The result is that it can sort 1mil Numbers (which are float64s) in around ~5.4ms.
```
$ node node.js
Sort the data x 183 ops/sec ±0.37% (82 runs sampled)
```

### Interpretation
Unsurprisingly, the results are extremely similar for float64 because it is very trivial code for both Go and Node to optimise.

Using only integers speeds it up by a significant factor, but is unlikely to be a bottleneck ...
