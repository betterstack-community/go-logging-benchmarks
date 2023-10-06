# Benchmarks for Go Logging Libraries

This repo compares the performance of several
[Go Logging Libraries](https://betterstack.com/community/guides/logging/best-golang-logging-libraries/)
. The results are automatically updated once a day.

The following libraries are being tested:

- [Zerolog](https://github.com/rs/zerolog).
- [Zap](https://github.com/uber-go/zap).
- [Apex/log](https://github.com/uber-go/zap).
- [Logrus](https://github.com/sirupsen/logrus).
- [Slog](https://pkg.go.dev/log/slog).
- [Phuslog](https://github.com/phuslu/log).
- [Log15](https://github.com/inconshreveable/log15).
- [Logf](https://github.com/zerodha/logf).

<!-- ![Screenshot or GIF of the application in action](screenshot.png). -->

## 🟢 Prerequisites

You only need [a recent version of Go](https://go.dev/doc/install) to execute
the benchmarks.

## 📦 Getting started

- Clone the repository to your computer:

```bash
git clone https://github.com/betterstack-community/go-logging-benchmarks && cd go-logging-benchmarks
```

- Download the dependencies:

```bash
go mod tidy
```

- Execute the benchmarks:

```bash
go test -bench=. -benchmem
```

## ⚖ License

The code used in this project and in the linked tutorial are licensed under the
[Apache License, Version 2.0](LICENSE).

## 🤝 Contribute

If you found an issue with the benchmarks, or you want to propose a new library
for benchmarking, please open an issue or pull request accordingly.
