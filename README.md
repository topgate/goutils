# goutils

Go 言語ユーティリティライブラリ集

[![GoDoc](https://godoc.org/github.com/topgate/goutils?status.svg)](https://godoc.org/github.com/topgate/goutils)
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

- Excel に対応した CSV の出力 (encoding/csv,gocsv に対応したパッケージ)
- AppEngine 用のロギング

## Licensing

https://github.com/topgate/goutils/blob/master/LICENSE
