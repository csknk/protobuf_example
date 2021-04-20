# Simple Demo of Protobuf in Go

Simple working implementation based on [this tutorial][2].

This is a much simpler method for serializing data structures than the approach taken by Bitcoin, which I've described [here][3] and [here][4].

Protobuf allows you to serialize and unserialize data such that only a minimal amount of information needs to be transferred. It works by defining a data structure in a `.proto` file, and using the protobuf tool to generate code that allows you to serialize and unserialize data appropriately.


Usage
-----
* Clone this repo
* Generate the go code using the protocol buffer compiler
* It it's not installed, run `go install google.golang.org/protobuf/cmd/protoc-gen-go`
* Within the project directory, run `protoc --go_out $GOPATH/pkg/mod addressbook.proto`
* This builds the required mod file in your `$GOPATH`
* `cd` into the project and run `go get`
* Play about with it!

References
----------
* [Protocol Buffers (a.k.a., protobuf) Github][1]
* [Protocol Buffer Basics: Go, Google developers][2]

[1]: https://github.com/protocolbuffers/protobuf
[2]: https://developers.google.com/protocol-buffers/docs/gotutorial
[3]: https://github.com/csknk/bitcoin-varint
[4]: https://github.com/csknk/parse-chainstate
