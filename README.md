# zzcsi [![builds.sr.ht status](https://builds.sr.ht/~mna/zzcsi.svg)](https://builds.sr.ht/~mna/zzcsi?) [![GoDoc](https://godoc.org/git.sr.ht/~mna/zzcsi?status.svg)](http://godoc.org/git.sr.ht/~mna/zzcsi) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/git.sr.ht/~mna/zzcsi)

* Canonical repository: https://git.sr.ht/~mna/zzcsi
* Issues: https://todo.sr.ht/~mna/zzcsi
* Builds: https://builds.sr.ht/~mna/zzcsi

Package zzcsi implements the Control Sequence Introducer (CSI) functions
supported by Xterm-compatible terminals. The most commonly used functions are
to erase a line or the screen, manipulate the cursor or set the display
attributes such as foreground and background colors, bold or italics, but many
more functions are available. See the [xterm documentation][xterm] for more
information on CSI sequences.

See the [package documentation][godoc] for details, API reference and usage
example (alternatively, on [pkg.go.dev][pgd]).

This is a package in the [zzterm][zzt] family of Go packages, which aim to
help build terminal-based programs with a focus on speed and memory efficiency.

## Benchmarks

When using a buffer of sufficient capacity, `CSI.AppendFunc` can append the
CSI escape sequence without allocation.

```
benchmark                       iter      time/iter   bytes alloc        allocs
---------                       ----      ---------   -----------        ------
BenchmarkCSI/Func-4          9259478   118.00 ns/op       24 B/op   2 allocs/op
BenchmarkCSI/FuncString-4    9003938   133.00 ns/op       32 B/op   3 allocs/op
BenchmarkCSI/AppendFunc-4   20808049    58.40 ns/op        0 B/op   0 allocs/op
```

## License

The [BSD 3-Clause license][bsd].

[bsd]: http://opensource.org/licenses/BSD-3-Clause
[godoc]: http://godoc.org/git.sr.ht/~mna/zzcsi
[pgd]: https://pkg.go.dev/git.sr.ht/~mna/zzcsi
[xterm]: https://invisible-island.net/xterm/ctlseqs/ctlseqs.html#h3-Functions-using-CSI-_-ordered-by-the-final-character_s_
[zzt]: https://git.sr.ht/~mna/?search=zz
