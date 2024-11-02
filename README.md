# jsony

A blazing fast and safe Go package for serializing JSON.

Features:

* 2-3 times faster than stdlib
* type safe and with no runtime errors or panics
* pure go
* reflection-free
* objects preserve elements' order
* objects can be constructed dynamically

## Installation

```bash
go get github.com/orsinium-labs/jsony
```

## Usage

```go
obj := jsony.Object{
   jsony.Field{"name", jsony.String("Aragorn")},
   jsony.Field{"age", jsony.Int(87)},
}
s := jsony.EncodeString(obj)
fmt.Println(s)
```

## Benchmarks

Each value is time (in ns) per operation, as reported by the Go built-in benchmark framework. Lower is better.

| category     | jsony     | stdlib   |
| ------------ | --------: | -------: |
| Int          |     15 🏆 |       54 |
| Float64      |     63 🏆 |      122 |
| String       |     22 🏆 |       88 |
| Object       |    134 🏆 |      136 |
| Map          |    246 🏆 |      662 |
| MixedArray   |    125 🏆 |      271 |
| IntArray     |     88 🏆 |      183 |
| BigArray     |  64262 🏆 |   115994 |

To reproduce, run `task bench`
