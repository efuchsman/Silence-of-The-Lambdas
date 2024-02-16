load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    integrity = "sha256-fHbWI2so/2laoozzX5XeMXqUcv0fsUrHl8m/aE8Js3w=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
    ],
)

# Rules for creating ZIP files
http_archive(
    name = "io_bazel_rules_pkg",
    sha256 = "a880a793b916b0312e8d3e6ebd26128a9b64e38aa48dd20ef2bcfe25420a22a7",
    urls = ["https://github.com/bazelbuild/rules_pkg/releases/download/0.1.0/rules_pkg-0.1.0.tar.gz"],
)

# Gazelle
http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-MpOL2hbmcABjA1R5Bj2dJMYO2o15/Uc5Vj9Q0zHLMgk=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

http_archive(
    name = "logrus",
    sha256 = "a3c60bbde616e7579dbd104e47f5c0e9153f4e672b88e8d3a87f7f0f8852ad4a",
    strip_prefix = "logrus-1.9.3",
    urls = ["https://github.com/sirupsen/logrus/archive/v1.9.3.zip"],
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)
load(
    "@bazel_gazelle//:deps.bzl",
    "gazelle_dependencies",
    "go_repository",
)

go_repository(
    name = "com_github_aws_aws_lambda_go",
    importpath = "github.com/aws/aws-lambda-go",
    sum = "h1:UWVnvh2h2gecOlFhHQfIPQcD8pL/f7pVCutmFl+oXU8=",
    version = "v1.46.0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:FaXvNwHG3Ri1paUEW16Ahk9zLVqSAdqa1M3phjZR35Q=",
    version = "v1.50.6",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath_internal_testify",
    importpath = "github.com/jmespath/go-jmespath/internal/testify",
    sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_joho_godotenv",
    importpath = "github.com/joho/godotenv",
    sum = "h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
    version = "v1.8.4",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:obN1ZagJSUGI0Ek/LBmuj4SNLPfIny3KsKFopxRdj10=",
    version = "v2.2.8",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
    version = "v3.0.1",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:pVaXccu2ozPjCXewfr1S7xza/zcXTity9cCdXQYSjIM=",
    version = "v0.17.0",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:0A+M6Uqn+Eje4kHMK80dtF3JCXC4ykBgQG4Fe06QRhQ=",
    version = "v0.0.0-20220715151400-c0bba94af5f8",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:ablQoSUd0tRdKxZewP80B+BaqeKJuVhuRxj/dkrun3k=",
    version = "v0.13.0",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
    version = "v1.9.3",
)

go_rules_dependencies()

go_register_toolchains(version = "1.20.7")

gazelle_dependencies()
