# azamcodec-go

[![Build Status](https://github.com/azam/azamcodec-go/actions/workflows/build.yml/badge.svg?version=latest)](https://github.com/azam/azamcodec-go/actions/workflows/build.yml)

An encoder and decoder implementation in Go for [Azam Codec](https://github.com/azam/azamcodec), a lexicographically sortable multi-section base16 encoding of byte array. Zero external dependencies.

## License

MIT Licence
Copyright (c) 2024 Azamshul Azizy

## Usage

Import the crate and start using it.

### Decoding

```go
import {
    "github.com/azam/azamcodec"
}

// Decode Azam-encoded string as slice of `uint`'s` (`[]uint`).
nums, err := azamcodec.Decode("xytxvyyfh5wgg1"); // nums = []uint {0xdeadbeef, 0x15, 0xc001}
```

### Encoding

```go
import {
    "github.com/azam/azamcodec"
}

// Encode uint value as Azam-encoded string.
// 0xdeadbeefu32 encodes to "xytxvyyf".
str = azamcodec.EncodeUint(0xdeadbeef) // "xytxvyyf"

// Encode multiple `uint` values as Azam-encoded string.
// 0xdeadbeefu32 encodes to "xytxvyyf".
// 0x15u8 encodes to "h5".
// 0xc001u16 encodes to "wgg1".
// Resulting encoded string appends in order.
str = azamcodec.EncodeUints(0xdeadbeef, 0x15, 0xc001) // "xytxvyyfh5wgg1"
```

## Development

Standard Go development applies. Benchmark is also included, executable via `go test -bench`.
