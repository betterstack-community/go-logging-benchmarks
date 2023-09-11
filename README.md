# Benchmarks for Go Logging Libraries

This repo compares the performance of several
[Go Logging Libraries](https://betterstack.com/community/guides/logging/best-golang-logging-libraries/)
. It always uses the latest version of each library, and the results are
automatically updated once a week.

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
