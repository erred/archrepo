FROM golang:alpine AS build
ARG CGO_ENABLED=0
ARG GOPROXY=https://proxy.golang.org,direct
ARG GOMODCACHE=/go/pkg/mod
ARG GOCACHE=/root/.cache/go-build
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' ./cmd/archrepod

FROM archlinux:base
WORKDIR /opt/archrepo
COPY --from=build /workspace/archrepod /usr/local/bin/archrepod
ENTRYPOINT ["/usr/local/bin/archrepod"]
