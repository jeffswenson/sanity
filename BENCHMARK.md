# Benchmarks

Sanity's benchmarks were designed to compare the performance of Sanity to the
performance of `text/html`. It is difficult fairly compare Sanity to
`text/html` because their performance penalties are asymteric.

`text/html` uses reflection when interacting with the template's model.
Reflection in Go is relatively slow, so the cost of most templates is driven by
the number `{{ }}` blocks in the template. If a template contains no dynamic
blocks, it compiles down to a single `io.Writer.Write` call and will outperform
Sanity.

The internal design of Sanity was driven by a desire to minimize the number of
allocations required to build a node tree, but even so every Sanity tag
requires exactly one allocation and benchmarking Sanity is effectively
benchmarking the Go allocator and garbage collector. So the cost of rendering a
Sanity view directly tracks the number of tags rendered. But one allocation is
much faster than one `{{ }}` block in a template, so most Sanity views will out
perform an equivalent `text/html` template.

## Micro Benchmarks

`internal/benchmark` contains micro benchmarks. Rendering a list with 1k
elements requires 145us using Sanity and 803us using `text/html`.

```
goos: linux
goarch: amd64
pkg: github.com/jeffswenson/sanity/internal/benchmark
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkSanityList1000-8           8078            145831 ns/op
PASS
ok      github.com/jeffswenson/sanity/internal/benchmark 1.201s
```

```
goos: linux
goarch: amd64
pkg: github.com/jeffswenson/sanity/internal/benchmark
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkGoTemplate1000-8           1501            803934 ns/op
PASS
ok      github.com/jeffswenson/sanity/internal/benchmark 1.295s
```

## HTTP Benchmarks

`internal/benchmarkhttp` contains benchmarks that use the stdlib http server to
render HTML content. The HTML content is similar to a popular orange tech news
site. The benchmark contains tests with and without gzip.

The tests were run with `GOMAXPROCS=1` and `GOMEMLIMIT=100Mib`. GOMAXPROCS was
set to measure the single threaded throughput. The memory limit was set to to
force frequent garbage collection. The tests saturated the server's capacity,
so the latency numbers are meaningless.

A key takeway from the test is gziping the HTML costs far more cpu cyles than
producing the HTML, so rendering the view is unlikely to be a bottleneck in
most applications.

```
Bombarding http://localhost:8080/sanity for 10s using 125 connection(s)
Statistics        Avg      Stdev        Max
  Reqs/sec      8693.36     851.13    9551.49
  Latency       14.42ms    31.99ms   381.55ms
  HTTP codes:
    1xx - 0, 2xx - 86619, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    92.30MB/s

Bombarding http://localhost:8080/sanity-gzip for 10s using 125 connection(s)
Statistics        Avg      Stdev        Max
  Reqs/sec      2698.47     462.90   11584.86
  Latency       46.44ms   187.45ms      2.05s
  HTTP codes:
    1xx - 0, 2xx - 26943, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     2.40MB/s

Bombarding http://localhost:8080/template for 10s using 125 connection(s)
Statistics        Avg      Stdev        Max
  Reqs/sec      2230.99     353.45    8751.29
  Latency       56.25ms    23.46ms   129.58ms
  HTTP codes:
    1xx - 0, 2xx - 22228, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    27.81MB/s

Bombarding http://localhost:8080/template-gzip for 10s using 125 connection(s)
Statistics        Avg      Stdev        Max
  Reqs/sec      1203.30      90.52    1350.24
  Latency      103.27ms   191.40ms      1.32s
  HTTP codes:
    1xx - 0, 2xx - 12157, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     1.50MB/s
```
