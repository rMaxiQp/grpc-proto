#!/usr/bin/env bash

# https://github.com/bazelbuild/rules_go/issues/3110#issuecomment-1123581408
# Load std lib from local instead of remote
GOPACKAGESDRIVER_BAZEL_BUILD_FLAGS=--strategy=GoStdlibList=local
exec bazelisk run -- @io_bazel_rules_go//go/tools/gopackagesdriver "${@}"