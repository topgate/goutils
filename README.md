# goutils

Go 言語ユーティリティライブラリ集

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/topgate/goutils)
[![CircleCI](https://circleci.com/gh/topgate/goutils.svg?style=shield)](https://circleci.com/gh/topgate/goutils)

## Installing / Getting started

A quick introduction of the minimal setup you need to get a hello world up &
running.

```shell
go get github.com/topgate/goutils
```

または、go mod を使用している場合、ソースコード中に `github.com/topgate/goutils` のパッケージを参照している状態で、

```shell
go mod tidy
```

を実行する

## Developing

```shell
git clone https://github.com/topgate/goutils.git
cd goutils
go mod tidy
```

### Testing

テスト実行手順を以下に示す

```shell
go test ./...
```

## Features

このライブラリで提供する機能

- [Excel に対応した CSV の出力](https://pkg.go.dev/github.com/topgate/goutils/interop/excel?tab=subdirectories) (encoding/csv,gocsv に対応したパッケージ)
- [AppEngine 用のロギング](https://pkg.go.dev/github.com/topgate/goutils/gcp/appengine/log)
- [構造体のバリデーション](https://pkg.go.dev/github.com/topgate/goutils/validate)
- [SQL 拡張](https://pkg.go.dev/github.com/topgate/goutils/database/sqlx)

## Licensing

https://github.com/topgate/goutils/blob/master/LICENSE
