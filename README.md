# go-simple-chat [![GoDoc](https://godoc.org/github.com/utrack/go-simple-chat?status.svg)](https://godoc.org/github.com/utrack/go-simple-chat)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/utrack/go-simple-chat/blob/master/LICENSE)
[![Go Report Card](http://goreportcard.com/badge/utrack/go-simple-chat)](http://goreportcard.com/report/utrack/go-simple-chat)
[![Build Status](https://travis-ci.org/utrack/go-simple-chat.svg)](https://travis-ci.org/utrack/go-simple-chat)
[![codecov.io](https://codecov.io/github/utrack/go-simple-chat/coverage.svg?branch=master)](https://codecov.io/github/utrack/go-simple-chat?branch=master)

Simple single-channel chat app written in Go.

# Requirements
Golang compiler and tools (v1.5 or later) are required. See the [official Getting Started guide](https://golang.org/doc/install) or your distro's docs for detailed instructions.

# Installation
```
go get -u github.com/utrack/go-simple-chat/cmd/gosimplechat
```
If you're using Go < 1.6 - you need to set envvar `GO15VENDOREXPERIMENT` to `1` before go-getting:
```
GO15VENDOREXPERIMENT=1 go get -u github.com/utrack/go-simple-chat/cmd/gosimplechat
```

# Running
Check that your `PATH` envvar has `$GOPATH\bin` and run the command:
```
gosimplechat
```

Open your browser, navigate to `localhost:8080` and chat away!

Use flag `-log debug` for more verbose logging.

# Testing
```
go test github.com/utrack/go-simple-chat/...
```
Tests are written using the [GoConvey](https://github.com/smartystreets/goconvey) framework. If you have `goconvey` tools installed in your `$PATH`, cd to the project's path and run `goconvey` to use its web interface.
