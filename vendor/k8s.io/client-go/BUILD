load("@io_bazel_rules_go//go:def.bzl", "go_library")

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
        "//staging/src/k8s.io/client-go/discovery:all-srcs",
        "//staging/src/k8s.io/client-go/dynamic:all-srcs",
        "//staging/src/k8s.io/client-go/examples/create-update-delete-deployment:all-srcs",
        "//staging/src/k8s.io/client-go/examples/dynamic-create-update-delete-deployment:all-srcs",
        "//staging/src/k8s.io/client-go/examples/fake-client:all-srcs",
        "//staging/src/k8s.io/client-go/examples/in-cluster-client-configuration:all-srcs",
        "//staging/src/k8s.io/client-go/examples/leader-election:all-srcs",
        "//staging/src/k8s.io/client-go/examples/out-of-cluster-client-configuration:all-srcs",
        "//staging/src/k8s.io/client-go/examples/workqueue:all-srcs",
        "//staging/src/k8s.io/client-go/informers:all-srcs",
        "//staging/src/k8s.io/client-go/kubernetes:all-srcs",
        "//staging/src/k8s.io/client-go/kubernetes_test:all-srcs",
        "//staging/src/k8s.io/client-go/listers/admissionregistration/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/admissionregistration/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/apps/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/apps/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/apps/v1beta2:all-srcs",
        "//staging/src/k8s.io/client-go/listers/auditregistration/v1alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/authentication/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/authentication/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/authorization/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/authorization/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/autoscaling/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/autoscaling/v2beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/autoscaling/v2beta2:all-srcs",
        "//staging/src/k8s.io/client-go/listers/batch/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/batch/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/batch/v2alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/certificates/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/coordination/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/coordination/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/core/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/discovery/v1alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/discovery/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/events/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/extensions/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/flowcontrol/v1alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/imagepolicy/v1alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/networking/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/networking/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/node/v1alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/node/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/policy/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/rbac/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/rbac/v1alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/rbac/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/scheduling/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/scheduling/v1alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/scheduling/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/settings/v1alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/storage/v1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/storage/v1alpha1:all-srcs",
        "//staging/src/k8s.io/client-go/listers/storage/v1beta1:all-srcs",
        "//staging/src/k8s.io/client-go/metadata:all-srcs",
        "//staging/src/k8s.io/client-go/pkg/apis/clientauthentication:all-srcs",
        "//staging/src/k8s.io/client-go/pkg/version:all-srcs",
        "//staging/src/k8s.io/client-go/plugin/pkg/client/auth:all-srcs",
        "//staging/src/k8s.io/client-go/rest:all-srcs",
        "//staging/src/k8s.io/client-go/restmapper:all-srcs",
        "//staging/src/k8s.io/client-go/scale:all-srcs",
        "//staging/src/k8s.io/client-go/testing:all-srcs",
        "//staging/src/k8s.io/client-go/third_party/forked/golang/template:all-srcs",
        "//staging/src/k8s.io/client-go/tools/auth:all-srcs",
        "//staging/src/k8s.io/client-go/tools/cache:all-srcs",
        "//staging/src/k8s.io/client-go/tools/clientcmd:all-srcs",
        "//staging/src/k8s.io/client-go/tools/events:all-srcs",
        "//staging/src/k8s.io/client-go/tools/leaderelection:all-srcs",
        "//staging/src/k8s.io/client-go/tools/metrics:all-srcs",
        "//staging/src/k8s.io/client-go/tools/pager:all-srcs",
        "//staging/src/k8s.io/client-go/tools/portforward:all-srcs",
        "//staging/src/k8s.io/client-go/tools/record:all-srcs",
        "//staging/src/k8s.io/client-go/tools/reference:all-srcs",
        "//staging/src/k8s.io/client-go/tools/remotecommand:all-srcs",
        "//staging/src/k8s.io/client-go/tools/watch:all-srcs",
        "//staging/src/k8s.io/client-go/transport:all-srcs",
        "//staging/src/k8s.io/client-go/util/cert:all-srcs",
        "//staging/src/k8s.io/client-go/util/certificate:all-srcs",
        "//staging/src/k8s.io/client-go/util/connrotation:all-srcs",
        "//staging/src/k8s.io/client-go/util/exec:all-srcs",
        "//staging/src/k8s.io/client-go/util/flowcontrol:all-srcs",
        "//staging/src/k8s.io/client-go/util/homedir:all-srcs",
        "//staging/src/k8s.io/client-go/util/jsonpath:all-srcs",
        "//staging/src/k8s.io/client-go/util/keyutil:all-srcs",
        "//staging/src/k8s.io/client-go/util/retry:all-srcs",
        "//staging/src/k8s.io/client-go/util/testing:all-srcs",
        "//staging/src/k8s.io/client-go/util/workqueue:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["clientset.go"],
    importmap = "k8s.io/kubernetes/vendor/k8s.io/client-go",
    importpath = "k8s.io/client-go",
    visibility = ["//visibility:public"],
    deps = [
        "//staging/src/k8s.io/client-go/discovery:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/admissionregistration/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/apps/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/apps/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/apps/v1beta2:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/auditregistration/v1alpha1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/authentication/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/authentication/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/authorization/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/authorization/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/autoscaling/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/batch/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/batch/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/batch/v2alpha1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/certificates/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/coordination/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/coordination/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/core/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/discovery/v1alpha1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/discovery/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/events/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/extensions/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/keti/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/networking/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/networking/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/node/v1alpha1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/node/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/policy/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/rbac/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/rbac/v1alpha1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/rbac/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/scheduling/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/scheduling/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/settings/v1alpha1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/storage/v1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/storage/v1alpha1:go_default_library",
        "//staging/src/k8s.io/client-go/kubernetes/typed/storage/v1beta1:go_default_library",
        "//staging/src/k8s.io/client-go/rest:go_default_library",
        "//staging/src/k8s.io/client-go/util/flowcontrol:go_default_library",
    ],
)
