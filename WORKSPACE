#workspace(name = "test-repo")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.6/rules_go-0.18.6.tar.gz"],
    sha256 = "f04d2373bcaf8aa09bccb08a98a57e721306c8f6043a2a0ee610fd6853dcde3d"
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@bazel_gazelle//:deps.bzl", "go_repository")

# Do *not* call *_dependencies(), etc, yet.  See comment at the end.

go_repository(
  name = "com_github_antlr",
  tag = "4.7.2",
  importpath = "github.com/antlr/antlr4",
)

go_repository(
  name = "org_golang_x_text",
  importpath = "golang.org/x/text",
  tag = "v0.3.0",
  remote = "https://github.com/golang/text",
  vcs = "git",
)


# Run the dependencies at the end.  These will silently try to import some
# of the above repositories but at different versions, so ours must come first.
go_rules_dependencies()
go_register_toolchains()
gazelle_dependencies()
