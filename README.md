# go-rob

[![GoDev][godev-image]][godev-url]

go-rob パッケージは []byte, string 間の低コストな相互キャストを提供します.

## Usage

```go
var b []byte = rob.Bytes("string")
var s string = rob.String(b)
```

## License

This software is released under the MIT License, see LICENSE.

## Author

17e10

[godev-image]: https://pkg.go.dev/badge/github.com/17e10/go-rob
[godev-url]: https://pkg.go.dev/github.com/17e10/go-rob
