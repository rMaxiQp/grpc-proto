# grpc-proto
GRPC Proto playground

## Scope
A few things that I'd like to achieve:
1. Compile [protobuf](https://protobuf.dev/) services with RPC & HTTP supports
2. Setup client/server for GRPC communication via Go (maybe extend it to JavaScript/Java later)
3. Play around with [Bazel](https://bazel.build/) a bit

## Setup Requirements

```.sh
# Setup prerequisite
$ brew install buildifier bazelisk

```

For Editor Setup, please checkout [this](https://github.com/bazelbuild/rules_go/wiki/Editor-setup) document

## For testing

```
# Start the go server at port 9090
$ bazelisk run //cmd/server

# Startt the go client talking to the server via tcp:9090
$ bazelisk run //cmd/client
```