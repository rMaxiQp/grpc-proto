# grpc-proto

GRPC Proto playground

## Scope

A few things that I'd like to achieve:

- [x] Compile [protobuf](https://protobuf.dev/) services with RPC & HTTP supports
- [x] Setup client/server for GRPC communication via Go & JavaScript
- [x] Play around with [Bazel](https://bazel.build/) a bit

## Setup Requirements

```.sh
# Setup prerequisite
$ brew install buildifier bazelisk
```

For Editor Setup, please checkout [this](https://github.com/bazelbuild/rules_go/wiki/Editor-setup) document

## Command lines

```
# To list out all existing rules (including auto-generated rules)
# Rules like `//:.aspect_rules_js/node_modules/@bufbuild+protobuf@1.10.0` is auto-generated
$ bazelisk query "//..."

# Start the Go server. Fixed ports: 9090 for gRPC, 9000 for HTTP
$ bazelisk run "//cmd/server"

# Start the JavaScript server
$ bazelisk run "//src:js_server"

# Send the HTTP payload to port 9000
$ bazelisk run "//cmd/client/http:client

# Send the gRPC payload to port 9090
$ bazelisk run "//cmd/client/rpc:client

# Send the gRPC payload to port 9090
$ bazelisk run "//src:js_client"

# Sync Go package dependencies with go.mod
$ gazelle update-repo

# Sync Go package dependencies in BUILD
$ gazelle update
```

## Learning

There are several components inside of this POC:

- Build system (Bazel)
- Type definition (Protobuf)
- Implementation (gRPC related packages)

### Bazel

Bazel is the open source version of Google's Blaze. It's a pain killer for managing projects in monorepo (after experiencing diffferent gotchas of `npm workspaces`). Via `deps` options, you cherry pick the packages you need to build the corresponding binaries. Being explicit for package imports helps me to trim unused dependencies, reducing the binary size. However, it has its own caveats:

- The learning curve is steep
  - Learn [Starlark](https://bazel.build/rules/language)
  - Understand Rules, Workspace, Module (and some [packages](https://github.com/bazelbuild/rules_go/issues/3685) does not work well with modules yet...)
- The community is still growing, and many things are still under development
  - `rules_go` has an official support via Protobuf already. But `rules_ts`'s support is still [WIP](https://docs.aspect.build/rulesets/aspect_rules_ts/docs/proto).

### Protobuf

Protobuf defines data schema that can be shared across different languages. But also, the protobuf can be sent as binary data and it's more packed than JSON. Other program languages can't interperate Protobuf files directly, and it needs a compiler to compile the proto definition into their languages. The common one is [`protoc`](https://grpc.io/docs/protoc-installation/).

### gRPC related packages

Protobuf has services, but they are RPC services. Thus, the default HTTP modules may struggle to inteperate them directly. Additional gRPC packages need to be installed/setup to make sure both client and server understand the protocol.
