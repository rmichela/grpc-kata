# gRPC in Go, C#, and Java: A microservices code Kata

In this repo, we have the same Hello World service written in three langaues, Java, Go, and C#. This repo demonstrates
how to use the gRPC framework from each language, and how gRPC facilitates cross-language interoperability.

## What is gRPC

"gRPC is a modern open source high performance RPC framework" - [grpc.io](https://grpc.io/)

* Polyglot - [10+ languages](https://grpc.io/docs/)
* Binary RPC - [protobuf](https://github.com/protocolbuffers/protobuf)
* Standards-based - [http/2](https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md)

![gRPC diagram](https://grpc.io/img/landing-2.svg)

## Why gRPC?

* Multiplexed RPC
* Bi-directional streaming
* Strong [community ecosystem](https://github.com/grpc-ecosystem/awesome-grpc)
* Strong security - mTLS 🔐
* First-class observability 👀 - [metrics and tracing](https://opencensus.io/guides/grpc/) built in
* Well defined separation of concerns - interceptors for middleware

## Where does gRPC fit?

* 🔥Works great behind the firewall
* 📱Works great with mobile clients
* 💻Works in [the browser](https://github.com/grpc/grpc-web)

## Basic gRPC workflow

1. Add library dependencies
1. Write proto service contract
1. Generate client and server stubs
1. Implement server stub and start server
1. Call service with client stub

## Building

This repo is divided into three subdirectories, one for each language. There are also a series of git tags you
can progress through to see each gRPC example progress one at a time. See the readme in each subdirectory
for language specific details.
