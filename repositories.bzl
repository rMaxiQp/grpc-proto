"""
This file contains the external repositories for the project.
"""
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def http_dependencies():
    """
    A list of external dependencies
    """
    http_archive(
        name = "io_bazel_rules_go",
        integrity = "sha256-fHbWI2so/2laoozzX5XeMXqUcv0fsUrHl8m/aE8Js3w=",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
            "https://github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
        ],
    )

    http_archive(
        name = "bazel_gazelle",
        integrity = "sha256-MpOL2hbmcABjA1R5Bj2dJMYO2o15/Uc5Vj9Q0zHLMgk=",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
            "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        ],
    )

    http_archive(
        name = "googleapis",
        sha256 = "9d1a930e767c93c825398b8f8692eca3fe353b9aaadedfbcf1fca2282c85df88",
        strip_prefix = "googleapis-64926d52febbf298cb82a8f472ade4a3969ba922",
        urls = [
            "https://github.com/googleapis/googleapis/archive/64926d52febbf298cb82a8f472ade4a3969ba922.zip",
        ],
    )

    http_archive(
        name = "rules_proto_grpc",
        sha256 = "2a0860a336ae836b54671cbbe0710eec17c64ef70c4c5a88ccfd47ea6e3739bd",
        strip_prefix = "rules_proto_grpc-4.6.0",
        urls = [
            "https://github.com/rules-proto-grpc/rules_proto_grpc/releases/download/4.6.0/rules_proto_grpc-4.6.0.tar.gz",
        ],
    )

    http_archive(
        name = "rules_nodejs",
        sha256 = "dddd60acc3f2f30359bef502c9d788f67e33814b0ddd99aa27c5a15eb7a41b8c",
        strip_prefix = "rules_nodejs-6.1.0",
        url = "https://github.com/bazelbuild/rules_nodejs/releases/download/v6.1.0/rules_nodejs-v6.1.0.tar.gz",
    )

    http_archive(
        name = "aspect_rules_ts",
        sha256 = "b11f5bd59983a58826842029b99240fd0eeb6f1291d710db10f744b327701646",
        strip_prefix = "rules_ts-2.3.0",
        url = "https://github.com/aspect-build/rules_ts/releases/download/v2.3.0/rules_ts-v2.3.0.tar.gz",
    )

    http_archive(
        name = "aspect_rules_js",
        sha256 = "bc9b4a01ef8eb050d8a7a050eedde8ffb1e45a56b0e4094e26f06c17d5fcf1d5",
        strip_prefix = "rules_js-1.41.2",
        url = "https://github.com/aspect-build/rules_js/releases/download/v1.41.2/rules_js-v1.41.2.tar.gz",
    )
