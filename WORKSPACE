workspace(name = "grpc_proto_playground")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

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

GOOGLE_APIS_VERSION = "64926d52febbf298cb82a8f472ade4a3969ba922"

http_archive(
    name = "com_google_googleapis",
    integrity = "sha256-nRqTDnZ8k8glOYuPhpLso/41O5qq3t+88fyiKCyF34g=",
    strip_prefix = "googleapis-" + GOOGLE_APIS_VERSION,
    urls = [
        "https://github.com/googleapis/googleapis/archive/%s.zip" % GOOGLE_APIS_VERSION,
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

go_rules_dependencies()

go_register_toolchains(version = "1.20.7")
gazelle_dependencies()

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    grpc = True,
    gapic = True,
    go = True,
)