# String Concatenation

There are some methods in golang for string concatenation:

```go
a, b := "aaa", "bbb"

// operator "+":
v1 := a + b

// fmt.Sprintf
v2 := fmt.Sprintf("%s%s", a, b)

// bytes.Buffer
buf := new(bytes.Buffer)
buf.WriteString(a)
buf.WriteString(b)
v3 := buf.String()

// strings.Builder
var builder strings.Builder
builder.WriteString(a)
builder.WriteString(b)
v4 := builder.String()

// other methods...
```

We can know the most efficient method above is `strings.Builder` by practical benchmarks or some tech articles. But I
don't think the writing form of `strings.Builder` is concise. Hope that I am not the only one :)

```go
// if the number of elements you want to concatenate is quite large, you have to write a lot of code lines.
var builder strings.Builder
builder.WriteString(a)
builder.WriteString(b)
builder.Write(c)
builder.Write(d)
builder.WriteRune(e)
```

So I find some different usages of `strings.Builder` from many most-starred repositories in github.

```go
// https://github.com/golang/go/blob/master/src/go/types/methodset.go
func (s *MethodSet) String() string {
    if s.Len() == 0 {
        return "MethodSet {}"
    }

    var buf strings.Builder
    fmt.Fprintln(&buf, "MethodSet {")
    for _, f := range s.list {
        fmt.Fprintf(&buf, "\t%s\n", f)
    }
    fmt.Fprintln(&buf, "}")
    return buf.String()
}
```

```go
// https://github.com/etcd-io/etcd/blob/main/raft/util.go
func DescribeHardState(hs pb.HardState) string {
    var buf strings.Builder
    fmt.Fprintf(&buf, "Term:%d", hs.Term)
    if hs.Vote != 0 {
        fmt.Fprintf(&buf, " Vote:%d", hs.Vote)
    }
    fmt.Fprintf(&buf, " Commit:%d", hs.Commit)
    return buf.String()
}
```

## Benchmark
Benchmark results from [string_concatenation_test.go](string_concatenation_test.go):

```shell
# go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/sasakiyori/efficient-go/string_concatenation
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkStringBuilderInCommonUsage-16          18878925                61.73 ns/op           48 B/op          2 allocs/op
BenchmarkStringBuilderInFmtf-16                  4856551               260.1 ns/op           104 B/op          4 allocs/op
BenchmarkStringBuilderInFmt-16                   2360289               546.5 ns/op           152 B/op          4 allocs/op
BenchmarkBytesBuffer-16                         15711445                70.99 ns/op           96 B/op          2 allocs/op
```

Regard `BenchmarkStringBuilderInCommonUsage` as measuring unit 1, then:
- `BenchmarkStringBuilderInFmtf` ≈ [4,5]
- `BenchmarkStringBuilderInFmt` ≈ [8,9]
- `BenchmarkBytesBuffer` ≈ [1,2]

Use `fmt.Fprintf` in `strings.Builder` is better than `fmt.Fprint` because it specifics the tpyes of arguments. But their performaces are worse, not recommended.
Thus, the original form to use `strings.Builder` is the best practice.