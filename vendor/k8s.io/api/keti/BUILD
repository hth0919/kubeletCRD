load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "addtoscheme_keti_v1.go",
        "apis.go",
    ],
    importmap = "k8s.io/kubernetes/vendor/k8s.io/api/keti",
    importpath = "k8s.io/api/keti",
    visibility = ["//visibility:public"],
    deps = [
        "//staging/src/k8s.io/api/keti/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//staging/src/k8s.io/api/keti/v1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
