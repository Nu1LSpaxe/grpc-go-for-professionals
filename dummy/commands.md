```bash=
bazel run //:gazelle

bazel run //:gazelle-update-repos
bazel run //:gazelle
bazel run //server:server 0.0.0.0:50051
```