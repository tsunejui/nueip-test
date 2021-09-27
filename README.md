## NUEiP Test

Quick Start (Testing on local)
---

1. compile binary for api server
```
go run build.go --target=api-server --pkg=./cmd/api-server --output=build/apps/api-server/bin --os=linux --arch=amd64
```

2. `cd build/apps/api-server` and start to build image
```
docker build . -f Dockerfile -t api-server
```
